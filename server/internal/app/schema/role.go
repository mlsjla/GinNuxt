package schema

import (
	"time"
)

// Role 角色对象
type Role struct {
	ID        uint64    `json:"id,string"`                             // 唯一标识
	Name      string    `json:"name" binding:"required"`               // 角色名称
	Sequence  int       `json:"sequence"`                              // 排序值
	Memo      string    `json:"memo"`                                  // 备注
	Status    int       `json:"status" binding:"required,max=2,min=1"` // 状态(1:启用 2:禁用)
	Creator   uint64    `json:"creator"`                               // 创建者
	CreatedAt time.Time `json:"created_at"`                            // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                            // 更新时间
}

// RoleQueryParam 查询条件
type RoleQueryParam struct {
	PaginationParam
	IDs        []uint64 `form:"-"`          // 唯一标识列表
	Name       string   `form:"-"`          // 角色名称
	QueryValue string   `form:"queryValue"` // 模糊查询
	Status     int      `form:"status"`     // 状态(1:启用 2:禁用)
}

// RoleQueryOptions 查询可选参数项
type RoleQueryOptions struct {
	OrderFields  []*OrderField // 排序字段
	SelectFields []string      // 查询字段
}

// RoleQueryResult 查询结果
type RoleQueryResult struct {
	Data       Roles
	PageResult *PaginationResult
}

// Roles 角色对象列表
type Roles []*Role

// ToNames 获取角色名称列表
func (a Roles) ToNames() []string {
	names := make([]string, len(a))
	for i, item := range a {
		names[i] = item.Name
	}
	return names
}

// ToMap 转换为键值存储
func (a Roles) ToMap() map[uint64]*Role {
	m := make(map[uint64]*Role)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}
