package app

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var AppSet = wire.NewSet(wire.Struct(new(AppRepo), "*"))

type AppRepo struct {
	DB *gorm.DB
}

func (a *AppRepo) getQueryOption(opts ...schema.AppQueryOptions) schema.AppQueryOptions {
	var opt schema.AppQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *AppRepo) Query(ctx context.Context, params schema.AppQueryParam, opts ...schema.AppQueryOptions) (*schema.AppQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetAppDB(ctx, a.DB)

	// TODO: Your where condition code here...

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Apps
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.AppQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaApps(),
	}

	return qr, nil
}

func (a *AppRepo) Get(ctx context.Context, id uint64, opts ...schema.AppQueryOptions) (*schema.App, error) {
	var item App
	ok, err := util.FindOne(ctx, GetAppDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaApp(), nil
}

func (a *AppRepo) Create(ctx context.Context, item schema.App) error {
	eitem := SchemaApp(item).ToApp()
	result := GetAppDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *AppRepo) Update(ctx context.Context, id uint64, item schema.App) error {
	eitem := SchemaApp(item).ToApp()
	result := GetAppDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *AppRepo) Delete(ctx context.Context, id uint64) error {
	result := GetAppDB(ctx, a.DB).Where("id=?", id).Delete(App{})
	return errors.WithStack(result.Error)
}
