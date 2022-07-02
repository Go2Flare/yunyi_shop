package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	//延迟消息，是发送后，经过了延迟时间后，消费这才能收到消息
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"120.78.192.25:9876"}))
	if err != nil {
		panic("生成producer失败")
	}

	if err = p.Start(); err != nil {panic("启动producer失败")}

	msg := primitive.NewMessage("flare", []byte("this is delay message"))
	//reference delay level definition: 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	msg.WithDelayTimeLevel(2)//30s
	res, err := p.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("发送失败: %s\n", err)
	}else{
		fmt.Printf("发送成功: %s\n", res.String())
	}

	if err = p.Shutdown(); err != nil {panic("关闭producer失败")}

	//支付的时候， 淘宝， 12306， 购票， 超时归还 - 定时执行逻辑
	//我可以去写一个轮询， 轮询的问题： 1. 多久执行一次轮询 30分钟
	//在12:00执行过一次轮询， 下一次执行就是在 12:30的时候 但是12:01的时候下了单， 12:31就应该超时 13:00时候才能超时
	//那我1分钟执行一次啊， 比如我的订单量没有这么大，1分钟执行一次， 其中29次查询都是无用， 而且你还还会轮询mysql，消耗的资源太多
	//rocketmq的延迟消息， 1. 时间一到就执行， 2. 消息中包含了订单编号，你只查询这种订单编号
}