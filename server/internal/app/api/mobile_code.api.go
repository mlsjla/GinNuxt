package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var MobileCodeSet = wire.NewSet(wire.Struct(new(MobileCodeAPI), "*"))

type MobileCodeAPI struct {
	MobileCodeSrv *service.MobileCodeSrv
}

func (a *MobileCodeAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.MobileCodeQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.MobileCodeSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *MobileCodeAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.MobileCodeSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *MobileCodeAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.MobileCode
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.MobileCodeSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *MobileCodeAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.MobileCode
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.MobileCodeSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *MobileCodeAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.MobileCodeSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
