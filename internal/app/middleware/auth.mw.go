package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/leoay/luna/pkg/auth"
	"github.com/leoay/luna/pkg/errors"
	"github.com/leoay/luna/pkg/logger"
	"luna-layout/internal/app/config"
	"luna-layout/internal/app/contextx"
	"luna-layout/internal/app/ginx"
)

func wrapGreetAuthContext(c *gin.Context, GreetID uint64, GreetName string) {
	ctx := contextx.NewGreetID(c.Request.Context(), GreetID)
	ctx = contextx.NewGreetName(ctx, GreetName)
	ctx = logger.NewGreetIDContext(ctx, GreetID)
	ctx = logger.NewGreetNameContext(ctx, GreetName)
	c.Request = c.Request.WithContext(ctx)
}

// Valid Greet token (jwt)
func GreetAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	if !config.C.JWTAuth.Enable {
		return func(c *gin.Context) {
			wrapGreetAuthContext(c, config.C.Root.GreetID, config.C.Root.GreetName)
			c.Next()
		}
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		tokenGreetID, err := a.ParseGreetID(c.Request.Context(), ginx.GetToken(c))
		if err != nil {
			if err == auth.ErrInvalidToken {
				if config.C.IsDebugMode() {
					wrapGreetAuthContext(c, config.C.Root.GreetID, config.C.Root.GreetName)
					c.Next()
					return
				}
				ginx.ResError(c, errors.ErrInvalidToken)
				return
			}
			ginx.ResError(c, errors.WithStack(err))
			return
		}

		idx := strings.Index(tokenGreetID, "-")
		if idx == -1 {
			ginx.ResError(c, errors.ErrInvalidToken)
			return
		}

		GreetID, _ := strconv.ParseUint(tokenGreetID[:idx], 10, 64)
		wrapGreetAuthContext(c, GreetID, tokenGreetID[idx+1:])
		c.Next()
	}
}
