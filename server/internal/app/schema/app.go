package schema

import "time"

// APP
type App struct {
	ID        uint64    `json:"id,string"`
	UserId    uint64    `json:"user_id" binding:"required"` // 创建用户 id
	Type      int       `json:"type"`                       // 类型
	Appid     string    `json:"appid" binding:"required"`   // Appid
	Data      string    `json:"data"`                       // 详情
	Ip        string    `json:"ip"`                         // ip
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Query parameters for db
type AppQueryParam struct {
	PaginationParam
}

// Query options for db (order or select fields)
type AppQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type AppQueryResult struct {
	Data       Apps
	PageResult *PaginationResult
}

// APP Object List
type Apps []*App
