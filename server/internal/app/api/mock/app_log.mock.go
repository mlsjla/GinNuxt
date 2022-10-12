package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var AppLogSet = wire.NewSet(wire.Struct(new(AppLogMock), "*"))

type AppLogMock struct{}

// @Tags APPLog
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Success 200 {object} schema.ListResult{list=[]schema.AppLog} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/app-logs [get]
func (a *AppLogMock) Query(c *gin.Context) {
}

// @Tags APPLog
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.AppLog
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/app-logs/{id} [get]
func (a *AppLogMock) Get(c *gin.Context) {
}

// @Tags APPLog
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.AppLog true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/app-logs [post]
func (a *AppLogMock) Create(c *gin.Context) {
}

// @Tags APPLog
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.AppLog true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/app-logs/{id} [put]
func (a *AppLogMock) Update(c *gin.Context) {
}

// @Tags APPLog
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/app-logs/{id} [delete]
func (a *AppLogMock) Delete(c *gin.Context) {
}
