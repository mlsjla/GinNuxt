package schema

import "time"

// 短信发送
type MobileCode struct {
	ID        uint64    `json:"id,string"`
	Mobile    string    `json:"mobile" binding:"required"` // 手机号
	Code      string    `json:"code" binding:"required"`   // 验证码
	Type      string    `json:"type" binding:"required"`   // 验证类型
	State     uint      `json:"state" binding:"required"`  // 验证状态
	Ip        string    `json:"ip"`                        // 排序
	ExpiredAt time.Time `json:"expired_at"`                // 过期时间
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Query parameters for db
type MobileCodeQueryParam struct {
	PaginationParam
}

// Query options for db (order or select fields)
type MobileCodeQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type MobileCodeQueryResult struct {
	Data       MobileCodes
	PageResult *PaginationResult
}

// 短信发送 Object List
type MobileCodes []*MobileCode
