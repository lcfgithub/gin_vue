package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"test/forumtest-api/common"
	"test/forumtest-api/dto"
	"test/forumtest-api/model"
	"test/forumtest-api/responce"
	"test/forumtest-api/until"
)

func Login(ctx *gin.Context)  {
	DB :=common.GetDB()
	var params = model.User{}
	ctx.Bind(&params)
	telphone := params.Telphone
	password := params.Password
	//数据验证
	if len(telphone) != 11 {
		responce.Error(ctx, nil, "手机号必须是11位")
		return
	}
	if len(password) < 6 {
		responce.Error(ctx, nil, "密码不能小于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telphone = ?", telphone).First(&user)
	if user.ID == 0 {
		responce.Error(ctx, nil, "该手机号还没有注册")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		responce.Error(ctx, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.GetToken(user)
	if err != nil {
		responce.Error(ctx, nil, "生成token错误")
		log.Printf("token generate error: %v", err)
		return
	}
	retToken, Claims, err := common.ParseToken(token);
	if err != nil {
		responce.Error(ctx, nil, "解析token错误")
		log.Printf("token parseToken error: %v", err)
		return
	}
	responce.Success(ctx, gin.H{
		"token": token,
		"retToken": retToken,
		"Claims": Claims,
	}, "登录成功")
}

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	//使用map接收数据
	//var requestMap = make(map[string]string)
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	//var requestUser = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)

	var params = model.User{}
	ctx.Bind(&params)
	name := params.Name
	telphone := params.Telphone
	password := params.Password

	//telphone := ctx.PostForm("telphone")
	//数据验证
	if len(telphone) != 11 {
		responce.Error(ctx, nil, "手机号必须是11位" )
		return
	}
	if len(password) < 6 {
		responce.Error(ctx, nil, "密码不能小于6位")
		return
	}
	if len(name) == 0 {
		name = until.GetRandomString(6)
	}
	//判断手机号是否存在
	if IsTelhponeExists(DB, telphone) {
		responce.Error(ctx, nil, "该手机号已注册")
		return
	}
	bcrypt_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		responce.Error(ctx, nil, "加密失败")
		return
	}
	//创建用户
	newUser := model.User{
		Model:    gorm.Model{},
		Name:     name,
		Telphone: telphone,
		Password: string(bcrypt_password),
	}
	DB.Create(&newUser)
	//发放token
	token, err := common.GetToken(newUser)
	if err != nil {
		responce.Error(ctx, nil, "生成token错误")
		log.Printf("token generate error: %v", err)
		return
	}
	//返回结果
	responce.Success(ctx,gin.H{
		"token": token,
	}, "注册成功")
}

func IsTelhponeExists(db *gorm.DB, telphone string) bool  {
	var user model.User
	db.Where("telphone = ?", telphone).First(&user)
	if user.Model.ID != 0 {
		return true
	}
	return false
}

func Info(ctx *gin.Context)  {
	user,_ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"user": dto.UserDto(user.(model.User)),
	})
}