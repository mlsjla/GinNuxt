package schema

import "time"

// APPLog
type AppLog struct {
	ID        uint64    `json:"id,string"`
	UserId    uint64    `json:"user_id" binding:"required"` // 创建用户 id
	Appid     string    `json:"appid" binding:"required"`   // Appid
	Data      string    `json:"data"`                       // 详情
	Ip        string    `json:"ip"`                         // ip
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AppPack struct {
	Appid      string `json:"appid"` // Appid
	Type       int    `json:"type"`  // 类型
	Privatekey string `json:"privatekey"`
	Api        string `json:"api"`
	Remark     string `json:"remark"`
}

// Query parameters for db
type AppLogQueryParam struct {
	PaginationParam
}

// Query options for db (order or select fields)
type AppLogQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type AppLogQueryResult struct {
	Data       AppLogs
	PageResult *PaginationResult
}

// APPLog Object List
type AppLogs []*AppLog
