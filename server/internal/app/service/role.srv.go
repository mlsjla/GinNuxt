package service

import (
	"context"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleSrv), "*"))

type RoleSrv struct {
	Enforcer  *casbin.SyncedEnforcer
	TransRepo *dao.TransRepo
	RoleRepo  *dao.RoleRepo
	UserRepo  *dao.UserRepo
}

func (a *RoleSrv) Query(ctx context.Context, params schema.RoleQueryParam, opts ...schema.RoleQueryOptions) (*schema.RoleQueryResult, error) {
	return a.RoleRepo.Query(ctx, params, opts...)
}

func (a *RoleSrv) Get(ctx context.Context, id uint64, opts ...schema.RoleQueryOptions) (*schema.Role, error) {
	item, err := a.RoleRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

func (a *RoleSrv) Create(ctx context.Context, item schema.Role) (*schema.IDResult, error) {
	err := a.checkName(ctx, item)
	if err != nil {
		return nil, err
	}

	item.ID = snowflake.MustID()
	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RoleRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}
	return schema.NewIDResult(item.ID), nil
}

func (a *RoleSrv) checkName(ctx context.Context, item schema.Role) error {
	result, err := a.RoleRepo.Query(ctx, schema.RoleQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		Name:            item.Name,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("名称不允许重复")
	}
	return nil
}

func (a *RoleSrv) Update(ctx context.Context, id uint64, item schema.Role) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Name != item.Name {
		err := a.checkName(ctx, item)
		if err != nil {
			return err
		}
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt
	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RoleRepo.Update(ctx, id, item)
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *RoleSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.RoleRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	userResult, err := a.UserRepo.Query(ctx, schema.UserQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		RoleIDs:         []uint64{id},
	})
	if err != nil {
		return err
	} else if userResult.PageResult.Total > 0 {
		return errors.New400Response("不允许删除已经存在用户的角色")
	}

	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RoleRepo.Delete(ctx, id)
	})
	if err != nil {
		return err
	}

	a.Enforcer.DeleteRole(strconv.FormatUint(id, 10))

	return nil
}

func (a *RoleSrv) UpdateStatus(ctx context.Context, id uint64, status int) error {
	oldItem, err := a.RoleRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Status == status {
		return nil
	}

	err = a.RoleRepo.UpdateStatus(ctx, id, status)
	if err != nil {
		return err
	}
	return nil
}
