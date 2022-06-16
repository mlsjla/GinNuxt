package attachment

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get Attachment db model
func GetAttachmentDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Attachment))
}

// Attachment
type SchemaAttachment schema.Attachment

// Convert to Attachment entity
func (a SchemaAttachment) ToAttachment() *Attachment {
	item := new(Attachment)
	structure.Copy(a, item)
	return item
}

// Attachment entity
type Attachment struct {
	util.Model
	Uuid       string `gorm:"string"`        // 发布者ID
	UserId     uint64 `gorm:"id,string"`     // 发布者ID
	TypeId     uint64 `gorm:"id,string"`     // 类型数据ID
	Order      int    `gorm:""`              // 附件排序
	Type       int    `gorm:""`              // 附件类型
	IsRemote   int    `gorm:""`              // 是否远程附件
	Attachment string `gorm:"size:255"`      // 文件系统生成的名称
	FilePath   string `gorm:"size:255"`      // 文件路径
	FileName   string `gorm:"size:255"`      // 文件原名称
	FileSize   uint64 `gorm:""`              // 文件大小
	FileWidth  uint64 `gorm:""`              // 宽度
	FileHeight uint64 `gorm:""`              // 高度
	FileType   string `gorm:"size:255"`      // 文件类型
	Ip         string `gorm:"size:45;index"` // ip

}

// Convert to Attachment schema
func (a Attachment) ToSchemaAttachment() *schema.Attachment {
	item := new(schema.Attachment)
	structure.Copy(a, item)
	return item
}

// Attachment entity list
type Attachments []*Attachment

// Convert to Attachment schema list
func (a Attachments) ToSchemaAttachments() []*schema.Attachment {
	list := make([]*schema.Attachment, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaAttachment()
	}
	return list
}
