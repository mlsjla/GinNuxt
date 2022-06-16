package role_menu

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var RoleMenuSet = wire.NewSet(wire.Struct(new(RoleMenuRepo), "*"))

type RoleMenuRepo struct {
	DB *gorm.DB
}

func (a *RoleMenuRepo) getQueryOption(opts ...schema.RoleMenuQueryOptions) schema.RoleMenuQueryOptions {
	var opt schema.RoleMenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *RoleMenuRepo) Query(ctx context.Context, params schema.RoleMenuQueryParam, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenuQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetRoleMenuDB(ctx, a.DB)

	// TODO: Your where condition code here...
	if v := params.RoleId; v != "" {
		db = db.Where("role_id=?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list RoleMenus
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.RoleMenuQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaRoleMenus(),
	}

	return qr, nil
}

func (a *RoleMenuRepo) Get(ctx context.Context, id uint64, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenu, error) {
	var item RoleMenu
	ok, err := util.FindOne(ctx, GetRoleMenuDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaRoleMenu(), nil
}

func (a *RoleMenuRepo) Create(ctx context.Context, item schema.RoleMenu) error {
	eitem := SchemaRoleMenu(item).ToRoleMenu()
	result := GetRoleMenuDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *RoleMenuRepo) Update(ctx context.Context, id uint64, item schema.RoleMenu) error {
	eitem := SchemaRoleMenu(item).ToRoleMenu()
	result := GetRoleMenuDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *RoleMenuRepo) Delete(ctx context.Context, id uint64) error {
	result := GetRoleMenuDB(ctx, a.DB).Where("id=?", id).Delete(RoleMenu{})
	return errors.WithStack(result.Error)
}
