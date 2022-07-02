package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	//Pull模型，客户端轮询服务器，不断询问有无消息
	//Push模型，服务器有了数据后，会主动推送消息到客户端，更省资源一些
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"120.78.192.25:9876"}),
		consumer.WithGroupName("flare"),//groupname，应该可以在同一group中实现负载均衡
	)

	if err := c.Subscribe("flare", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("RMQ获取到值： %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	//不能让主goroutine退出
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
