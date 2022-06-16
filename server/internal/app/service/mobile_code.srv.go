package service

import (
	"context"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var MobileCodeSet = wire.NewSet(wire.Struct(new(MobileCodeSrv), "*"))

type MobileCodeSrv struct {
	TransRepo      *dao.TransRepo
	MobileCodeRepo *dao.MobileCodeRepo
}

func (a *MobileCodeSrv) Query(ctx context.Context, params schema.MobileCodeQueryParam, opts ...schema.MobileCodeQueryOptions) (*schema.MobileCodeQueryResult, error) {
	result, err := a.MobileCodeRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *MobileCodeSrv) Get(ctx context.Context, id uint64, opts ...schema.MobileCodeQueryOptions) (*schema.MobileCode, error) {
	item, err := a.MobileCodeRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *MobileCodeSrv) Create(ctx context.Context, item schema.MobileCode) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.MobileCodeRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *MobileCodeSrv) Update(ctx context.Context, id uint64, item schema.MobileCode) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.MobileCodeRepo.Update(ctx, id, item)
	})
}

func (a *MobileCodeSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.MobileCodeRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.MobileCodeRepo.Delete(ctx, id)
	})
}
