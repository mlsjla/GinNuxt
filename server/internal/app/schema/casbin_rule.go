package schema

import "time"

// 权限设置
type CasbinRule struct {
	ID        uint64    `json:"id,string"`
	PType     string    `json:"p_type" binding:"required"` // 规则类型
	V0        string    `json:"v0" binding:"required"`     // 角色ID
	V1        string    `json:"v1" binding:"required"`     // api路径
	V2        string    `json:"v2" binding:"required"`     // api访问方法
	V3        string    `json:"v3" binding:""`
	V4        string    `json:"v4" binding:""`
	V5        string    `json:"v5" binding:""`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Query parameters for db
type CasbinRuleQueryParam struct {
	PaginationParam
	V0 string `form:"v0,string"` //
}

// Query options for db (order or select fields)
type CasbinRuleQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type CasbinRuleQueryResult struct {
	Data       CasbinRules
	PageResult *PaginationResult
}

// 权限设置 Object List
type CasbinRules []*CasbinRule
