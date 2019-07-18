package basic_grammar

//import (
//	"fmt"
//	"github.com/confluentinc/confluent-kafka-go/kafka"
//)
//
//func KafkaTest() {
//
//	c, err := kafka.NewConsumer(&kafka.ConfigMap{
//		"bootstrap.servers": "localhost",
//		"group.id":          "test001",
//		"auto.offset.reset": "earliest",
//	})
//
//	if err != nil {
//		panic(err)
//	}
//
//	c.SubscribeTopics([]string{"topic_demo", "^aRegex.*[Tt]opic"}, nil)
//
//	for {
//		msg, err := c.ReadMessage(-1)
//		if err == nil {
//			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
//		} else {
//			// The client will automatically try to recover from all errors.
//			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
//		}
//	}
//
//	c.Close()
//}
