package components

import (
	"gitlab.deepwisdomai.com/infra/go-zero/core/stores/redis"
)

func GetClient(){
	client := redis.NewRedis("","","")
	client.Ping()
}