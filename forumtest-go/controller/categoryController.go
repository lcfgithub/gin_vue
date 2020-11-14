package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"test/forumtest-api/model"
	"test/forumtest-api/repository"
	"test/forumtest-api/responce"
	"test/forumtest-api/vo"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	//DB *gorm.DB
	repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	//db := common.GetDB()
	//db.AutoMigrate(model.Category{})
	//
	//return CategoryController{DB: db}
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		responce.Error(ctx, nil, "分类名称验证错误")
		return
	}
	//var requestCategory model.Category
	//ctx.Bind(&requestCategory)
	//if requestCategory.Name == "" {
	//	responce.Error(ctx, nil, "分类名称不能为空")
	//	return
	//}
	//c.DB.Create(&requestCategory)
	/*
	category := model.Category{Name: requestCategory.Name}
	c.DB.Create(&category)
	*/
	category, err := c.repository.Create(requestCategory.Name)
	if err != nil {
		//responce.Error(ctx, nil, "创建失败")
		panic(err)
		return
	}
	responce.Success(ctx, gin.H{
		"category": category,
	}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//获取body参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		responce.Error(ctx, nil, "分类名称不能为空")
		return
	}
	//获取path参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//var updateCategory model.Category
	//c.DB.First(&updateCategory, categoryId)
	//if updateCategory.ID == 0 {
	//	responce.Error(ctx, nil, "分类不存在")
	//	return
	//}

	updateCategory, err := c.repository.SelectById(categoryId)
	if err != nil {
		responce.Error(ctx, nil, "分类不存在")
		return
	}

	//更新分类
	//c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	category,err := c.repository.Update(updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}
	responce.Success(ctx, gin.H{
		"name": category,
	}, "更新分类成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	//var findCategory model.Category
	//c.DB.First(&findCategory, categoryId)
	//if findCategory.ID == 0 {
	//	responce.Error(ctx, nil, "记录不存在")
	//	return
	//}
	_, err := c.repository.SelectById(categoryId)
	if err != nil {
		responce.Error(ctx, nil, "分类不存在")
		return
	}
	//if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
	//	responce.Error(ctx, nil, "删除失败")
	//	return
	//}
	if err := c.repository.DeleteById(categoryId); err != nil {
		responce.Error(ctx, nil, "删除失败，请重试")
	}
	responce.Success(ctx, nil, "删除成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	//var findCategory model.Category
	//c.DB.First(&findCategory, categoryId)
	//if findCategory.ID == 0 {
	//	responce.Error(ctx, nil, "记录不存在")
	//	return
	//}
	category, err := c.repository.SelectById(categoryId)
	if err != nil {
		responce.Error(ctx, nil, "分类不存在")
		return
	}
	responce.Success(ctx, gin.H{
		"rows": category,
	}, "")
}

