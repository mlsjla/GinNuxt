package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var AppSet = wire.NewSet(wire.Struct(new(AppAPI), "*"))

type AppAPI struct {
	AppSrv *service.AppSrv
}

func (a *AppAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.AppQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.AppSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *AppAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.AppSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *AppAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.App
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	UserId := contextx.FromUserID(ctx)
	if UserId == 0 {
		ginx.ResError(c, errors.New("用户数据异常"))
		return
	}
	item.UserId = UserId

	result, err := a.AppSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *AppAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.App
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.AppSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *AppAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AppSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
