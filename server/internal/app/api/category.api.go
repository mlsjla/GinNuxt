package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var CategorySet = wire.NewSet(wire.Struct(new(CategoryAPI), "*"))

type CategoryAPI struct {
	CategorySrv *service.CategorySrv
}

func (a *CategoryAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.CategoryQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.CategorySrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *CategoryAPI) QueryTree(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.CategoryQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.CategorySrv.Query(ctx, params, schema.CategoryQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("sort", schema.OrderByDESC)),
	})
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResList(c, result.Data.ToTree())
}

func (a *CategoryAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CategorySrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *CategoryAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Category
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.CategorySrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *CategoryAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Category
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.CategorySrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *CategoryAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.CategorySrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
