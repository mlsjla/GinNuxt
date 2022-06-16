package schema

import (
	"time"

	"github.com/shopspring/decimal"
)

// 文章
type Thread struct {
	ID               uint64          `json:"id,string"`
	UserId           uint64          `json:"user_id"`                               // 创建用户 id
	lastPostedUserId uint64          `json:"last_posted_user_id"`                   // 最后回复用户
	CategoryId       int             `json:"category_id,string" binding:"required"` // 分类 id
	Type             int             `json:"type" binding:"numeric"`                // 类型
	Title            string          `json:"title" binding:"required"`
	Cover            string          `json:"cover"`
	Summary          string          `json:"summary"`
	Content          string          `json:"content"`
	Username         string          `json:"username"`                                     // 类型
	Price            decimal.Decimal `json:"price" binding:"numeric"`                      // 价格
	AttachmentPrice  decimal.Decimal `json:"attachment_price" binding:"numeric","numeric"` // 附件价格
	FreeWords        decimal.Decimal `json:"free_words" binding:"numeric";numeric"`        // 免费字数百分比
	PostCount        int             `json:"post_count"`                                   // 回复数
	ViewCount        int             `json:"view_count"`                                   // 查看数
	RewardedCount    int             `json:"rewarded_count"`                               // 打赏数
	PaidCount        int             `json:"paid_count"`                                   // 付费数
	ShareCount       int             `json:"share_count"`                                  // 分享数
	Longitude        decimal.Decimal `json:"longitude" binding:"numeric";"numeric"`        // 经度
	Latitude         decimal.Decimal `json:"latitude" binding:"numeric";"numeric"`         // 纬度
	Address          string          `json:"address"`                                      // 地址
	Location         string          `json:"location"`                                     // 位置
	PostedAt         time.Time       `json:"posted_at"`                                    // 最新评论时间
	IssueAt          time.Time       `json:"issue_at"`                                     // 审核变更的时间记录
	IsApproved       int             `json:"is_approved"`                                  // 是否合法
	IsSticky         int             `json:"is_sticky"`                                    // 是否置顶
	IsEssence        int             `json:"is_essence"`                                   // 是否加精
	IsSite           int             `json:"is_site"`                                      // 是否推荐到首页
	IsAnonymous      int             `json:"is_anonymous"`                                 // 是否匿名
	IsDisplay        int             `json:"is_display"`                                   // 是否显示
	IsRedPacket      int             `json:"is_red_packet"`                                // 是否添加红包
	IsDraft          int             `json:"is_draft"`                                     // 是否为草稿
	Source           int             `json:"source"`                                       // 来源
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

// Query parameters for db
type ThreadQueryParam struct {
	CategoryId string `form:"category_id"` // 分类 id
	PaginationParam
}

// Query options for db (order or select fields)
type ThreadQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type ThreadQueryResult struct {
	Data       Threads
	PageResult *PaginationResult
}

// 文章 Object List
type Threads []*Thread
