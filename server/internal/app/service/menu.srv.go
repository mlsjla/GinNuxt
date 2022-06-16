package service

import (
	"context"
	"fmt"
	"os"

	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/snowflake"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/yaml"
)

var MenuSet = wire.NewSet(wire.Struct(new(MenuSrv), "*"))

type MenuSrv struct {
	TransRepo *dao.TransRepo
	MenuRepo  *dao.MenuRepo
}

func (a *MenuSrv) InitData(ctx context.Context, dataFile string) error {
	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return nil
	}

	data, err := a.readData(dataFile)
	if err != nil {
		return err
	}

	return a.createMenus(ctx, 0, data)
}

func (a *MenuSrv) readData(name string) (schema.MenuTrees, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data schema.MenuTrees
	d := yaml.NewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	return data, err
}

func (a *MenuSrv) createMenus(ctx context.Context, parentID uint64, list schema.MenuTrees) error {
	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		for _, item := range list {
			sitem := schema.Menu{
				Name:     item.Name,
				Sequence: item.Sequence,
				Icon:     item.Icon,
				Router:   item.Router,
				ParentID: parentID,
				Status:   1,
				IsShow:   1,
				Actions:  item.Actions,
			}
			if v := item.IsShow; v > 0 {
				sitem.IsShow = v
			}

			nsitem, err := a.Create(ctx, sitem)
			if err != nil {
				return err
			}

			if item.Children != nil && len(*item.Children) > 0 {
				err := a.createMenus(ctx, nsitem.ID, *item.Children)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (a *MenuSrv) Query(ctx context.Context, params schema.MenuQueryParam, opts ...schema.MenuQueryOptions) (*schema.MenuQueryResult, error) {
	result, err := a.MenuRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *MenuSrv) Get(ctx context.Context, id uint64, opts ...schema.MenuQueryOptions) (*schema.Menu, error) {
	item, err := a.MenuRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *MenuSrv) checkName(ctx context.Context, item schema.Menu) error {
	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		ParentID: &item.ParentID,
		Name:     item.Name,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("名称不能重复")
	}
	return nil
}

func (a *MenuSrv) Create(ctx context.Context, item schema.Menu) (*schema.IDResult, error) {
	if err := a.checkName(ctx, item); err != nil {
		return nil, err
	}

	ParentRouter, err := a.getParentRouter(ctx, item.ParentID)
	if err != nil {
		return nil, err
	}
	item.ParentRouter = ParentRouter
	item.ID = snowflake.MustID()

	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.MenuRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *MenuSrv) getParentRouter(ctx context.Context, parentID uint64) (string, error) {
	if parentID == 0 {
		return "", nil
	}

	pitem, err := a.MenuRepo.Get(ctx, parentID)
	if err != nil {
		return "", err
	} else if pitem == nil {
		return "", errors.ErrInvalidParent
	}

	return a.joinParentRouter(pitem.ParentRouter, pitem.ID), nil
}

func (a *MenuSrv) joinParentRouter(parent string, id uint64) string {
	if parent != "" {
		parent += "/"
	}

	return fmt.Sprintf("%s%d", parent, id)
}

func (a *MenuSrv) Update(ctx context.Context, id uint64, item schema.Menu) error {
	if id == item.ParentID {
		return errors.ErrInvalidParent
	}

	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Name != item.Name {
		if err := a.checkName(ctx, item); err != nil {
			return err
		}
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	if oldItem.ParentID != item.ParentID {
		ParentRouter, err := a.getParentRouter(ctx, item.ParentID)
		if err != nil {
			return err
		}
		item.ParentRouter = ParentRouter
	} else {
		item.ParentRouter = oldItem.ParentRouter
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {

		err = a.updateChildParentRouter(ctx, *oldItem, item)
		if err != nil {
			return err
		}

		return a.MenuRepo.Update(ctx, id, item)
	})
}

func (a *MenuSrv) updateChildParentRouter(ctx context.Context, oldItem, newItem schema.Menu) error {
	if oldItem.ParentID == newItem.ParentID {
		return nil
	}

	opath := a.joinParentRouter(oldItem.ParentRouter, oldItem.ID)
	result, err := a.MenuRepo.Query(contextx.NewNoTrans(ctx), schema.MenuQueryParam{
		PrefixParentRouter: opath,
	})
	if err != nil {
		return err
	}

	npath := a.joinParentRouter(newItem.ParentRouter, newItem.ID)
	for _, menu := range result.Data {
		err = a.MenuRepo.UpdateParentRouter(ctx, menu.ID, npath+menu.ParentRouter[len(opath):])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *MenuSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		ParentID:        &id,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("forbid delete")
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.MenuRepo.Delete(ctx, id)
	})
}

func (a *MenuSrv) UpdateStatus(ctx context.Context, id uint64, status int) error {
	oldItem, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Status == status {
		return nil
	}

	return a.MenuRepo.UpdateStatus(ctx, id, status)
}
