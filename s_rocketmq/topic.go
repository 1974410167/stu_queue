package s_rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)
var nameSrvAddr = []string{"127.0.0.1:9876"}
var brokerAddr = "127.0.0.1:10911"
var testAdmin, err = admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))

func CreateTopic(topicName string){
	topic := topicName
	//clusterName := "DefaultCluster"

	//create topic
	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr),
	)
	if err != nil {
		fmt.Println("Create topic error:", err.Error())
	}
	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}
}

func DeleteTopic(topicName string){
	topic := topicName

	err = testAdmin.DeleteTopic(
		context.Background(),
		admin.WithTopicDelete(topic),
		//admin.WithBrokerAddrDelete(brokerAddr),
		//admin.WithNameSrvAddr(nameSrvAddr),
	)
	if err != nil {
		fmt.Println("Delete topic error:", err.Error())
	}

	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}
}
