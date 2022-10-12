package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(wire.Struct(new(AppMock), "*"))

type AppMock struct{}

// @Tags APP
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Success 200 {object} schema.ListResult{list=[]schema.App} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/apps [get]
func (a *AppMock) Query(c *gin.Context) {
}

// @Tags APP
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.App
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/apps/{id} [get]
func (a *AppMock) Get(c *gin.Context) {
}

// @Tags APP
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.App true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/apps [post]
func (a *AppMock) Create(c *gin.Context) {
}

// @Tags APP
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.App true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/apps/{id} [put]
func (a *AppMock) Update(c *gin.Context) {
}

// @Tags APP
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/apps/{id} [delete]
func (a *AppMock) Delete(c *gin.Context) {
}
