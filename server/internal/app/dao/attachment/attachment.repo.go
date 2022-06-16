package attachment

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var AttachmentSet = wire.NewSet(wire.Struct(new(AttachmentRepo), "*"))

type AttachmentRepo struct {
	DB *gorm.DB
}

func (a *AttachmentRepo) getQueryOption(opts ...schema.AttachmentQueryOptions) schema.AttachmentQueryOptions {
	var opt schema.AttachmentQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *AttachmentRepo) Query(ctx context.Context, params schema.AttachmentQueryParam, opts ...schema.AttachmentQueryOptions) (*schema.AttachmentQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetAttachmentDB(ctx, a.DB)

	// TODO: Your where condition code here...

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Attachments
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.AttachmentQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaAttachments(),
	}

	return qr, nil
}

func (a *AttachmentRepo) Get(ctx context.Context, id uint64, opts ...schema.AttachmentQueryOptions) (*schema.Attachment, error) {
	var item Attachment
	ok, err := util.FindOne(ctx, GetAttachmentDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaAttachment(), nil
}

func (a *AttachmentRepo) Create(ctx context.Context, item schema.Attachment) error {
	eitem := SchemaAttachment(item).ToAttachment()
	result := GetAttachmentDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *AttachmentRepo) Update(ctx context.Context, id uint64, item schema.Attachment) error {
	eitem := SchemaAttachment(item).ToAttachment()
	result := GetAttachmentDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *AttachmentRepo) Delete(ctx context.Context, id uint64) error {
	result := GetAttachmentDB(ctx, a.DB).Where("id=?", id).Delete(Attachment{})
	return errors.WithStack(result.Error)
}
