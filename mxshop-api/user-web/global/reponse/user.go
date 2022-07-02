package reponse

import (
	"fmt"
	"time"
)

type JsonTime time.Time

// MarshalJSON 用time的同类型实现 Marshaler接口，转json时会自动使用该接口的方法
func (j JsonTime) MarshalJSON() ([]byte, error) {
	//格式化时间
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 15:04:05"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"name"`
	//Birthday string `json:"birthday"`//我们只需要用结构体接收json的信息，不用使用，所以能解析为非string的time更好
	//，但转成指定格式得实现接口
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
}
