package app

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get App db model
func GetAppDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(App))
}

// App
type SchemaApp schema.App

// Convert to App entity
func (a SchemaApp) ToApp() *App {
	item := new(App)
	structure.Copy(a, item)
	return item
}

// App entity
type App struct {
	util.Model
	UserId uint64 `gorm:"index"`         // 创建用户 id
	Type   int    `gorm:""`              // 类型
	Appid  string `gorm:"string;index"`  // Appid
	Data   string `gorm:"type:text"`     // 详情
	Ip     string `gorm:"size:45;index"` // ip

}

// Convert to App schema
func (a App) ToSchemaApp() *schema.App {
	item := new(schema.App)
	structure.Copy(a, item)
	return item
}

// App entity list
type Apps []*App

// Convert to App schema list
func (a Apps) ToSchemaApps() []*schema.App {
	list := make([]*schema.App, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaApp()
	}
	return list
}
