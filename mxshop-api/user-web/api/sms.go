package api

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"user-web/forms"

	"user-web/global"
)

func GenerateSmsCode(witdh int) string {
	//生成width长度的短信验证码

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
//新版sdk
func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}
func SendSms(ctx *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	smsCode := GenerateSmsCode(6)
	fmt.Println("smsCode:", smsCode)
	client, err := CreateClient(tea.String(global.ServerConfig.AliSmsInfo.ApiKey), tea.String(global.ServerConfig.AliSmsInfo.ApiSecrect))
	if err != nil {
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName: tea.String("耀斑小记"),
		TemplateCode: tea.String("SMS_230640005"),
		PhoneNumbers: tea.String("13790887214"),
		TemplateParam: tea.String("{\"code\":" + smsCode + "}"),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, err = client.SendSms(sendSmsRequest)
	if err != nil {
		panic(err)
	}
	//将验证码保存起来 - redis
	rdb := redis.NewClient(&redis.Options{
		Addr:fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	// Set 键值对，手机号->短信验证码，过期时间
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":"发送成功",
	})
}

