package app_log

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get AppLog db model
func GetAppLogDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(AppLog))
}

// AppLog
type SchemaAppLog schema.AppLog

// Convert to AppLog entity
func (a SchemaAppLog) ToAppLog() *AppLog {
	item := new(AppLog)
	structure.Copy(a, item)
	return item
}

// AppLog entity
type AppLog struct {
	util.Model
	UserId uint64 `gorm:"index"`         // 创建用户 id
	Appid  string `gorm:"string;index"`  // Appid
	Data   string `gorm:"type:text"`     // 详情
	Ip     string `gorm:"size:45;index"` // ip

}

// Convert to AppLog schema
func (a AppLog) ToSchemaAppLog() *schema.AppLog {
	item := new(schema.AppLog)
	structure.Copy(a, item)
	return item
}

// AppLog entity list
type AppLogs []*AppLog

// Convert to AppLog schema list
func (a AppLogs) ToSchemaAppLogs() []*schema.AppLog {
	list := make([]*schema.AppLog, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaAppLog()
	}
	return list
}
