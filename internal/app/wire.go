//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"server/internal/app/api"
	"server/internal/app/dao"
	"server/internal/app/router"
	"server/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		InitAuth,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
