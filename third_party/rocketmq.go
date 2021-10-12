package third_party

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

const (
	RetryTime = 3
	SleepTime = time.Millisecond * 10
)

type RocketMQClient struct {
	nameServer []string `json:"name_server"`
}

var (
	MqProducer            rocketmq.Producer
	MqPushConsumerSuccess rocketmq.PushConsumer
)

func (rmq RocketMQClient) Init() (err error) {
	MqProducer, err = rocketmq.NewProducer(
		producer.WithGroupName("notify_producer"),
		producer.WithNameServer(rmq.nameServer),
		producer.WithRetry(RetryTime),
	)
	if err != nil {
		panic(fmt.Sprintf("init rocket mq producer err:%v", err))
		return
	}

	err = MqProducer.Start()
	if err != nil {
		panic(fmt.Sprintf("producer mq start err:%v", err))
		return
	}

	MqPushConsumerSuccess, err = rocketmq.NewPushConsumer(
		consumer.WithGroupName("notify_consumer_success"),
		consumer.WithNameServer(rmq.nameServer),
	)
	if err != nil {
		panic(fmt.Sprintf("init rocket mq push consumer err:%v", err))
		return
	}
	return
}

func ShutDownMq() {
	_ = MqProducer.Shutdown()
	_ = MqPushConsumerSuccess.Shutdown()
}

func (rmq RocketMQClient) SubScribe() {
	var err error
	err = MqPushConsumerSuccess.Subscribe("", consumer.MessageSelector{}, callBackSucc)
	if err != nil {
		panic(fmt.Sprintf("consumer subscribe TopicNotifySucess err:%v", err))
		return
	}

	err = MqPushConsumerSuccess.Start()
	if err != nil {
		panic(fmt.Sprintf("MqPushConsumerSuccess start err:%v", err))
		return
	}
}

func callBackSucc(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for i := range msgs {
		fmt.Println(string(msgs[i].Body))
	}
	return consumer.ConsumeSuccess, nil
}

func (rmq RocketMQClient) SendSuccNotifyByMq(taskID string) (err error) {
	msg := primitive.NewMessage("", []byte(taskID))
	//msg.WithTag(include.TagSuccNotify)
	for i := 0; i < RetryTime; i++ {
		res, err := MqProducer.SendSync(context.Background(), msg)
		if err != nil {
			time.Sleep(SleepTime)
			continue
		}
		fmt.Println(res)
		break
	}
	return

}
