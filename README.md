# Auroraride Adapter

## 各项服务绑定端口说明



> 端口基本规则: 
>
> 	1. 4位数字
> 	1. 前两位为`服务代号`，后两位为  `端口代号`
> 	1. 对外提供服务端口前一律加11400, 例如: 5110 -> 16510
> 	1. 测试环境端口一律加13000, 例如: 5530 -> 18530, 16510 -> 29510



#### 业务及代号

|   业务   | 代号 |     备注     |
| :------: | :--: | :----------: |
| aurservd |  50  |    主业务    |
|  kxcab   |  51  |  凯信智能柜  |
|  xcbms   |  52  | 星创智能电池 |



#### 端口代号

| 端口类型 | 代号 |
| :------: | :--: |
|   HTTP   |  1x  |
|   TCP    |  2x  |
|   gRPC   |  3x  |



#### 端口列表


- aurservd
    - api: `127.0.0.1:5010`
    - tcp (对微服务提供连接服务)
        - kxcab: `127.0.0.1:5020`
        - xcbms:`127.0.0.1:5021`
- kxcab
    - api: `127.0.0.1:5110`
    - tcp: `127.0.0.1:16510`
- xcbms
  
    - api: `127.0.0.1:5210`
    - exhook: `127.0.0.1:5230`
    





## 日志数据库

- ELK
- ClickHouse

## 日志库

- [一文告诉你如何用好uber开源的zap日志库](https://tonybai.com/2021/07/14/uber-zap-advanced-usage/)
- [一文告诉你如何用好uber开源的zap日志库](https://mp.weixin.qq.com/s?__biz=MzIyNzM0MDk0Mg%3D%3D&chksm=e863f0fadf1479ec6a0138cede9923f44ca158a5e3dcab3d22de56deb6eca56bb0fd9db2e367&idx=1&mid=2247489307&scene=21&sn=0fd725e4be08b40d1e73e53600433910)
- [golang使用Zap日志库](https://zhuanlan.zhihu.com/p/371547318)
- [ECS Logging Go (zap) Reference](https://www.elastic.co/guide/en/ecs-logging/go-zap/current/setup.html)
- [Run Filebeat on Dockeredit](https://www.elastic.co/guide/en/beats/filebeat/master/running-on-docker.html)
- [Configuring Centralized logging with Kafka and ELK stack](https://2much2learn.com/centralized-logging-with-kafka-and-elk-stack/)
- [[Go] 基于 Zap 与 ELK 的日志分析实践](https://juejin.cn/post/6844904039793033223)
- [Go语言高性能日志库zap使用](https://huangzhongde.cn/post/Golang/2020-03-07-golang_logger_zap/)
- [在Go中集成ELK服务](https://jasonkayzk.github.io/2021/05/16/%E5%9C%A8Go%E4%B8%AD%E9%9B%86%E6%88%90ELK%E6%9C%8D%E5%8A%A1/)

## MQTT

- [emqx: 多语言 - 钩子扩展](https://www.emqx.io/docs/zh/v5.0/advanced/lang-exhook.html)
- [emqx: MQTT Go 客户端库](https://www.emqx.io/docs/zh/v5.0/development/go.html)
- [如何在 Golang 中使用 MQTT](https://www.emqx.com/zh/blog/how-to-use-mqtt-in-golang)

## 字节

- [Golang binary包——byte数组如何转int？](https://studygolang.com/articles/1122)