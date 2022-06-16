package schema

import "time"

// 菜单权限
type RoleMenu struct {
	ID        uint64    `json:"id,string"`
	RoleId    uint64    `json:"role_id,string" binding:"required"` // 角色ID
	MenuId    uint64    `json:"menu_id,string"`                    // 菜单ID
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Query parameters for db
type RoleMenuQueryParam struct {
	PaginationParam
	RoleId string `form:"role_id,string"` //
}

// Query options for db (order or select fields)
type RoleMenuQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type RoleMenuQueryResult struct {
	Data       RoleMenus
	PageResult *PaginationResult
}

// 菜单权限 Object List
type RoleMenus []*RoleMenu
