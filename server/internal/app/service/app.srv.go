package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var AppSet = wire.NewSet(wire.Struct(new(AppSrv), "*"))

type AppSrv struct {
	TransRepo *dao.TransRepo
	AppRepo   *dao.AppRepo
}

func (a *AppSrv) Query(ctx context.Context, params schema.AppQueryParam, opts ...schema.AppQueryOptions) (*schema.AppQueryResult, error) {
	result, err := a.AppRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *AppSrv) Get(ctx context.Context, id uint64, opts ...schema.AppQueryOptions) (*schema.App, error) {
	item, err := a.AppRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *AppSrv) Create(ctx context.Context, item schema.App) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AppRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *AppSrv) Update(ctx context.Context, id uint64, item schema.App) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AppRepo.Update(ctx, id, item)
	})
}

func (a *AppSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.AppRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AppRepo.Delete(ctx, id)
	})
}
