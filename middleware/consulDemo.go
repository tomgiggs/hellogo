package middleware

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

func ServiceDiscover() {
	var lastIndex uint64
	config := api.DefaultConfig()
	config.Address = "172.26.43.219:8500"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return
	}
	services, metainfo, err := client.Health().Service("go-consul-redis01", "", true, &api.QueryOptions{
		WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})
	if err != nil {
		logrus.Warn("error retrieving instances from Consul: %v", err)
	}
	lastIndex = metainfo.LastIndex

	addrs := map[string]struct{}{}
	for _, service := range services {
		fmt.Println("node.Address:", service.Service.Address, "node.Port:", service.Service.Port)
		addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}
}

func RegisterHttp()  {
	// 创建连接consul服务配置
	config := api.DefaultConfig()
	config.Address = "172.26.43.219:8500"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(api.AgentServiceRegistration)
	registration.ID = "nginxdemo"
	registration.Name = "go-consul-nginxdemo01"
	registration.Port = 8080
	registration.Tags = []string{"go-consul-nginxdemo01"}
	registration.Address = "172.26.43.219"

	// 增加consul健康检查回调函数
	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}
func RegisterMysql()  {
	// 创建连接consul服务配置
	config := api.DefaultConfig()
	config.Address = "172.26.43.219:8500"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(api.AgentServiceRegistration)
	registration.ID = "mysql"
	registration.Name = "go-consul-mysql01"
	registration.Port = 3306
	registration.Tags = []string{"go-consul-mysql01"}
	registration.Address = "192.168.233.3"

	// 增加consul健康检查回调函数
	check := new(api.AgentServiceCheck)
	check.TCP = "192.168.233.3:3306"
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}
func RegisterRedis()  {
	// 创建连接consul服务配置
	config := api.DefaultConfig()
	config.Address = "172.26.43.219:8500"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(api.AgentServiceRegistration)
	registration.ID = "redis02"//节点id
	registration.Name = "go-consul-redis01"//服务名称
	registration.Port = 46379 //节点端口
	registration.Tags = []string{"go-consul-redis02"}
	registration.Address = "172.26.43.219"//节点ip

	// 增加consul健康检查回调函数
	check := new(api.AgentServiceCheck)
	check.TCP = "172.26.43.219:46379"
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}
