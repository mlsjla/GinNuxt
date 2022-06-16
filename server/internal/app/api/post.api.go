package api

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var PostSet = wire.NewSet(wire.Struct(new(PostAPI), "*"))

type PostAPI struct {
	PostSrv *service.PostSrv
}

func (a *PostAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PostQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	fmt.Println("params", params)
	params.Pagination = true
	result, err := a.PostSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	fmt.Println("result.Data", result.Data)
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *PostAPI) QueryTree(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PostQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.PostSrv.Query(ctx, params, schema.PostQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("sort", schema.OrderByDESC)),
	})
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResList(c, result.Data.ToTree())
}

func (a *PostAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PostSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *PostAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Post
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
	result, err := a.PostSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *PostAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Post
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.PostSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *PostAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PostSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
