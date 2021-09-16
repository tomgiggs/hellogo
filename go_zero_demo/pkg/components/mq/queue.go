package mq

import (
	"context"
)

type ConsumeStatus int

const (
	AckWithSuccess         ConsumeStatus = 0
	AckWithFailed          ConsumeStatus = 1
	AckWithInvalidMessage  ConsumeStatus = 2
	AckWithMaxConsumeTimes ConsumeStatus = 3
)

type MessageResult struct {
	Code      ConsumeStatus `json:"code"`
	MessageId string        `json:"message_id"`
	Message   string        `json:"message"`
}

type Queue interface {
	Init(ctx context.Context)
	ConsumeWithCallback(ctx context.Context, callback func())
	Enqueue(ctx context.Context, key string, message string, args ...interface{}) (ok bool, err error)
	Dequeue(ctx context.Context, key string, args ...interface{}) (message string, tag string, token string, dequeueCount int64, err error)
	AckMsg(ctx context.Context, key string, token string, args ...interface{}) (ok bool, err error)
	BatchEnqueue(ctx context.Context, key string, messages []string, args ...interface{}) (ok bool, err error)
}
