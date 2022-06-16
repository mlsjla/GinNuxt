package category

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get Category db model
func GetCategoryDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Category))
}

// Category
type SchemaCategory schema.Category

// Convert to Category entity
func (a SchemaCategory) ToCategory() *Category {
	item := new(Category)
	structure.Copy(a, item)
	return item
}

// Category entity
type Category struct {
	util.Model
	Name        string `gorm:"size:255;index;"` // 分类名称
	Description string `gorm:"type:text;"`      // 分类描述
	Icon        string `gorm:"size:255;"`       // 图标
	Sort        int    `gorm:"index"`           // 排序
	Property    int    `gorm:""`                // 属性
	ThreadCount uint64 `gorm:""`                // 主题数
	Moderators  string `gorm:"size:255;"`       // 排序
	Ip          string `gorm:"size:45;"`        // 排序
	Parentid    uint64 `gorm:""`                // 父ID

}

// Convert to Category schema
func (a Category) ToSchemaCategory() *schema.Category {
	item := new(schema.Category)
	structure.Copy(a, item)
	return item
}

// Category entity list
type Categories []*Category

// Convert to Category schema list
func (a Categories) ToSchemaCategories() []*schema.Category {
	list := make([]*schema.Category, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaCategory()
	}
	return list
}
