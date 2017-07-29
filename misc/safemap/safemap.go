/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/08/11 17:28
 */

package safemap

import "sync"

type safeMap struct {
	lock *sync.RWMutex
	sm   map[string]interface{}
}

func NewSafeMap() *safeMap {
	return &safeMap{
		lock: new(sync.RWMutex),
		sm:   make(map[string]interface{}),
	}
}

//Get from maps return the k's value
func (m *safeMap) Get(k string) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.sm[k]; ok {
		return val
	}
	return nil
}

//Set Maps the given key and value. Returns false if the key is already in the map and changes nothing.
func (m *safeMap) Set(k string, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.sm[k]; !ok {
		m.sm[k] = v
	} else if val != v {
		m.sm[k] = v
	} else {
		return false
	}
	return true
}

//Check returns true if k is exist in the map.
func (m *safeMap) Check(k string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.sm[k]; ok {
		return true
	}
	return false
}

func (m *safeMap) Delete(k string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.sm, k)
}
