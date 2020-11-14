package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"test/forumtest-api/common"
	"test/forumtest-api/model"
	"test/forumtest-api/responce"
	"test/forumtest-api/vo"
)

type IPostController interface {
	PageList(ctx *gin.Context)
	RestController
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() PostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}
func (p PostController) PageList(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.DefaultPostForm("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultPostForm("pageSize", "20"))
	log.Print(pageNum, pageSize)

	//分页
	var post []model.Post
	p.DB.Select("id,created_at,title").Order("created_at desc").Offset((pageNum-1) * pageSize).Limit(pageSize).Find(&post)

	//总记录数
	var total int64
	p.DB.Model(model.Post{}).Count(&total)

	responce.Success(ctx, gin.H{
		"data": post,
		"total": total,
	}, "OK")

}
func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := ctx.ShouldBind(&requestPost); err != nil {
		responce.Error(ctx, nil, "数据验证错误")
		return
	}
	//获取登录用户
	user, _ := ctx.Get("user")
	//创建Post
	Post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	if err := p.DB.Create(&Post).Error; err != nil {
		panic(err)
		return
	}
	responce.Success(ctx, nil, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := ctx.ShouldBind(&requestPost); err != nil {
		responce.Error(ctx, nil, "数据验证错误")
		return
	}
	//获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	p.DB.Where("id = ?", postId).First(&post)
	if post.UserId == 0 {
		responce.Error(ctx, nil, "文章不存在")
		return
	}
	//获取登录用户
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		responce.Error(ctx, nil, "无权限操作其他作者的文章")
		return
	}
	//更新文章
	if err := p.DB.Model(&post).Updates(model.Post{
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}).Error; err != nil {
		responce.Error(ctx, nil, "更新文章失败")
		return
	}
	responce.Success(ctx, nil, "更新文章成功")

}

func (p PostController) Delete(ctx *gin.Context) {
	//获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	p.DB.Where("id = ?", postId).First(&post)
	if post.UserId == 0 {
		responce.Error(ctx, nil, "文章不存在")
		return
	}
	//获取登录用户
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		responce.Error(ctx, nil, "无权限操作其他作者的文章")
		return
	}
	if err := p.DB.Delete(&post).Error; err != nil {
		responce.Error(ctx, nil, "文章删除失败")
		return
	}
	responce.Success(ctx, nil, "文章删除成功")
}

func (p PostController) Show(ctx *gin.Context) {
	//获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	//p.DB.Preload("Category").Where("id = ?", postId).First(&post)
	p.DB.Where("id = ?", postId).First(&post)
	if post.UserId == 0 {
		responce.Error(ctx, nil, "文章不存在")
		return
	}
	responce.Success(ctx, gin.H{
		"post": post,
	}, "")
}
