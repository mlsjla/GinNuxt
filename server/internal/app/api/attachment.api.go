package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var AttachmentSet = wire.NewSet(wire.Struct(new(AttachmentAPI), "*"))

type AttachmentAPI struct {
	AttachmentSrv *service.AttachmentSrv
}

func (a *AttachmentAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.AttachmentQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.AttachmentSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *AttachmentAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.AttachmentSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *AttachmentAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Attachment
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.AttachmentSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *AttachmentAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Attachment
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.AttachmentSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *AttachmentAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AttachmentSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
