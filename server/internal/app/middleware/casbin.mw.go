package middleware

import (
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/mlsjla/gin-nuxt/server/internal/app/config"
	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Valid use interface permission
func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, UserSrv *service.UserSrv, allowGuest []SkipperFunc, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := config.C.Casbin
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		if SkipHandler(c, allowGuest...) {
			c.Next()
			return
		}

		p := c.Request.URL.Path
		m := c.Request.Method
		userID := contextx.FromUserID(c.Request.Context())
		ctx := c.Request.Context()
		user, err := UserSrv.Get(ctx, userID)
		if err != nil && userID != 1 {
			ginx.ResError(c, err)
			return
		}
		// 校验权限
		if user != nil {
			userRoles := user.UserRoles
			fmt.Println("user", user)
			for _, role := range userRoles {
				if b, err := enforcer.Enforce(strconv.FormatUint(role.RoleID, 10), p, m); err != nil {
					ginx.ResError(c, errors.WithStack(err))
					return
				} else if b {
					c.Next()
					return
				}
			}
		}

		if b, err := enforcer.Enforce(strconv.FormatUint(userID, 10), p, m); err != nil {
			ginx.ResError(c, errors.WithStack(err))
			return
		} else if !b {
			ginx.ResError(c, errors.ErrNoPerm)
			return
		}
		c.Next()
	}
}
