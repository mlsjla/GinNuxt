package post

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var PostSet = wire.NewSet(wire.Struct(new(PostRepo), "*"))

type PostRepo struct {
	DB *gorm.DB
}

func (a *PostRepo) getQueryOption(opts ...schema.PostQueryOptions) schema.PostQueryOptions {
	var opt schema.PostQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *PostRepo) Query(ctx context.Context, params schema.PostQueryParam, opts ...schema.PostQueryOptions) (*schema.PostQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetPostDB(ctx, a.DB)

	// TODO: Your where condition code here...
	if v := params.ThreadId; v > 0 {
		db = db.Where("thread_id=?", v)
	}
	if v := params.IsFirst; v > 0 {
		db = db.Where("is_first=?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Posts
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.PostQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaPosts(),
	}

	return qr, nil
}

func (a *PostRepo) Get(ctx context.Context, id uint64, opts ...schema.PostQueryOptions) (*schema.Post, error) {
	var item Post
	ok, err := util.FindOne(ctx, GetPostDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaPost(), nil
}

func (a *PostRepo) Create(ctx context.Context, item schema.Post) error {
	eitem := SchemaPost(item).ToPost()
	result := GetPostDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *PostRepo) Update(ctx context.Context, id uint64, item schema.Post) error {
	eitem := SchemaPost(item).ToPost()
	result := GetPostDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *PostRepo) Delete(ctx context.Context, id uint64) error {
	result := GetPostDB(ctx, a.DB).Where("id=?", id).Delete(Post{})
	return errors.WithStack(result.Error)
}
