package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RoleMenuSet = wire.NewSet(wire.Struct(new(RoleMenuMock), "*"))

type RoleMenuMock struct{}

// @Tags 菜单权限
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Success 200 {object} schema.ListResult{list=[]schema.RoleMenu} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/role-menus [get]
func (a *RoleMenuMock) Query(c *gin.Context) {
}

// @Tags 菜单权限
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.RoleMenu
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/role-menus/{id} [get]
func (a *RoleMenuMock) Get(c *gin.Context) {
}

// @Tags 菜单权限
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.RoleMenu true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/role-menus [post]
func (a *RoleMenuMock) Create(c *gin.Context) {
}

// @Tags 菜单权限
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.RoleMenu true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/role-menus/{id} [put]
func (a *RoleMenuMock) Update(c *gin.Context) {
}

// @Tags 菜单权限
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/role-menus/{id} [delete]
func (a *RoleMenuMock) Delete(c *gin.Context) {
}
