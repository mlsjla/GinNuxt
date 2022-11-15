package dao

import (
	"strings"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/config"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/attachment"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/casbin_rule"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/category"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/menu"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/mobile_code"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/post"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/role"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/role_menu"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/setting"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/thread"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/user"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
) // end

// RepoSet repo injection
var RepoSet = wire.NewSet(
	util.TransSet,
	menu.MenuSet,
	role.RoleSet,
	user.UserRoleSet,
	user.UserSet,
	thread.ThreadSet,
	category.CategorySet,
	post.PostSet,
	attachment.AttachmentSet,
	setting.SettingSet,
	mobile_code.MobileCodeSet,
	casbin_rule.CasbinRuleSet,
	role_menu.RoleMenuSet,
) // end

// Define repo type alias
type (
	TransRepo      = util.Trans
	MenuRepo       = menu.MenuRepo
	RoleRepo       = role.RoleRepo
	UserRoleRepo   = user.UserRoleRepo
	UserRepo       = user.UserRepo
	ThreadRepo     = thread.ThreadRepo
	CategoryRepo   = category.CategoryRepo
	PostRepo       = post.PostRepo
	AttachmentRepo = attachment.AttachmentRepo
	SettingRepo    = setting.SettingRepo
	MobileCodeRepo = mobile_code.MobileCodeRepo
	CasbinRuleRepo = casbin_rule.CasbinRuleRepo
	RoleMenuRepo   = role_menu.RoleMenuRepo
) // end

// Auto migration for given models
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(menu.Menu),
		new(role.Role),
		new(user.UserRole),
		new(user.User),
		new(thread.Thread),
		new(category.Category),
		new(post.Post),
		new(attachment.Attachment),
		new(setting.Setting),
		new(mobile_code.MobileCode),
		new(casbin_rule.CasbinRule),
		new(role_menu.RoleMenu),
	) // end
}
