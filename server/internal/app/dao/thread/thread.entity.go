package thread

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
	"github.com/shopspring/decimal"
)

// Get Thread db model
func GetThreadDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Thread))
}

// Thread
type SchemaThread schema.Thread

// Convert to Thread entity
func (a SchemaThread) ToThread() *Thread {
	item := new(Thread)
	structure.Copy(a, item)
	return item
}

// Thread entity
type Thread struct {
	util.Model
	UserId           uint64          `gorm:"index"`            // 创建用户 id
	lastPostedUserId uint64          `gorm:""`                 // 最后回复用户
	CategoryId       int             `gorm:""`                 // 分类 id
	Type             int             `gorm:""`                 // 类型
	Title            string          `gorm:"size:255;index;"`  // 类型
	Cover            string          `gorm:"size:255;index;"`  // 类型
	Summary          string          `gorm:"size:1000;index;"` // 类型
	Username         string          `gorm:"size:255;index;"`  // 类型
	Price            decimal.Decimal `gorm:""`                 // 价格
	AttachmentPrice  decimal.Decimal `gorm:""`                 // 附件价格
	FreeWords        decimal.Decimal `gorm:""`                 // 免费字数百分比
	PostCount        int             `gorm:""`                 // 回复数
	ViewCount        int             `gorm:""`                 // 查看数
	RewardedCount    int             `gorm:""`                 // 打赏数
	PaidCount        int             `gorm:""`                 // 付费数
	ShareCount       int             `gorm:""`                 // 分享数
	Longitude        decimal.Decimal `gorm:""`                 // 经度
	Latitude         decimal.Decimal `gorm:""`                 // 纬度
	Address          string          `gorm:"size:128;index;"`  // 地址
	Location         string          `gorm:"size:128;index;"`  // 位置
	PostedAt         time.Time       `gorm:""`                 // 最新评论时间
	IssueAt          time.Time       `gorm:""`                 // 审核变更的时间记录
	IsApproved       int             `gorm:""`                 // 是否合法
	IsSticky         int             `gorm:""`                 // 是否置顶
	IsEssence        int             `gorm:""`                 // 是否加精
	IsSite           int             `gorm:""`                 // 是否推荐到首页
	IsAnonymous      int             `gorm:""`                 // 是否匿名
	IsDisplay        int             `gorm:""`                 // 是否显示
	IsRedPacket      int             `gorm:""`                 // 是否添加红包
	IsDraft          int             `gorm:""`                 // 是否为草稿
	Source           int             `gorm:""`                 // 来源

}

// Convert to Thread schema
func (a Thread) ToSchemaThread() *schema.Thread {
	item := new(schema.Thread)
	structure.Copy(a, item)
	return item
}

// Thread entity list
type Threads []*Thread

// Convert to Thread schema list
func (a Threads) ToSchemaThreads() []*schema.Thread {
	list := make([]*schema.Thread, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaThread()
	}
	return list
}
