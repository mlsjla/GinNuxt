package casbin_rule

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get CasbinRule db model
func GetCasbinRuleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(CasbinRule))
}

// CasbinRule
type SchemaCasbinRule schema.CasbinRule

// Convert to CasbinRule entity
func (a SchemaCasbinRule) ToCasbinRule() *CasbinRule {
	item := new(CasbinRule)
	structure.Copy(a, item)
	return item
}

// CasbinRule entity
type CasbinRule struct {
	util.Model
	PType string `gorm:"size:100"` // 规则类型
	V0    string `gorm:"size:100"` // 角色ID
	V1    string `gorm:"size:100"` // api路径
	V2    string `gorm:"size:100"` // api访问方法
	V3    string `gorm:"size:100"`
	V4    string `gorm:"size:100"`
	V5    string `gorm:"size:100"`
}

// Convert to CasbinRule schema
func (a CasbinRule) ToSchemaCasbinRule() *schema.CasbinRule {
	item := new(schema.CasbinRule)
	structure.Copy(a, item)
	return item
}

// CasbinRule entity list
type CasbinRules []*CasbinRule

// Convert to CasbinRule schema list
func (a CasbinRules) ToSchemaCasbinRules() []*schema.CasbinRule {
	list := make([]*schema.CasbinRule, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaCasbinRule()
	}
	return list
}
