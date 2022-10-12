package service

import (
	"context"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/h2non/filetype"

	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/upload"
)

var AttachmentSet = wire.NewSet(wire.Struct(new(AttachmentSrv), "*"))

type AttachmentSrv struct {
	TransRepo      *dao.TransRepo
	AttachmentRepo *dao.AttachmentRepo
}

func (a *AttachmentSrv) Query(ctx context.Context, params schema.AttachmentQueryParam, opts ...schema.AttachmentQueryOptions) (*schema.AttachmentQueryResult, error) {
	result, err := a.AttachmentRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *AttachmentSrv) Get(ctx context.Context, id uint64, opts ...schema.AttachmentQueryOptions) (*schema.Attachment, error) {
	item, err := a.AttachmentRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *AttachmentSrv) Create(ctx context.Context, item schema.Attachment) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AttachmentRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *AttachmentSrv) Update(ctx context.Context, id uint64, item schema.Attachment) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AttachmentRepo.Update(ctx, id, item)
	})
}

func (a *AttachmentSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.AttachmentRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AttachmentRepo.Delete(ctx, id)
	})
}

//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader string
//@return: err error, file model.ExaFileUploadAndDownload

func (e *AttachmentSrv) UploadFile(ctx *gin.Context, header *multipart.FileHeader) (file schema.Attachment, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		return file, uploadErr
	}
	// TypeId 类型数据ID(post_id,dialog_message_id…)
	// Type 0帖子附件，1帖子图片，2帖子视频，3帖子音频，4消息图片
	// s := strings.Split(header.Filename, ".")

	hander, _ := header.Open()
	head := make([]byte, 1024)
	hander.Read(head)

	kind, _ := filetype.Match(head)
	// if kind == filetype.Unknown {
	// 	return file, err
	// }

	mime := strings.Split(kind.MIME.Value, "/")

	var FileWidth uint64 = 0
	var FileHeight uint64 = 0
	if mime[0] == "image" {
		hander.Seek(0, 0)
		config, _, err := image.DecodeConfig(hander)
		if err != nil {
			return file, err
		}
		if config.Width > 0 {
			FileWidth = uint64(config.Width)
		}
		if config.Height > 0 {
			FileHeight = uint64(config.Height)
		}

	}
	UserId := contextx.FromUserID(ctx.Request.Context())

	f := schema.Attachment{
		ID:         0,
		Uuid:       "",
		UserId:     UserId,
		TypeId:     0,
		Order:      0,
		Type:       0,
		IsRemote:   0,
		Attachment: key,
		FilePath:   filePath,
		FileName:   header.Filename,
		FileSize:   uint64(header.Size),
		FileWidth:  FileWidth,
		FileHeight: FileHeight,
		FileType:   kind.MIME.Value,
		Ip:         ctx.ClientIP(),
	}
	_, err = e.Create(ctx, f)
	return f, err
}
