// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/mlsjla/gin-nuxt/server/internal/app/api"
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
	"github.com/mlsjla/gin-nuxt/server/internal/app/module/adapter"
	"github.com/mlsjla/gin-nuxt/server/internal/app/router"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

import (
	_ "github.com/mlsjla/gin-nuxt/server/internal/app/swagger"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	roleRepo := &role.RoleRepo{
		DB: db,
	}
	userRepo := &user.UserRepo{
		DB: db,
	}
	userRoleRepo := &user.UserRoleRepo{
		DB: db,
	}
	casbinRuleRepo := &casbin_rule.CasbinRuleRepo{
		DB: db,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleRepo:       roleRepo,
		UserRepo:       userRepo,
		UserRoleRepo:   userRoleRepo,
		CasbinRuleRepo: casbinRuleRepo,
	}
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	trans := &util.Trans{
		DB: db,
	}
	userSrv := &service.UserSrv{
		Enforcer:     syncedEnforcer,
		TransRepo:    trans,
		UserRepo:     userRepo,
		UserRoleRepo: userRoleRepo,
		RoleRepo:     roleRepo,
	}
	menuRepo := &menu.MenuRepo{
		DB: db,
	}
	loginSrv := &service.LoginSrv{
		Auth:         auther,
		UserRepo:     userRepo,
		UserRoleRepo: userRoleRepo,
		RoleRepo:     roleRepo,
		MenuRepo:     menuRepo,
	}
	menuSrv := &service.MenuSrv{
		TransRepo: trans,
		MenuRepo:  menuRepo,
	}
	roleMenuRepo := &role_menu.RoleMenuRepo{
		DB: db,
	}
	roleMenuSrv := &service.RoleMenuSrv{
		TransRepo:    trans,
		RoleMenuRepo: roleMenuRepo,
	}
	loginAPI := &api.LoginAPI{
		UserSrv:     userSrv,
		LoginSrv:    loginSrv,
		MenuSrv:     menuSrv,
		RoleMenuSrv: roleMenuSrv,
	}
	menuAPI := &api.MenuAPI{
		MenuSrv: menuSrv,
	}
	roleSrv := &service.RoleSrv{
		Enforcer:  syncedEnforcer,
		TransRepo: trans,
		RoleRepo:  roleRepo,
		UserRepo:  userRepo,
	}
	roleAPI := &api.RoleAPI{
		RoleSrv: roleSrv,
	}
	userAPI := &api.UserAPI{
		UserSrv: userSrv,
	}
	threadRepo := &thread.ThreadRepo{
		DB: db,
	}
	threadSrv := &service.ThreadSrv{
		TransRepo:  trans,
		ThreadRepo: threadRepo,
	}
	postRepo := &post.PostRepo{
		DB: db,
	}
	postSrv := &service.PostSrv{
		TransRepo: trans,
		PostRepo:  postRepo,
	}
	threadAPI := &api.ThreadAPI{
		ThreadSrv: threadSrv,
		PostSrv:   postSrv,
	}
	categoryRepo := &category.CategoryRepo{
		DB: db,
	}
	categorySrv := &service.CategorySrv{
		TransRepo:    trans,
		CategoryRepo: categoryRepo,
	}
	categoryAPI := &api.CategoryAPI{
		CategorySrv: categorySrv,
	}
	postAPI := &api.PostAPI{
		PostSrv: postSrv,
	}
	attachmentRepo := &attachment.AttachmentRepo{
		DB: db,
	}
	attachmentSrv := &service.AttachmentSrv{
		TransRepo:      trans,
		AttachmentRepo: attachmentRepo,
	}
	attachmentAPI := &api.AttachmentAPI{
		AttachmentSrv: attachmentSrv,
	}
	uploadAPI := &api.UploadAPI{
		AttachmentSrv: attachmentSrv,
	}
	settingRepo := &setting.SettingRepo{
		DB: db,
	}
	settingSrv := &service.SettingSrv{
		TransRepo:   trans,
		SettingRepo: settingRepo,
	}
	settingAPI := &api.SettingAPI{
		SettingSrv: settingSrv,
	}
	mobileCodeRepo := &mobile_code.MobileCodeRepo{
		DB: db,
	}
	mobileCodeSrv := &service.MobileCodeSrv{
		TransRepo:      trans,
		MobileCodeRepo: mobileCodeRepo,
	}
	mobileCodeAPI := &api.MobileCodeAPI{
		MobileCodeSrv: mobileCodeSrv,
	}
	casbinRuleSrv := &service.CasbinRuleSrv{
		TransRepo:      trans,
		CasbinRuleRepo: casbinRuleRepo,
	}
	casbinRuleAPI := &api.CasbinRuleAPI{
		CasbinRuleSrv: casbinRuleSrv,
	}
	roleMenuAPI := &api.RoleMenuAPI{
		RoleMenuSrv: roleMenuSrv,
	}
	routerRouter := &router.Router{
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		LoginAPI:       loginAPI,
		MenuAPI:        menuAPI,
		RoleAPI:        roleAPI,
		UserAPI:        userAPI,
		ThreadAPI:      threadAPI,
		CategoryAPI:    categoryAPI,
		PostAPI:        postAPI,
		AttachmentAPI:  attachmentAPI,
		UploadAPI:      uploadAPI,
		SettingAPI:     settingAPI,
		MobileCodeAPI:  mobileCodeAPI,
		CasbinRuleAPI:  casbinRuleAPI,
		UserSrv:        userSrv,
		RoleMenuAPI:    roleMenuAPI,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuSrv:        menuSrv,
		SettingSrv:     settingSrv,
		UserSrv:        userSrv,
	}
	return injector, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
