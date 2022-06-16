package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var UploadSet = wire.NewSet(wire.Struct(new(UploadAPI), "*"))

type UploadAPI struct {
	AttachmentSrv *service.AttachmentSrv
}

// @Tags FileUpload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Router /api/v1/upload/upload [post]
func (u *UploadAPI) UploadFile(c *gin.Context) {
	// _, header, err := c.Request.FormFile("file")
	//设置内存大小
	err := c.Request.ParseMultipartForm(32 << 30)
	if err != nil {
		ginx.ResError(c, err)
	}
	//获取上传的文件组
	files := c.Request.MultipartForm.File["file[]"]
	len := len(files)
	results := make([]schema.Attachment, len)
	for i := 0; i < len; i++ {
		//打开上传文件
		file := files[i]
		result, err := u.AttachmentSrv.UploadFile(c, file)
		if err != nil {
			ginx.ResError(c, err)
			return
		}
		results[i] = result
	}
	ginx.ResSuccess(c, results)
}

// @Tags FileUpload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body schema.Attachment true "传入文件里面id即可"
// @Router /api/v1/upload/deleteFile [post]
func (u *UploadAPI) DeleteFile(c *gin.Context) {
	var file schema.Attachment
	_ = c.ShouldBindJSON(&file)

	if err := u.AttachmentSrv.Delete(c, ginx.ParseParamID(c, "id")); err != nil {
		// if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}

// @Tags FileUpload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body schema.AttachmentQueryParam true "页码, 每页大小"
// @Router /api/v1/upload/getFileList [post]
func (u *UploadAPI) GetFileList(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.AttachmentQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	params.Pagination = true
	_ = c.ShouldBindJSON(&params)
	result, err := u.AttachmentSrv.Query(ctx, params)
	// err, list, total := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		ginx.ResError(c, err)
	} else {
		ginx.ResSuccess(c, result)
	}
}
