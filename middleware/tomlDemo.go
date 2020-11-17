package middleware

import "fmt"
import "github.com/BurntSushi/toml"


type Consul struct {
	Node []string
}

type Access struct {
	Port int32
}

type Conf struct {
	Region string
	Consul *Consul
	Access *Access
}

func Init() *Conf {
	return &Conf{
		Region: "",
		Consul: &Consul{
			Node: []string{},
		},
		Access: &Access{
			Port: 0,
		},
	}
}

func TomlDemo(){
	config := Init()
	_, err := toml.DecodeFile("./middleware/accessConf.toml", &config)
	if err != nil{
		panic(err)
	}
	fmt.Println(config.Region,config.Access)
}