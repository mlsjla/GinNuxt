package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var MobileCodeSet = wire.NewSet(wire.Struct(new(MobileCodeMock), "*"))

type MobileCodeMock struct{}

// @Tags 短信发送
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Success 200 {object} schema.ListResult{list=[]schema.MobileCode} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/mobile-codes [get]
func (a *MobileCodeMock) Query(c *gin.Context) {
}

// @Tags 短信发送
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.MobileCode
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/mobile-codes/{id} [get]
func (a *MobileCodeMock) Get(c *gin.Context) {
}

// @Tags 短信发送
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.MobileCode true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/mobile-codes [post]
func (a *MobileCodeMock) Create(c *gin.Context) {
}

// @Tags 短信发送
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.MobileCode true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/mobile-codes/{id} [put]
func (a *MobileCodeMock) Update(c *gin.Context) {
}

// @Tags 短信发送
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/mobile-codes/{id} [delete]
func (a *MobileCodeMock) Delete(c *gin.Context) {
}
