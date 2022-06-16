package post

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

// Get Post db model
func GetPostDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Post))
}

// Post
type SchemaPost schema.Post

// Convert to Post entity
func (a SchemaPost) ToPost() *Post {
	item := new(Post)
	structure.Copy(a, item)
	return item
}

// Post entity
type Post struct {
	util.Model
	UserId        uint64 `gorm:"id,string"`        // 发布者ID
	ThreadId      uint64 `gorm:"id,string"`        // 文章ID
	ReplyPostId   uint64 `gorm:"id,string;index"`  // 回复postID
	ReplyUserId   uint64 `gorm:"id,string;index"`  // 回复userID
	CommentPostId uint64 `gorm:"id,string;index"`  // 评论postID
	CommentUserId uint64 `gorm:"id,string;index"`  // 评论userID
	Content       string `gorm:"type:text"`        // 详情
	Ip            string `gorm:"size:45;index"`    // 排序
	Port          int    `gorm:""`                 // 端口
	ReplyCount    uint64 `gorm:""`                 // 回复数
	LikeCount     uint64 `gorm:""`                 // 赞数
	IsFirst       int    `gorm:"type:tinyint(1);"` // 是否主文章
	IsComment     int    `gorm:"type:tinyint(1);"` // 是否是回复回帖内容
	IsApproved    int    `gorm:"type:tinyint(1);"` // 是否合法

}

// Convert to Post schema
func (a Post) ToSchemaPost() *schema.Post {
	item := new(schema.Post)
	structure.Copy(a, item)
	return item
}

// Post entity list
type Posts []*Post

// Convert to Post schema list
func (a Posts) ToSchemaPosts() []*schema.Post {
	list := make([]*schema.Post, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaPost()
	}
	return list
}
