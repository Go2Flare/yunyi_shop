package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type OrderListener struct{}

//执行本地的事务，事务的一致性，可以执行rollback
func (o *OrderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行本地逻辑")
	time.Sleep(time.Second*3)
	fmt.Println("执行本地逻辑失败")
	//本地执行逻辑无缘无故失败 代码异常 宕机
	//UnknowState状态会进行回查
	return primitive.UnknowState
	//return primitive.CommitMessageState //正常情况，逻辑执行成功。不会回查
	//return primitive.RollbackMessageState //异常情况，逻辑执行失败。直接rollback
}

// 消息回查逻辑
func (o *OrderListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("rocketmq的消息回查")
	time.Sleep(time.Second*10)
	return primitive.CommitMessageState
}

func main() {
	//新建一个事务的producer
	p, err := rocketmq.NewTransactionProducer(
		&OrderListener{},
		producer.WithNameServer([]string{"120.78.192.25:9876"}),
		)
	if err != nil {
		panic("生成producer失败")
	}

	if err = p.Start(); err != nil {panic("启动producer失败")}

	res, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage("flare", []byte("this is a transaction message")))
	if err != nil {
		fmt.Printf("发送失败: %s\n", err)
	}else{
		fmt.Printf("发送成功: %s\n", res.String())
	}
	
	time.Sleep(time.Hour)
	if err = p.Shutdown(); err != nil {panic("关闭producer失败")}
}
