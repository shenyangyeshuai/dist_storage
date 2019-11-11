package heartbeat

import (
	"../../rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()

	// 当前的数据服务节点
	// 每隔 5 秒发送一个心跳
	// 给 apiServers exchange
	// 告知当前节点的监听地址
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(time.Second * 5)
	}
}
