package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	//str := ""
	str := "name: user-web\nhost: 192.168.0.103\ntags: \n  - mxshop\n  - imooc\n  - bobby\n  - user\n  - web\nport: 8021\nuser_srv: \n  host: 192.168.0.103\n  port: 50051\n  name: user-srv\njwt: \n  key: 5$!UEmvB#nRB@Iwab#Sy!zofKEOGLRtE\nsms: \n  key: LTAI4FzGkwzJyrKfCex9kpP1\n  secrect: r5aWSxybxkuT4ROcwpRqqusXtcwxt5\n  expire: 300\nredis: \n  host: 192.168.0.104\n  port: 6379\nconsul: \n  host: 192.168.0.104\n  port: 8500"
	//str := "name: user-web\nhost: localhost\ntag: \n  - mxshop\n  - imooc\n  - bobby\n  - user\n  - web\nport: 8021\nuser_srv: \n  host: 120.24.221.188\n  port: 52001\n  name: user-srv\njwt: \n  key: Go2Flare12138\nsms: \n  key: LTAI5tQcYeruM4ouBix44XQZ\n  secrect: WKaxXlmzyItvSj4TIBmmcfKHDt7sid\nredis: \n  host: 47.106.87.191\n  port: 6379\n  expire: 3000\nconsul: \n  host: 120.24.221.188\n  port: 8500"
	//str := "{\n  \"name\": \"user-web\",\n  \"host\": \"localhost\",\n  \"tag\": [\n    \"mxshop\",\n    \"imooc\",\n    \"bobby\",\n    \"user\",\n    \"web\"\n  ],\n  \"port\": 8021,\n  \"user_srv\": {\n    \"name\": \"user-srv\"\n  },\n  \"jwt\": {\n    \"key\": \"Go2Flare12138\"\n  },\n  \"sms\": {\n    \"key\": \"LTAI5tQcYeruM4ouBix44XQZ\",\n    \"secrect\": \"WKaxXlmzyItvSj4TIBmmcfKHDt7sid\"\n  },\n  \"redis\": {\n    \"host\": \"47.106.87.191\",\n    \"port\": 6379,\n    \"expire\": 3000\n  },\n  \"consul\": {\n    \"host\": \"120.24.221.188\",\n    \"port\": 8500\n  }\n}"
	res := ServerConfig{}
	err := json.Unmarshal([]byte(str), &res)

	fmt.Println(err, res)
}
