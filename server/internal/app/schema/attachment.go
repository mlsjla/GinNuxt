package schema

import "time"

// 附件
type Attachment struct {
	ID         uint64    `json:"id,string"`
	Uuid       string    `json:"uuid" binding:"required"`    // 发布者ID
	UserId     uint64    `json:"user_id" binding:"required"` // 发布者ID
	TypeId     uint64    `json:"type_id" binding:"required"` // 类型数据ID
	Order      int       `json:"order" binding:"required"`   // 附件排序
	Type       int       `json:"type" binding:"required"`    // 附件类型
	IsRemote   int       `json:"is_remote"`                  // 是否远程附件
	Attachment string    `json:"attachment"`                 // 文件系统生成的名称
	FilePath   string    `json:"file_path"`                  // 文件路径
	FileName   string    `json:"file_name"`                  // 文件原名称
	FileSize   uint64    `json:"file_size"`                  // 文件大小
	FileWidth  uint64    `json:"file_width"`                 // 宽度
	FileHeight uint64    `json:"file_height"`                // 高度
	FileType   string    `json:"file_type"`                  // 文件类型
	Ip         string    `json:"ip"`                         // ip
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Query parameters for db
type AttachmentQueryParam struct {
	PaginationParam
}

// Query options for db (order or select fields)
type AttachmentQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type AttachmentQueryResult struct {
	Data       Attachments
	PageResult *PaginationResult
}

// 附件 Object List
type Attachments []*Attachment
