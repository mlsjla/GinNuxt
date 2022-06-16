package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var RoleMenuSet = wire.NewSet(wire.Struct(new(RoleMenuAPI), "*"))

type RoleMenuAPI struct {
	RoleMenuSrv *service.RoleMenuSrv
}

func (a *RoleMenuAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.RoleMenuQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	fmt.Println("params roleId=>", params.RoleId)
	result, err := a.RoleMenuSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *RoleMenuAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.RoleMenuSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *RoleMenuAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.RoleMenu
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.RoleMenuSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *RoleMenuAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.RoleMenu
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.RoleMenuSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *RoleMenuAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.RoleMenuSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
