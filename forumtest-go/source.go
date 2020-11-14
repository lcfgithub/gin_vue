package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name string `gorm:"varchar(20) not null"`
	Telphone string `gorm:"char(11) not null unique"`
	Password string `gorm:"size:255 not null"`
}

func source()  {
	db := InitDB()
	db.Debug()
	//db.AutoMigrate(&User{})
	r := gin.Default();
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		telphone := ctx.PostForm("telphone")
		password := ctx.PostForm("password")
		//数据验证
		if len(telphone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": len(telphone),
				"msg" : "手机号必须是11位",
			})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg" : "密码不能小于6位",
			})
			return
		}
		if len(name) == 0 {
			name = GetRandomString(6)
		}
		//判断手机号是否存在
		if IsTelhponeExists(db, telphone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg" : "该手机号已注册",
			})
			return
		}
		//创建用户
		newUser := User{
			Model:    gorm.Model{},
			Name:     name,
			Telphone: telphone,
			Password: password,
		}
		db.Create(&newUser)
		//返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run(":8080"))
}

func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, letters[r.Intn(len(letters))])
	}
	return string(result)
}

func IsTelhponeExists(db *gorm.DB, telphone string) bool  {
	var user User
	db.Where("telphone = ?", telphone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/forumtest?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                        // default size for string fields
		DisableDatetimePrecision:  true,                                                                       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                      // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Println("database connect fail err:", err.Error())
	}
	return db
}