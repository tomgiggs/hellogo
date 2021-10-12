package middleware
//
//import (
//	"context"
//	"github.com/Shopify/sarama"
//	"gitlab.deepwisdomai.com/infra/go-zero/core/logx"
//	"log"
//	"os"
//	"os/signal"
//	"pinggo/internal/config"
//	"sync"
//	"syscall"
//)
//
////var _ Queue = (*KafkaConsumer)(nil)
//
//var (
//	kafkaConsumerInstance *KafkaConsumer
//	kafkaConsumerOnce     = sync.Once{}
//)
//
//func GetKafkaConsumerInstance() *KafkaConsumer {
//	kafkaConsumerOnce.Do(func() {
//		kafkaConsumerInstance = new(KafkaConsumer)
//		kafkaConsumerInstance.Init(context.TODO(), *config.AppConfig)
//	})
//	return kafkaConsumerInstance
//}
//
//type KafkaConsumer struct {
//	client     sarama.ConsumerGroup
//	clientConf *sarama.Config
//}
//
//func (kc *KafkaConsumer) Init(ctx context.Context, conf config.Config) (err error) {
//	kc.clientConf = sarama.NewConfig()
//	kc.client, err = sarama.NewConsumerGroup(conf.KafkaConfig.Hosts, "test001", kc.clientConf)
//	if err != nil {
//		logx.Errorf("KafkaConsumer:Init:%v", err)
//		return err
//	}
//	return nil
//}
//
//func (kc *KafkaConsumer) ConsumeWithCallback(ctx context.Context, topic string, callback func(msg *sarama.ConsumerMessage) error) {
//	consumer := QueueMessageConsumer{
//		ready: make(chan bool),
//	}
//
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for {
//			err := kc.client.Consume(context.TODO(), []string{topic}, &consumer)
//			if err != nil {
//				logx.Errorf("KafkaConsumer:ConsumeWithCallback:%v", err)
//			}
//			// check if context was cancelled, signaling that the consumer should stop
//			if ctx.Err() != nil {
//				return
//			}
//			consumer.ready = make(chan bool)
//		}
//	}()
//
//	<-consumer.ready // Await till the consumer has been set up
//	log.Println("Sarama consumer up and running!...")
//
//	sigterm := make(chan os.Signal, 1)
//	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
//	select {
//	case <-ctx.Done():
//		log.Println("terminating: context cancelled")
//	case <-sigterm:
//		log.Println("terminating: via signal")
//	}
//	wg.Wait()
//	if err := kc.client.Close(); err != nil {
//		log.Panicf("Error closing client: %v", err)
//	}
//}
//
//type QueueMessageConsumer struct {
//	ready chan bool
//}
//
//func (consumer *QueueMessageConsumer) Setup(sarama.ConsumerGroupSession) error {
//	close(consumer.ready)
//	return nil
//}
//
//func (consumer *QueueMessageConsumer) Cleanup(sarama.ConsumerGroupSession) error {
//	return nil
//}
//
//func (consumer *QueueMessageConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
//	for message := range claim.Messages() {
//		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
//		session.MarkMessage(message, "")
//	}
//	session.Commit()
//
//	return nil
//}
