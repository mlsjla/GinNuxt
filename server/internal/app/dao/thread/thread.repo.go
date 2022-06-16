package thread

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var ThreadSet = wire.NewSet(wire.Struct(new(ThreadRepo), "*"))

type ThreadRepo struct {
	DB *gorm.DB
}

func (a *ThreadRepo) getQueryOption(opts ...schema.ThreadQueryOptions) schema.ThreadQueryOptions {
	var opt schema.ThreadQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *ThreadRepo) Query(ctx context.Context, params schema.ThreadQueryParam, opts ...schema.ThreadQueryOptions) (*schema.ThreadQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetThreadDB(ctx, a.DB)

	// TODO: Your where condition code here...
	if v := params.CategoryId; v != "" {
		db = db.Where("category_id=?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Threads
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.ThreadQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaThreads(),
	}

	return qr, nil
}

func (a *ThreadRepo) Get(ctx context.Context, id uint64, opts ...schema.ThreadQueryOptions) (*schema.Thread, error) {
	var item Thread
	ok, err := util.FindOne(ctx, GetThreadDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaThread(), nil
}

func (a *ThreadRepo) Create(ctx context.Context, item schema.Thread) error {
	eitem := SchemaThread(item).ToThread()
	result := GetThreadDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *ThreadRepo) Update(ctx context.Context, id uint64, item schema.Thread) error {
	eitem := SchemaThread(item).ToThread()
	result := GetThreadDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *ThreadRepo) Delete(ctx context.Context, id uint64) error {
	result := GetThreadDB(ctx, a.DB).Where("id=?", id).Delete(Thread{})
	return errors.WithStack(result.Error)
}
