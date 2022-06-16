package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SettingSet = wire.NewSet(wire.Struct(new(SettingMock), "*"))

type SettingMock struct{}

// @Tags 系统设置
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param key query string false "key" default("")
// @Param tag query string false "tag" default("")
// @Success 200 {object} schema.ListResult{list=[]schema.Setting} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/settings [get]
func (a *SettingMock) Query(c *gin.Context) {
}

// @Tags 系统设置
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.Setting
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/settings/{id} [get]
func (a *SettingMock) Get(c *gin.Context) {
}

// @Tags 系统设置
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.Setting true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/settings [post]
func (a *SettingMock) Create(c *gin.Context) {
}

// @Tags 系统设置
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.Setting true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/settings/{id} [put]
func (a *SettingMock) Update(c *gin.Context) {
}

// @Tags 系统设置
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/settings/{id} [delete]
func (a *SettingMock) Delete(c *gin.Context) {
}
