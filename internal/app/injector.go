package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/leoay/luna/pkg/auth"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine *gin.Engine
	Auth   auth.Auther
}
