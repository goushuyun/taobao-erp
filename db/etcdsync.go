package db

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/wothing/log"
	"golang.org/x/net/context"
)

const (
	defaultTTL   = 60
	defaultTry   = 3
	deleteAction = "delete"
	expireAction = "expire"
)

// A Mutex is a mutual exclusion lock which is distributed across a cluster.
type Mutex struct {
	key    string
	id     string // The identity of the caller
	client client.Client
	kapi   client.KeysAPI
	ctx    context.Context
	ttl    time.Duration
	mutex  *sync.Mutex
}

func NewEtcdSync(key string, ttl int) (*Mutex, error) {
	endpoints := strings.Split(GetEtcd(), ",")
	ret, err := newEtcdSync(key, ttl, endpoints)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// New creates a Mutex with the given key which must be the same
// across the cluster nodes.
// machines are the ectd cluster addresses
func newEtcdSync(key string, ttl int, machines []string) (*Mutex, error) {
	cfg := client.Config{
		Endpoints:               machines,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	if len(key) == 0 || len(machines) == 0 {
		return nil, errors.New("Key is Empty Or machines is Empty!")
	}

	if key[0] != '/' {
		key = "/" + key
	}

	if ttl < 1 {
		ttl = defaultTTL
	}

	return &Mutex{
		key:    key,
		id:     fmt.Sprintf("%v-%v-%v", hostname, os.Getpid(), time.Now().Format("20060102-15:04:05.999999999")),
		client: c,
		kapi:   client.NewKeysAPI(c),
		ctx:    context.TODO(),
		ttl:    time.Second * time.Duration(ttl),
		mutex:  new(sync.Mutex),
	}, nil
}

// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *Mutex) Lock() (err error) {
	m.mutex.Lock()
	for try := 1; try <= defaultTry; try++ {
		if m.lock() == nil {
			return nil
		}
		log.Debugf("Lock node %v ERROR %v", m.key, err)
		if try < defaultTry {
			log.Debugf("Try to lock node %v again", m.key, err)
		}
	}
	return err
}

//lock函数，创建key,如果创建成功，返回nil，创建失败，则监控已经存在key,如果key被删除，或者超时则返回nil。

func (m *Mutex) lock() (err error) {
	log.Debugf("Trying to create a node : key=%v", m.key)
	setOptions := &client.SetOptions{
		PrevExist: client.PrevNoExist,
		TTL:       m.ttl,
	}
	resp, err := m.kapi.Set(m.ctx, m.key, m.id, setOptions)
	if err == nil {
		log.Debugf("Create node %v OK [%q]", m.key, resp)
		return nil
	}
	log.Debugf("Create node %v failed [%v]", m.key, err)
	e, ok := err.(client.Error)
	if !ok {
		return err
	}

	if e.Code != client.ErrorCodeNodeExist {
		return err
	}

	// Get the already node's value.
	resp, err = m.kapi.Get(m.ctx, m.key, nil)
	if err != nil {
		return err
	}
	log.Debugf("Get node %v OK", m.key)
	watcherOptions := &client.WatcherOptions{
		AfterIndex: resp.Index,
		Recursive:  false,
	}
	watcher := m.kapi.Watcher(m.key, watcherOptions)
	for {
		log.Debugf("Watching %v ...", m.key)
		resp, err = watcher.Next(m.ctx)
		if err != nil {
			return err
		}

		log.Debug("Received an event : %q", resp)
		if resp.Action == deleteAction || resp.Action == expireAction {
			return nil
		}
	}

}

// Unlock unlocks m.
// It is a run-time error if m is not locked on entry to Unlock.
//
// A locked Mutex is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
func (m *Mutex) Unlock() (err error) {
	defer m.mutex.Unlock()
	for i := 1; i <= defaultTry; i++ {
		var resp *client.Response
		resp, err = m.kapi.Delete(m.ctx, m.key, nil)
		if err == nil {
			log.Debugf("Delete %v OK", m.key)
			return nil
		}
		log.Debugf("Delete %v falied: %q", m.key, resp)
		e, ok := err.(client.Error)
		if ok && e.Code == client.ErrorCodeKeyNotFound {
			return nil
		}
	}
	return err
}
