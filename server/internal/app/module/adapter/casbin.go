package adapter

import (
	"context"
	"fmt"

	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/logger"
)

var _ persist.Adapter = (*CasbinAdapter)(nil)

var CasbinAdapterSet = wire.NewSet(wire.Struct(new(CasbinAdapter), "*"), wire.Bind(new(persist.Adapter), new(*CasbinAdapter)))

type CasbinAdapter struct {
	RoleRepo       *dao.RoleRepo
	UserRepo       *dao.UserRepo
	UserRoleRepo   *dao.UserRoleRepo
	CasbinRuleRepo *dao.CasbinRuleRepo
}

func (a *CasbinAdapter) Query(ctx context.Context, params schema.UserRoleQueryParam, opts ...schema.UserRoleQueryOptions) (*schema.UserRoleQueryResult, error) {
	return a.UserRoleRepo.Query(ctx, params)
}

// Loads all policy rules from the storage.
func (a *CasbinAdapter) LoadPolicy(model casbinModel.Model) error {
	ctx := context.Background()
	err := a.loadDbPolicy(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("Load casbin role policy error: %s", err.Error())
		return err
	}

	return nil
}

// Load role policy (p,role_id,path,method)
func (a *CasbinAdapter) loadDbPolicy(ctx context.Context, m casbinModel.Model) error {
	result, err := a.CasbinRuleRepo.Query(ctx, schema.CasbinRuleQueryParam{})
	if err != nil {
		return err
	}
	for _, item := range result.Data {
		line := fmt.Sprintf("%s,%s,%s,%s", "p", item.V0, item.V1, item.V2)
		persist.LoadPolicyLine(line, m)
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *CasbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
