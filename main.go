package main

import (
	"stu_queue/s_rocketmq"
	"time"
)

func main(){
	s_rocketmq.Producer1()
	time.Sleep(1000*time.Second)
	s_rocketmq.Consumer1()

}
