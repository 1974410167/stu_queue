package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"html"
	"log"
	"net/http"
	"stu_queue/s_kafka"
	"stu_queue/s_rocketmq"
	"time"
)

func main(){
	runKafka()

}

func runRocketMq(){
	s_rocketmq.CreateTopic("Topic-test")
	s_rocketmq.Producer1()
	time.Sleep(1000*time.Second)
	s_rocketmq.Consumer1()
}

func runKafka(){
	topic := "sample-kafka"
	producer, err := s_kafka.NewProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	consumer, err := sarama.NewConsumer(s_kafka.Brokers, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	s_kafka.Subscribe(topic, consumer)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello Sarama!") })

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		r.ParseForm()
		msg := s_kafka.PrepareMessage(topic, r.FormValue("q"))
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Fprintf(w, "%s error occured.", err.Error())
		} else {
			fmt.Fprintf(w, "Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
		}
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, html.EscapeString(s_kafka.GetMessage())) })

	log.Fatal(http.ListenAndServe(":8088", nil))
}