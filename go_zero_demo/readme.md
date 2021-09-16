# 项目说明
这个项目是用于处理客户更新训练数据，用户调用接口后对数据格式进行校验，成功的就发送给下游进行更新，失败的丢弃，同时统计用户调用接口的数据，项目从python改造而来。

# 开发步骤
项目基于go-zero框架开发
- 安装goctl
- 在 pinggo.api中新增对应接口
- 使用 goctl api go -api pinggo.api -dir . 生成接口代码
- 在internal/handler里面修改对应接口的实现逻辑即可

# 选型
## kafka依赖包
- https://github.com/Shopify/sarama
  github上star最多的项目，使用者众多，开发者活跃
- https://github.com/segmentio/kafka-go
  github上star第二的项目，使用者众多
- https://github.com/confluentinc/confluent-kafka-go
  kafka官方出品


# 附录
##根据proto生成.pb文件
### 工具安装
// TODO
### 生成命令
protoc --proto_path=./proto --plugin=protoc-gen-go.exe --go_out=proto/ --go_opt=paths=source_relative proto/message.proto





