package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	//consul地址
	cfg.Address = "localhost:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	//check := &api.AgentServiceCheck{
	//	//consul回调的健康检查地址
	//	HTTP:                           fmt.Sprintf("http://%s:%s/health",address,port),
	//	Timeout:                        "5s",
	//	Interval:                       "5s",
	//	DeregisterCriticalServiceAfter: "10s",
	//}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	//registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "120.24.221.188:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}
func FilterSerivice() {
	cfg := api.DefaultConfig()
	cfg.Address = "120.24.221.188:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func main() {
	_ = Register("localhost", 8881, "server", []string{"test", "demo"}, "server")
	//AllServices()
	//FilterSerivice()
	//fmt.Println(fmt.Sprintf(`Service == "%s"`, "user-srv"))
}
