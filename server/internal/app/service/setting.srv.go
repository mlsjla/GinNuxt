package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var SettingSet = wire.NewSet(wire.Struct(new(SettingSrv), "*"))

type SettingSrv struct {
	TransRepo   *dao.TransRepo
	SettingRepo *dao.SettingRepo
}

func (a *SettingSrv) Query(ctx context.Context, params schema.SettingQueryParam, opts ...schema.SettingQueryOptions) (*schema.SettingQueryResult, error) {
	result, err := a.SettingRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *SettingSrv) Get(ctx context.Context, id uint64, opts ...schema.SettingQueryOptions) (*schema.Setting, error) {
	item, err := a.SettingRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *SettingSrv) Create(ctx context.Context, item schema.Setting) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.SettingRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *SettingSrv) Update(ctx context.Context, id uint64, item schema.Setting) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.SettingRepo.Update(ctx, id, item)
	})
}

func (a *SettingSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.SettingRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.SettingRepo.Delete(ctx, id)
	})
}
