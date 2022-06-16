package setting

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get Setting db model
func GetSettingDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Setting))
}

// Setting
type SchemaSetting schema.Setting

// Convert to Setting entity
func (a SchemaSetting) ToSetting() *Setting {
	item := new(Setting)
	structure.Copy(a, item)
	return item
}

// Setting entity
type Setting struct {
	util.Model
	Key   string `gorm:"size:255;uniqueIndex"`                 // 设置项key
	Value string `gorm:"type:text"`                            // 设置项value
	Tag   string `gorm:"size:255;default:'default';not null;"` // 设置项tag

}

// Convert to Setting schema
func (a Setting) ToSchemaSetting() *schema.Setting {
	item := new(schema.Setting)
	structure.Copy(a, item)
	return item
}

// Setting entity list
type Settings []*Setting

// Convert to Setting schema list
func (a Settings) ToSchemaSettings() []*schema.Setting {
	list := make([]*schema.Setting, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaSetting()
	}
	return list
}
