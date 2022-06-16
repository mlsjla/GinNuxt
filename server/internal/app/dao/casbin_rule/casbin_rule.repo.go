package casbin_rule

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var CasbinRuleSet = wire.NewSet(wire.Struct(new(CasbinRuleRepo), "*"))

type CasbinRuleRepo struct {
	DB *gorm.DB
}

func (a *CasbinRuleRepo) getQueryOption(opts ...schema.CasbinRuleQueryOptions) schema.CasbinRuleQueryOptions {
	var opt schema.CasbinRuleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *CasbinRuleRepo) Query(ctx context.Context, params schema.CasbinRuleQueryParam, opts ...schema.CasbinRuleQueryOptions) (*schema.CasbinRuleQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetCasbinRuleDB(ctx, a.DB)

	// TODO: Your where condition code here...
	if v := params.V0; v != "" {
		db = db.Where("v0=?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list CasbinRules
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.CasbinRuleQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaCasbinRules(),
	}

	return qr, nil
}

func (a *CasbinRuleRepo) Get(ctx context.Context, id uint64, opts ...schema.CasbinRuleQueryOptions) (*schema.CasbinRule, error) {
	var item CasbinRule
	ok, err := util.FindOne(ctx, GetCasbinRuleDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaCasbinRule(), nil
}

func (a *CasbinRuleRepo) Create(ctx context.Context, item schema.CasbinRule) error {
	eitem := SchemaCasbinRule(item).ToCasbinRule()
	result := GetCasbinRuleDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *CasbinRuleRepo) Update(ctx context.Context, id uint64, item schema.CasbinRule) error {
	eitem := SchemaCasbinRule(item).ToCasbinRule()
	result := GetCasbinRuleDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *CasbinRuleRepo) Delete(ctx context.Context, id uint64) error {
	result := GetCasbinRuleDB(ctx, a.DB).Where("id=?", id).Delete(CasbinRule{})
	return errors.WithStack(result.Error)
}
