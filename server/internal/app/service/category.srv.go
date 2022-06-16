package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var CategorySet = wire.NewSet(wire.Struct(new(CategorySrv), "*"))

type CategorySrv struct {
	TransRepo    *dao.TransRepo
	CategoryRepo *dao.CategoryRepo
}

func (a *CategorySrv) Query(ctx context.Context, params schema.CategoryQueryParam, opts ...schema.CategoryQueryOptions) (*schema.CategoryQueryResult, error) {
	result, err := a.CategoryRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *CategorySrv) Get(ctx context.Context, id uint64, opts ...schema.CategoryQueryOptions) (*schema.Category, error) {
	item, err := a.CategoryRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *CategorySrv) Create(ctx context.Context, item schema.Category) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.CategoryRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *CategorySrv) Update(ctx context.Context, id uint64, item schema.Category) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.CategoryRepo.Update(ctx, id, item)
	})
}

func (a *CategorySrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.CategoryRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.CategoryRepo.Delete(ctx, id)
	})
}
