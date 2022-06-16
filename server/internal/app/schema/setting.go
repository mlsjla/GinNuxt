package schema

import "time"

// 系统设置
type Setting struct {
	ID        uint64    `json:"id,string"`
	Key       string    `json:"key";uniqueIndex;binding:"required"` // 设置项key
	Value     string    `json:"value";binding:"required"`           // 设置项value
	Tag       string    `json:"tag";default:"default"`              // 设置项tag
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Query parameters for db
type SettingQueryParam struct {
	PaginationParam
	Tag        string `form:"tag"`
	Key        string `form:"key"`
	QueryValue string `form:"queryValue"` // 模糊查询
	Status     int    `form:"status"`     // 用户状态(1:启用 2:停用)
}

// Query options for db (order or select fields)
type SettingQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type SettingQueryResult struct {
	Data       Settings
	PageResult *PaginationResult
}

// 系统设置 Object List
type Settings []*Setting
