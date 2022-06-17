## rocketmq
### 普通安装
1. 需要java环境
2. 下载
> wget https://archive.apache.org/dist/rocketmq/4.9.3/rocketmq-all-4.9.3-bin-release.zip
3. 解压
> unzip rocketmq-all-4.9.3-bin-release.zip

4. cd rocketmq-4.9.3/bin
5. nohup sh mqnamesrv & // 启动命名服务
6. nohup sh bin/mqbroker -n localhost:9876 & // 启动broker

### docker安装
1. 拉取镜像
> docker pull apache/rocketmq
2. 启动命名服务,默认9876端口
> sudo docker run -it --net=host apache/rocketmq ./mqnamesrv
3. 启动broker
> sudo docker run -it --net=host apache/rocketmq ./mqbroker -n localhost:9876
4. 启动控制台
> docker run -e "JAVA_OPTS=-Drocketmq.namesrv.addr={namesrv_addr} -Dcom.rocketmq.sendMessageWithVIPChannel=false" -p 8080:8080 -t styletang/rocketmq-console-ng

docker run --net=host -e "JAVA_OPTS=-Drocketmq.namesrv.addr=127.0.0.1:9876 -Dcom.rocketmq.sendMessageWithVIPChannel=false" -t styletang/rocketmq-console-ng

5. 浏览器打开http://127.0.0.1:8080/即可访问

## kafka

1. wget https://archive.apache.org/dist/zookeeper/zookeeper-3.4.14/zookeeper-3.4.14.tar.gz
2. bash zkServer.sh start
3. bin/kafka-server-start.sh config/server.properties