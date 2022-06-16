package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var CasbinRuleSet = wire.NewSet(wire.Struct(new(CasbinRuleMock), "*"))

type CasbinRuleMock struct{}

// @Tags 权限设置
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Success 200 {object} schema.ListResult{list=[]schema.CasbinRule} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/casbin-rules [get]
func (a *CasbinRuleMock) Query(c *gin.Context) {
}

// @Tags 权限设置
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.CasbinRule
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/casbin-rules/{id} [get]
func (a *CasbinRuleMock) Get(c *gin.Context) {
}

// @Tags 权限设置
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.CasbinRule true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/casbin-rules [post]
func (a *CasbinRuleMock) Create(c *gin.Context) {
}

// @Tags 权限设置
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.CasbinRule true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/casbin-rules/{id} [put]
func (a *CasbinRuleMock) Update(c *gin.Context) {
}

// @Tags 权限设置
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/casbin-rules/{id} [delete]
func (a *CasbinRuleMock) Delete(c *gin.Context) {
}
