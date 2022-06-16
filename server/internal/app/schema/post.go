package schema

import "time"

// 文章详情
type Post struct {
	ID            uint64    `json:"id,string"`
	UserId        uint64    `json:"user_id"`                                    // 发布者ID
	ThreadId      uint64    `json:"thread_id,string"`                           // 文章ID
	ReplyPostId   uint64    `json:"reply_post_id"`                              // 回复postID
	ReplyUserId   uint64    `json:"reply_user_id"`                              // 回复userID
	CommentPostId uint64    `json:"comment_post_id"`                            // 评论postID
	CommentUserId uint64    `json:"comment_user_id"`                            // 评论userID
	Content       string    `json:"content"`                                    // 详情
	Ip            string    `json:"ip"`                                         // 排序
	Port          int       `json:"port"`                                       // 端口
	ReplyCount    uint64    `gorm:"index";json:"reply_count" binding:"numeric"` // 回复数
	LikeCount     uint64    `gorm:"index";json:"like_count" binding:"numeric"`  // 赞数
	IsFirst       int       `json:"is_first,string"`                            // 是否主文章
	IsComment     int       `json:"is_comment"`                                 // 是否是回复回帖内容
	IsApproved    int       `json:"is_approved"`                                // 是否合法
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Query parameters for db
type PostQueryParam struct {
	PaginationParam
	ThreadId uint64 `form:"threadId"`
	IsFirst  int    `form:"isFirst"`
}

// Query options for db (order or select fields)
type PostQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type PostQueryResult struct {
	Data       Posts
	PageResult *PaginationResult
}

// 文章详情 Object List
type Posts []*Post

type PostTree struct {
	ID            uint64     `json:"id,string"`
	UserId        uint64     `json:"user_id" binding:"required"`      // 发布者ID
	ThreadId      uint64     `json:"thread_id"`                       // 文章ID
	ReplyPostId   uint64     `json:"reply_post_id"`                   // 回复postID
	ReplyUserId   uint64     `json:"reply_user_id"`                   // 回复userID
	CommentPostId uint64     `json:"comment_post_id"`                 // 评论postID
	CommentUserId uint64     `json:"comment_user_id"`                 // 评论userID
	Content       string     `json:"content"`                         // 详情
	Ip            string     `json:"ip"`                              // 排序
	Port          int        `json:"port"`                            // 端口
	ReplyCount    uint64     `json:"reply_count" binding:"LikeCount"` // 回复数
	LikeCount     uint64     `json:"like_count" binding:""`           // 赞数
	IsFirst       int        `json:"is_first"`                        // 是否主文章
	IsComment     int        `json:"is_comment"`                      // 是否是回复回帖内容
	IsApproved    int        `json:"is_approved"`                     // 是否合法
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Children      *PostTrees `yaml:"children,omitempty" json:"children,omitempty"` // 子级树
}

// PostTrees 分类树列表
type PostTrees []*PostTree

// ToTree 转换为树形结构
func (a PostTrees) ToTree() PostTrees {
	mi := make(map[uint64]*PostTree)
	for _, item := range a {
		mi[item.ID] = item
	}

	var list PostTrees
	for _, item := range a {
		if item.ReplyPostId == 0 {
			list = append(list, item)
			continue
		}
		if pitem, ok := mi[item.ReplyPostId]; ok {
			if pitem.Children == nil {
				children := PostTrees{item}
				pitem.Children = &children
				continue
			}
			*pitem.Children = append(*pitem.Children, item)
		}
	}
	return list
}

// ToTree 转换为分类树
// ToTree 转换为菜单树
func (a Posts) ToTree() PostTrees {
	list := make(PostTrees, len(a))
	for i, item := range a {
		list[i] = &PostTree{
			ID:         item.ID,
			UserId:     item.UserId,
			ReplyCount: item.ReplyCount,
			LikeCount:  item.LikeCount,
			IsFirst:    item.IsFirst,
		}
	}
	return list.ToTree()
}
