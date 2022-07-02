package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"user_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

//Init 初始化连接
func Init(){
	var err error
	//复制全局变量
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

//TestGetUserList 测试list接口，checkPassword接口
func TestGetUserList(){
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	//使用列表中的user
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)
		//调用顺便测试校验密码接口，
		checkRsp, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

func TestCreateUser(){
	for i := 0; i<10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("bobby%d",i),
			Mobile: fmt.Sprintf("1878222222%d",i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()
	//TestCreateUser()
	TestGetUserList()

	conn.Close()
}