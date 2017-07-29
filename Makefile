.PHONY: coupon sms leaflet profile get search proposal credit

all: idl ve para

idl:
	rm -rf pb/*.pb.go
	protoc -I=. pb/*.proto --go_out=plugins=grpc:.
ve:
	rm -rf ver
	mkdir ver
	echo package ver > ver/version.go
	echo const GitCommit string = \"$$(git rev-parse HEAD)\" >> ver/version.go
	echo const BuildDate string = \"$$(date)\" >> ver/version.go

para: acct act ba cfdnt confi coupon credit eam gw gouds hsptl leaflet ms mes oder o pay profile proposal pu search sha sms stat st usr vfc vers wc jigsaw

acct:
	cd account; GOBIN=$(TD) go install

act:
	cd activity; GOBIN=$(TD) go install

ba:
	cd banner; GOBIN=$(TD) go install

cfdnt:
	cd confidant; GOBIN=$(TD) go install

confi:
	cd config; GOBIN=$(TD) go install

coupon:
	cd coupon; GOBIN=$(TD) go install

credit:
	cd credit; GOBIN=$(TD) go install

eam:
	cd easemob; GOBIN=$(TD) go install

gw:
	cd gateway/appway; GOBIN=$(TD) go install
	cd gateway/interway; GOBIN=$(TD) go install
	cd gateway/hospway; GOBIN=$(TD) go install

gouds:
	cd goods; GOBIN=$(TD) go install

hsptl:
	cd hospital; GOBIN=$(TD) go install

leaflet:
	cd leaflet; GOBIN=$(TD) go install

ms:
	cd mediastore; GOBIN=$(TD) go install

mes:
	cd message; GOBIN=$(TD) go install

oder:
	cd order; GOBIN=$(TD) go install

o:
	cd other; GOBIN=$(TD) go install

pay:
	cd payment; GOBIN=$(TD) go install

profile:
	cd profile; GOBIN=$(TD) go install

proposal:
	cd proposal; GOBIN=$(TD) go install

pu:
	cd push; GOBIN=$(TD) go install

search:
	cd search; GOBIN=$(TD) go install

sha:
	cd share; GOBIN=$(TD) go install

sms:
	cd sms; GOBIN=$(TD) go install

stat:
	cd statistic; GOBIN=$(TD) go install

st:
	cd story; GOBIN=$(TD) go install

usr:
	cd user; GOBIN=$(TD) go install

vfc:
	cd verification; GOBIN=$(TD) go install

vers:
	cd version; GOBIN=$(TD) go install

wc:
	cd wechat; GOBIN=$(TD) go install

jigsaw:
	cd jigsaw; GOBIN=$(TD) go install

knowlebase:
	cd knowlebase; GOBIN=$(TD) go install

etcdTe:
	curl -Ls ${ETCD}/v2/keys/17mei/mode -XPUT -d value="test" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/pgsql/host -XPUT -d value="127.0.0.1" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/pgsql/port -XPUT -d value="5432" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/pgsql/name -XPUT -d value="meidb" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/pgsql/user -XPUT -d value="postgres" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/pgsql/password -XPUT -d value="" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/redis/host -XPUT -d value="127.0.0.1" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/redis/port -XPUT -d value="6379" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/mediastore/mode -XPUT -d value="test" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/payment/mode -XPUT -d value="test" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/proposal/mode -XPUT -d value="test" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/push/apns -XPUT -d value="false" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/nsql/host -XPUT -d value="127.0.0.1" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/nsql/port -XPUT -d value="4161" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/nsqd/host -XPUT -d value="127.0.0.1" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/nsqd/port -XPUT -d value="4150" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/wechat/web/appid -XPUT -d value="wxc3a713d594283b00" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/wechat/web/appsecret -XPUT -d value="66edd83a09789b1fb88535e3f14ae94c" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/wechat/app/appid -XPUT -d value="wxdd27e57ca55e6d97" > /dev/null
	curl -Ls ${ETCD}/v2/keys/17mei/wechat/app/appsecret -XPUT -d value="cca8a5c15d5b32c4ecab58d85b9cb044" > /dev/null

etcdPro:
	curl -L ${ETCD}/v2/keys/17mei/mode -XPUT -d value="prod" > /dev/null
	curl -L ${ETCD}/v2/keys/17mei/pgsql/host -XPUT -d value="192.168.0.5"
	curl -L ${ETCD}/v2/keys/17mei/pgsql/port -XPUT -d value="5432"
	curl -L ${ETCD}/v2/keys/17mei/pgsql/name -XPUT -d value="meidb"
	curl -L ${ETCD}/v2/keys/17mei/pgsql/user -XPUT -d value="postgres"
	curl -L ${ETCD}/v2/keys/17mei/pgsql/password -XPUT -d value="NOT-PUBLIC"
	curl -L ${ETCD}/v2/keys/17mei/redis/host -XPUT -d value="192.168.0.8"
	curl -L ${ETCD}/v2/keys/17mei/redis/port -XPUT -d value="6379"
	curl -L ${ETCD}/v2/keys/17mei/redis/password -XPUT -d value="NOT-PUBLIC"

	curl -L ${ETCD}/v2/keys/17mei/mediastore/mode -XPUT -d value="live"

	curl -L ${ETCD}/v2/keys/17mei/payment/mode -XPUT -d value="live"
	curl -L ${ETCD}/v2/keys/17mei/proposal/mode -XPUT -d value="live"
	curl -L ${ETCD}/v2/keys/17mei/push/apns -XPUT -d value="true"
	curl -L ${ETCD}/v2/keys/17mei/nsql/host -XPUT -d value="192.168.0.8"
	curl -L ${ETCD}/v2/keys/17mei/nsql/port -XPUT -d value="4161"
	curl -L ${ETCD}/v2/keys/17mei/nsqd/host -XPUT -d value="192.168.0.8"
	curl -L ${ETCD}/v2/keys/17mei/nsqd/port -XPUT -d value="4150"
	curl -L ${ETCD}/v2/keys/17mei/wechat/web/appid -XPUT -d value="wx6cacb7fe38e16098"
	curl -L ${ETCD}/v2/keys/17mei/wechat/web/appsecret -XPUT -d value="b899e84146f673e99e02adad3414deb8"
	curl -L ${ETCD}/v2/keys/17mei/wechat/app/appid -XPUT -d value="wxdd27e57ca55e6d97"
	curl -L ${ETCD}/v2/keys/17mei/wechat/app/appsecret -XPUT -d value="cca8a5c15d5b32c4ecab58d85b9cb044"

get:
	go get github.com/lib/pq
	go get github.com/jackc/pgx
	go get github.com/garyburd/redigo/redis
	go get github.com/coreos/etcd/client
	go get github.com/nsqio/go-nsq
	go get golang.org/x/crypto/bcrypt
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/urfave/negroni
	go get github.com/julienschmidt/httprouter
	go get github.com/gorilla/context
	go get github.com/dgrijalva/jwt-go
	go get github.com/pborman/uuid
	go get github.com/wothing/log
	go get github.com/wothing/worc
	go get github.com/wothing/worpc
	go get github.com/wothing/wonaming/etcd
	go get github.com/elgs/gostrgen
	go get qiniupkg.com/api.v7/kodo
	go get qiniupkg.com/x/url.v7
	go get github.com/ylywyn/jpush-api-go-client
	go get github.com/pingplusplus/pingpp-go/pingpp
	go get github.com/smartystreets/assertions
	go get github.com/smartystreets/goconvey
	go get github.com/bitly/go-simplejson
	go get github.com/tealeg/xlsx
	go get github.com/googollee/go-socket.io

startAll:
	for app in `ls /drone/app`; do \
	  nohup /drone/app/$$app -etcd=$(ETCD) & \
	done
	sleep 90
	go test -v gateway/tests/*.go

gv:
	dot -T png 17mei.gv -o 17mei.png
