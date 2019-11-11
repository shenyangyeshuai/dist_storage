package locate

import (
	"../../rabbitmq"
	"../objects"
	"os"
)

func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()

	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		// 消息的正文是接口服务发送过来的需要做定位的对象的名字
		object, e := strconv.Unqote(string(msg.Body))
		if e != nil {
			panic(e)
		}

		if Locate(objects.STORAGE_ROOT + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}

func Locate(name string) bool {
	_, e := os.Stat(name)
	return !os.IsNotExist(e)
}
