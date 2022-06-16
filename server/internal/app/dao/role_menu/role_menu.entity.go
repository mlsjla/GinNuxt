package role_menu

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get RoleMenu db model
func GetRoleMenuDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(RoleMenu))
}

// RoleMenu
type SchemaRoleMenu schema.RoleMenu

// Convert to RoleMenu entity
func (a SchemaRoleMenu) ToRoleMenu() *RoleMenu {
	item := new(RoleMenu)
	structure.Copy(a, item)
	return item
}

// RoleMenu entity
type RoleMenu struct {
	util.Model
	RoleId uint64 `gorm:"id,string"` // 角色ID
	MenuId uint64 `gorm:"id,string"` // 菜单ID

}

// Convert to RoleMenu schema
func (a RoleMenu) ToSchemaRoleMenu() *schema.RoleMenu {
	item := new(schema.RoleMenu)
	structure.Copy(a, item)
	return item
}

// RoleMenu entity list
type RoleMenus []*RoleMenu

// Convert to RoleMenu schema list
func (a RoleMenus) ToSchemaRoleMenus() []*schema.RoleMenu {
	list := make([]*schema.RoleMenu, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaRoleMenu()
	}
	return list
}
