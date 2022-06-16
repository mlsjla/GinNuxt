package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var CategorySet = wire.NewSet(wire.Struct(new(CategoryMock), "*"))

type CategoryMock struct{}

// @Tags 文章分类
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Success 200 {object} schema.ListResult{list=[]schema.Category} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/categories [get]
func (a *CategoryMock) Query(c *gin.Context) {
}

// @Tags 文章分类
// @Summary 查询分类树
// @Security ApiKeyAuth
// @Param parentid query int false "父级ID"
// @Success 200 {object} schema.ListResult{list=[]schema.CategoryTree} "查询结果"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/Categorys.tree [get]
func (a *CategoryMock) QueryTree(c *gin.Context) {
}

// @Tags 文章分类
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.Category
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/categories/{id} [get]
func (a *CategoryMock) Get(c *gin.Context) {
}

// @Tags 文章分类
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.Category true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/categories [post]
func (a *CategoryMock) Create(c *gin.Context) {
}

// @Tags 文章分类
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.Category true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/categories/{id} [put]
func (a *CategoryMock) Update(c *gin.Context) {
}

// @Tags 文章分类
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/categories/{id} [delete]
func (a *CategoryMock) Delete(c *gin.Context) {
}
