package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
	"user_srv/model"
)

func genMd5(code string) string{
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	//dsn := "root:root@tcp(192.168.0.104:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:4.234.23123aa@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:4.234.23123aa@tcp(47.106.87.191:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,         // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(newPassword)
	for i:=0; i<10; i++{
		user := model.User{
			NickName: fmt.Sprintf("haha%d",i),
			Mobile: fmt.Sprintf("1878222212%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}

	////设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
	////sql才是最重要的，本着这个原则我尽量的给大家看到每个api背后的sql语句是什么
	//
	////定义一个表结构， 将表结构直接生成对应的表 - migrations
	//// 迁移 schema
	_ = db.AutoMigrate(&model.User{}) //此处应该有sql语句

	fmt.Println(genMd5("flare_123456"))
	//将用户的密码变一下 随机字符串+用户密码
	//暴力破解 123456 111111 000000 彩虹表 盐值
	//47f2426c49ebe7fe9b0fed770d5a9573

	// Using custom options
	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("generic password", options)
	//
	////保存在数据库的密码需要盐值才能验证，那为了避免侵入性，一般将盐值和加密后的密码以及使用的算法组合成新的字符串
	//newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println(len(newPassword))
	//fmt.Println(salt, newPassword)
	//passwordInfo := strings.Split(newPassword, "$")
	//fmt.Println(passwordInfo)
	//
	////使用包中的验证函数，验证用户输入的密码是否正确
	//check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(check) // true
}