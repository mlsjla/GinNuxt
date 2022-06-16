package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var ThreadSet = wire.NewSet(wire.Struct(new(ThreadSrv), "*"))

type ThreadSrv struct {
	TransRepo  *dao.TransRepo
	ThreadRepo *dao.ThreadRepo
}

func (a *ThreadSrv) Query(ctx context.Context, params schema.ThreadQueryParam, opts ...schema.ThreadQueryOptions) (*schema.ThreadQueryResult, error) {
	result, err := a.ThreadRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *ThreadSrv) Get(ctx context.Context, id uint64, opts ...schema.ThreadQueryOptions) (*schema.Thread, error) {
	item, err := a.ThreadRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *ThreadSrv) Create(ctx context.Context, item schema.Thread) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()
	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.ThreadRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *ThreadSrv) Update(ctx context.Context, id uint64, item schema.Thread) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.ThreadRepo.Update(ctx, id, item)
	})
}

func (a *ThreadSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.ThreadRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.ThreadRepo.Delete(ctx, id)
	})
}
