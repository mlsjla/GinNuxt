package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var SettingSet = wire.NewSet(wire.Struct(new(SettingAPI), "*"))

type SettingAPI struct {
	SettingSrv *service.SettingSrv
}

func (a *SettingAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SettingQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.SettingSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *SettingAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SettingSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *SettingAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Setting
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SettingSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *SettingAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Setting
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.SettingSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *SettingAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SettingSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
