package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
	"github.com/mlsjla/gin-nuxt/server/internal/app/setting"
)

var CasbinRuleSet = wire.NewSet(wire.Struct(new(CasbinRuleAPI), "*"))

type CasbinRuleAPI struct {
	CasbinRuleSrv *service.CasbinRuleSrv
}

func (a *CasbinRuleAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.CasbinRuleQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.CasbinRuleSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *CasbinRuleAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CasbinRuleSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *CasbinRuleAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.CasbinRule
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.CasbinRuleSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *CasbinRuleAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.CasbinRule
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.CasbinRuleSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *CasbinRuleAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.CasbinRuleSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *CasbinRuleAPI) QueryUseApi(c *gin.Context) {
	routers := setting.Routers

	ginx.ResSuccess(c, routers)
}
