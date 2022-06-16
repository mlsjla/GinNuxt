package mobile_code

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get MobileCode db model
func GetMobileCodeDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(MobileCode))
}

// MobileCode
type SchemaMobileCode schema.MobileCode

// Convert to MobileCode entity
func (a SchemaMobileCode) ToMobileCode() *MobileCode {
	item := new(MobileCode)
	structure.Copy(a, item)
	return item
}

// MobileCode entity
type MobileCode struct {
	util.Model
	Mobile    string    `gorm:"size:20"`       // 手机号
	Code      string    `gorm:"size:20"`       // 验证码
	Type      string    `gorm:"size:20"`       // 验证类型
	State     uint      `gorm:"size:20"`       // 验证状态
	Ip        string    `gorm:"size:45;index"` // 排序
	ExpiredAt time.Time `gorm:""`              // 过期时间

}

// Convert to MobileCode schema
func (a MobileCode) ToSchemaMobileCode() *schema.MobileCode {
	item := new(schema.MobileCode)
	structure.Copy(a, item)
	return item
}

// MobileCode entity list
type MobileCodes []*MobileCode

// Convert to MobileCode schema list
func (a MobileCodes) ToSchemaMobileCodes() []*schema.MobileCode {
	list := make([]*schema.MobileCode, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaMobileCode()
	}
	return list
}
