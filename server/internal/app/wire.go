//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/api"
	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/module/adapter"
	"github.com/mlsjla/gin-nuxt/server/internal/app/router"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
