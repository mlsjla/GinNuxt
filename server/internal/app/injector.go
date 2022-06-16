package app

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
	"github.com/mlsjla/gin-nuxt/server/pkg/auth"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine         *gin.Engine
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	MenuSrv        *service.MenuSrv
	SettingSrv     *service.SettingSrv
	UserSrv        *service.UserSrv
}
