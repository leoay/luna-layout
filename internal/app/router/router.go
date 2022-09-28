package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/leoay/luna/pkg/auth"
	"luna-layout/internal/app/api"
	"luna-layout/internal/app/middleware"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	Auth     auth.Auther
	GreetAPI *api.GreetAPI
} // end

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	g.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
	))

	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{
		//pub := v1.Group("/pub")
		//{
		//	gLogin := pub.Group("login")
		//	{
		//		//gLogin.POST("exit", a.LoginAPI.Logout)
		//	}
		//}

		gGreet := v1.Group("greet")
		{
			gGreet.GET("", a.GreetAPI.Query)
		}
	} // v1 end
}
