package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"github.com/mlsjla/gin-nuxt/server/internal/app/config"
	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"golang.org/x/time/rate"
)

// Request rate limter (per minute)
func RateLimiterMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := config.C.RateLimiter
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	rc := config.C.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       cfg.RedisDB,
	})

	limiter := redis_rate.NewLimiter(ring)
	limiter.Fallback = rate.NewLimiter(rate.Inf, 0)

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID := contextx.FromUserID(c.Request.Context())
		if userID != 0 {
			limit := cfg.Count
			rate, delay, allowed := limiter.AllowMinute(fmt.Sprintf("%d", userID), limit)
			if !allowed {
				h := c.Writer.Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(limit-rate, 10))
				delaySec := int64(delay / time.Second)
				h.Set("X-RateLimit-Delay", strconv.FormatInt(delaySec, 10))
				ginx.ResError(c, errors.ErrTooManyRequests)
				return
			}
		}

		c.Next()
	}
}
