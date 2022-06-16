package setting

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var SettingSet = wire.NewSet(wire.Struct(new(SettingRepo), "*"))

type SettingRepo struct {
	DB *gorm.DB
}

func (a *SettingRepo) getQueryOption(opts ...schema.SettingQueryOptions) schema.SettingQueryOptions {
	var opt schema.SettingQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *SettingRepo) Query(ctx context.Context, params schema.SettingQueryParam, opts ...schema.SettingQueryOptions) (*schema.SettingQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetSettingDB(ctx, a.DB)

	// TODO: Your where condition code here...
	if v := params.Key; v != "" {
		db = db.Where("`key`=?", v)
	}
	if v := params.Tag; v != "" {
		db = db.Where("tag=?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Settings
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SettingQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaSettings(),
	}

	return qr, nil
}

func (a *SettingRepo) Get(ctx context.Context, id uint64, opts ...schema.SettingQueryOptions) (*schema.Setting, error) {
	var item Setting
	ok, err := util.FindOne(ctx, GetSettingDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaSetting(), nil
}

func (a *SettingRepo) Create(ctx context.Context, item schema.Setting) error {
	eitem := SchemaSetting(item).ToSetting()
	result := GetSettingDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *SettingRepo) Update(ctx context.Context, id uint64, item schema.Setting) error {
	eitem := SchemaSetting(item).ToSetting()
	result := GetSettingDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *SettingRepo) Delete(ctx context.Context, id uint64) error {
	result := GetSettingDB(ctx, a.DB).Where("id=?", id).Delete(Setting{})
	return errors.WithStack(result.Error)
}
