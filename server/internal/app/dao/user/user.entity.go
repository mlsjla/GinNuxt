package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/structure"
)

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(User))
}

type SchemaUser schema.User

func (a SchemaUser) ToUser() *User {
	item := new(User)
	structure.Copy(a, item)
	return item
}

type User struct {
	util.Model
	Username       string  `gorm:"size:64;uniqueIndex;default:'';not null;"` // 用户名
	Realname       string  `gorm:"size:64;index;default:'';"`                // 真实姓名
	Password       string  `gorm:"size:64;default:'';"`                      // 密码
	Introduce      string  `gorm:"size:255;default:'';"`                     // 个人介绍
	Email          *string `gorm:"size:255;"`                                // 邮箱
	Mobile         *string `gorm:"size:20;"`                                 // 手机号
	Status         int     `gorm:"index;default:0;"`                         // 状态(1:启用 2:停用)
	Creator        uint64  `gorm:""`                                         // 创建者
	Nickname       string  `gorm:"size:64;index;default:'';"`                // 昵称
	PayPassword    string  `gorm:"size:64;default:'';"`                      // 支付密码
	LastLoginIp    string  `gorm:"size:45;index"`                            // ip
	RegisterIp     string  `gorm:"size:45;index"`                            // ip
	RegisterReason string  `gorm:"size:50"`                                  //
	RejectReason   string  `gorm:"size:100"`                                 //
	UsernameBout   int     `gorm:"default:0;"`                               //
	ThreadCount    int     `gorm:"default:0;"`
	FollowCount    int     `gorm:"default:0;"`
	FansCount      int     `gorm:"default:0;"`
	LickdCount     int     `gorm:"default:0;"`
	QuestionCount  int     `gorm:"default:0;"`
	Avatar         string  `gorm:"size:255;"` // 邮箱
}

func (a User) ToSchemaUser() *schema.User {
	item := new(schema.User)
	structure.Copy(a, item)
	return item
}

type Users []*User

func (a Users) ToSchemaUsers() []*schema.User {
	list := make([]*schema.User, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUser()
	}
	return list
}
