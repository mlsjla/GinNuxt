package app_log

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var AppLogSet = wire.NewSet(wire.Struct(new(AppLogRepo), "*"))

type AppLogRepo struct {
	DB *gorm.DB
}

func (a *AppLogRepo) getQueryOption(opts ...schema.AppLogQueryOptions) schema.AppLogQueryOptions {
	var opt schema.AppLogQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *AppLogRepo) Query(ctx context.Context, params schema.AppLogQueryParam, opts ...schema.AppLogQueryOptions) (*schema.AppLogQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetAppLogDB(ctx, a.DB)

	// TODO: Your where condition code here...

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list AppLogs
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.AppLogQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaAppLogs(),
	}

	return qr, nil
}

func (a *AppLogRepo) Get(ctx context.Context, id uint64, opts ...schema.AppLogQueryOptions) (*schema.AppLog, error) {
	var item AppLog
	ok, err := util.FindOne(ctx, GetAppLogDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaAppLog(), nil
}

func (a *AppLogRepo) Create(ctx context.Context, item schema.AppLog) error {
	eitem := SchemaAppLog(item).ToAppLog()
	result := GetAppLogDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *AppLogRepo) Update(ctx context.Context, id uint64, item schema.AppLog) error {
	eitem := SchemaAppLog(item).ToAppLog()
	result := GetAppLogDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *AppLogRepo) Delete(ctx context.Context, id uint64) error {
	result := GetAppLogDB(ctx, a.DB).Where("id=?", id).Delete(AppLog{})
	return errors.WithStack(result.Error)
}
