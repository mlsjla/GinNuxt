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

var ThreadSet = wire.NewSet(wire.Struct(new(ThreadAPI), "*"))

type ThreadAPI struct {
	ThreadSrv *service.ThreadSrv
	PostSrv   *service.PostSrv
}

func (a *ThreadAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ThreadQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.ThreadSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *ThreadAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ThreadSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	var params schema.PostQueryParam
	params.ThreadId = ginx.ParseParamID(c, "id")
	params.IsFirst = 1
	post, err := a.PostSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	if len(post.Data) > 0 {
		item.Content = post.Data[0].Content
	}

	ginx.ResSuccess(c, item)
}

func (a *ThreadAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Thread
	if err := ginx.ParseJSONWidth(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	UserId := contextx.FromUserID(ctx)
	FromUserName := contextx.FromUserName(ctx)
	if UserId == 0 {
		ginx.ResError(c, errors.New("用户数据异常"))
		return
	}

	item.UserId = UserId
	item.Username = FromUserName

	result, err := a.ThreadSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	// 存入post信息
	var postItem schema.Post
	if err := ginx.ParseJSONWidth(c, &postItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	postItem.UserId = UserId
	postItem.ThreadId = result.ID
	postItem.IsFirst = 1
	_, err = a.PostSrv.Create(ctx, postItem)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, result)
}

func (a *ThreadAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Thread
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.ThreadSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	// 修改content

	ginx.ResOK(c)
}

func (a *ThreadAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ThreadSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
