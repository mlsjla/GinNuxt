package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var CasbinRuleSet = wire.NewSet(wire.Struct(new(CasbinRuleSrv), "*"))

type CasbinRuleSrv struct {
	TransRepo      *dao.TransRepo
	CasbinRuleRepo *dao.CasbinRuleRepo
}

func (a *CasbinRuleSrv) Query(ctx context.Context, params schema.CasbinRuleQueryParam, opts ...schema.CasbinRuleQueryOptions) (*schema.CasbinRuleQueryResult, error) {
	result, err := a.CasbinRuleRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *CasbinRuleSrv) Get(ctx context.Context, id uint64, opts ...schema.CasbinRuleQueryOptions) (*schema.CasbinRule, error) {
	item, err := a.CasbinRuleRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *CasbinRuleSrv) Create(ctx context.Context, item schema.CasbinRule) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.CasbinRuleRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *CasbinRuleSrv) Update(ctx context.Context, id uint64, item schema.CasbinRule) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.CasbinRuleRepo.Update(ctx, id, item)
	})
}

func (a *CasbinRuleSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.CasbinRuleRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.CasbinRuleRepo.Delete(ctx, id)
	})
}
