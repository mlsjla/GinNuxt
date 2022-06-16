package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mlsjla/gin-nuxt/server/internal/app/config"
	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/pkg/auth"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/logger"
)

func wrapUserAuthContext(c *gin.Context, userID uint64, userName string) {
	ctx := contextx.NewUserID(c.Request.Context(), userID)
	ctx = contextx.NewUserName(ctx, userName)
	ctx = logger.NewUserIDContext(ctx, userID)
	ctx = logger.NewUserNameContext(ctx, userName)
	c.Request = c.Request.WithContext(ctx)
}

// Valid user token (jwt)
func UserAuthMiddleware(a auth.Auther, allowGuest []SkipperFunc, skippers ...SkipperFunc) gin.HandlerFunc {
	if !config.C.JWTAuth.Enable {
		return func(c *gin.Context) {
			wrapUserAuthContext(c, config.C.Root.UserID, config.C.Root.Username)
			c.Next()
		}
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		tokenUserID, err := a.ParseUserID(c.Request.Context(), ginx.GetToken(c))
		if err != nil {
			if err == auth.ErrInvalidToken {
				if false && config.C.IsDebugMode() {
					wrapUserAuthContext(c, config.C.Root.UserID, config.C.Root.Username)
					c.Next()
					return
				} else if SkipHandler(c, allowGuest...) {
					wrapUserAuthContext(c, 0, "游客")
					c.Next()
					return
				}
				ginx.ResError(c, errors.ErrInvalidToken)
				return
			}
			ginx.ResError(c, errors.WithStack(err))
			return
		}

		idx := strings.Index(tokenUserID, "-")
		if idx == -1 {
			ginx.ResError(c, errors.ErrInvalidToken)
			return
		}

		userID, _ := strconv.ParseUint(tokenUserID[:idx], 10, 64)
		wrapUserAuthContext(c, userID, tokenUserID[idx+1:])
		c.Next()
	}
}
