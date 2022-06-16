package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var PostSet = wire.NewSet(wire.Struct(new(PostSrv), "*"))

type PostSrv struct {
	TransRepo *dao.TransRepo
	PostRepo  *dao.PostRepo
}

func (a *PostSrv) Query(ctx context.Context, params schema.PostQueryParam, opts ...schema.PostQueryOptions) (*schema.PostQueryResult, error) {
	result, err := a.PostRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *PostSrv) Get(ctx context.Context, id uint64, opts ...schema.PostQueryOptions) (*schema.Post, error) {
	item, err := a.PostRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *PostSrv) Create(ctx context.Context, item schema.Post) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.PostRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *PostSrv) Update(ctx context.Context, id uint64, item schema.Post) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.PostRepo.Update(ctx, id, item)
	})
}

func (a *PostSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.PostRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.PostRepo.Delete(ctx, id)
	})
}
