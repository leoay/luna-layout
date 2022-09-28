//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"luna-layout/internal/app/api"
	"luna-layout/internal/app/dao"
	"luna-layout/internal/app/router"
	"luna-layout/internal/app/service"
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
