package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var AppLogSet = wire.NewSet(wire.Struct(new(AppLogSrv), "*"))

type AppLogSrv struct {
	TransRepo  *dao.TransRepo
	AppLogRepo *dao.AppLogRepo
}

func (a *AppLogSrv) Query(ctx context.Context, params schema.AppLogQueryParam, opts ...schema.AppLogQueryOptions) (*schema.AppLogQueryResult, error) {
	result, err := a.AppLogRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *AppLogSrv) Get(ctx context.Context, id uint64, opts ...schema.AppLogQueryOptions) (*schema.AppLog, error) {
	item, err := a.AppLogRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *AppLogSrv) Create(ctx context.Context, item schema.AppLog) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AppLogRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *AppLogSrv) Update(ctx context.Context, id uint64, item schema.AppLog) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AppLogRepo.Update(ctx, id, item)
	})
}

func (a *AppLogSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.AppLogRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.AppLogRepo.Delete(ctx, id)
	})
}
