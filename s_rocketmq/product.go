package s_rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
)

func Producer1() {
	p, _ := rocketmq.NewProducer(
		// 设置  nameSrvAddr
		// nameSrvAddr 是 Topic 路由注册中心
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		// 指定发送失败时的重试时间
		producer.WithRetry(2),
		// 设置 Group
		producer.WithGroupName("testGroup"),
	)
	// 开始连接
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})))
	if err != nil{
		fmt.Println("admin创建失败",err)
		os.Exit(1)
	}
	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate("Topic-test"),
		admin.WithBrokerAddrCreate("192.169.1.2:10911"),
	)
	if err != nil{
		fmt.Println("topic创建失败",err)
		os.Exit(1)
	}
	// 设置节点名称
	topic := "Topic-test"
	// 循坏发送信息 (同步发送)
	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ Go Client" + strconv.Itoa(i)),
		}
		// 发送信息
		res, err := p.SendSync(context.Background(),msg)
		if err != nil {
			fmt.Printf("send message error:%s\n",err)
		}else {
			fmt.Printf("send message success: result=%s\n",res.String())
		}
	}
	// 关闭生产者
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error:%s",err.Error())
	}
}