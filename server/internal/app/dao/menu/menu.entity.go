package menu

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

func GetMenuDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Menu))
}

type SchemaMenu schema.Menu

func (a SchemaMenu) ToMenu() *Menu {
	item := new(Menu)
	structure.Copy(a, item)
	return item
}

type Menu struct {
	util.Model
	Name         string  `gorm:"size:50;index;default:'';not null;"` // 菜单名称
	Icon         *string `gorm:"size:255;"`                          // 菜单图标
	Router       *string `gorm:"size:255;"`                          // 访问路由
	ParentID     *uint64 `gorm:"index;default:0;"`                   // 父级内码
	ParentRouter *string `gorm:"size:512;default:'';"`               // 父级路径
	IsShow       int     `gorm:"index;default:0;"`                   // 是否显示(1:显示 2:隐藏)
	Status       int     `gorm:"index;default:0;"`                   // 状态(1:启用 2:禁用)
	Sequence     int     `gorm:"index;default:0;"`                   // 排序值
	Memo         *string `gorm:"size:1024;"`                         // 备注
	Creator      uint64  `gorm:""`                                   // 创建人
}

func (a Menu) ToSchemaMenu() *schema.Menu {
	item := new(schema.Menu)
	structure.Copy(a, item)
	return item
}

type Menus []*Menu

func (a Menus) ToSchemaMenus() []*schema.Menu {
	list := make([]*schema.Menu, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaMenu()
	}
	return list
}
