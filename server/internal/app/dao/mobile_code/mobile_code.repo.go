package mobile_code

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var MobileCodeSet = wire.NewSet(wire.Struct(new(MobileCodeRepo), "*"))

type MobileCodeRepo struct {
	DB *gorm.DB
}

func (a *MobileCodeRepo) getQueryOption(opts ...schema.MobileCodeQueryOptions) schema.MobileCodeQueryOptions {
	var opt schema.MobileCodeQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *MobileCodeRepo) Query(ctx context.Context, params schema.MobileCodeQueryParam, opts ...schema.MobileCodeQueryOptions) (*schema.MobileCodeQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetMobileCodeDB(ctx, a.DB)

	// TODO: Your where condition code here...

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list MobileCodes
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.MobileCodeQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaMobileCodes(),
	}

	return qr, nil
}

func (a *MobileCodeRepo) Get(ctx context.Context, id uint64, opts ...schema.MobileCodeQueryOptions) (*schema.MobileCode, error) {
	var item MobileCode
	ok, err := util.FindOne(ctx, GetMobileCodeDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaMobileCode(), nil
}

func (a *MobileCodeRepo) Create(ctx context.Context, item schema.MobileCode) error {
	eitem := SchemaMobileCode(item).ToMobileCode()
	result := GetMobileCodeDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *MobileCodeRepo) Update(ctx context.Context, id uint64, item schema.MobileCode) error {
	eitem := SchemaMobileCode(item).ToMobileCode()
	result := GetMobileCodeDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *MobileCodeRepo) Delete(ctx context.Context, id uint64) error {
	result := GetMobileCodeDB(ctx, a.DB).Where("id=?", id).Delete(MobileCode{})
	return errors.WithStack(result.Error)
}
