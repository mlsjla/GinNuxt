package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var RoleMenuSet = wire.NewSet(wire.Struct(new(RoleMenuSrv), "*"))

type RoleMenuSrv struct {
	TransRepo    *dao.TransRepo
	RoleMenuRepo *dao.RoleMenuRepo
}

func (a *RoleMenuSrv) Query(ctx context.Context, params schema.RoleMenuQueryParam, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenuQueryResult, error) {
	result, err := a.RoleMenuRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *RoleMenuSrv) Get(ctx context.Context, id uint64, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenu, error) {
	item, err := a.RoleMenuRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *RoleMenuSrv) Create(ctx context.Context, item schema.RoleMenu) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RoleMenuRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *RoleMenuSrv) Update(ctx context.Context, id uint64, item schema.RoleMenu) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RoleMenuRepo.Update(ctx, id, item)
	})
}

func (a *RoleMenuSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.RoleMenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RoleMenuRepo.Delete(ctx, id)
	})
}
