package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/global"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
	"net/http"
	"time"
)

type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter
	publicAPIRateLimiter      *limiter.Limiter
	userPrivateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	rateLimit := &RateLimiter{
		globalRateLimiter:         rateLimiter("100-S"),
		publicAPIRateLimiter:      rateLimiter("80-S"),
		userPrivateAPIRateLimiter: rateLimiter("60-S"),
	}
	return rateLimit
}

func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(global.Rdb, limiter.StoreOptions{
		Prefix:          "rate-limiter",
		MaxRetry:        3,
		CleanUpInterval: time.Hour,
	})
	if err != nil {
		return nil
	}
	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic(err)
	}
	instance := limiter.New(store, rate)
	return instance
}

func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global-rate-limiter"
		log.Println(key)
		limitContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			fmt.Println("Failed to check rate limiter GLOBAL rate limiter.", err)
			c.Next()
			return
		}
		if limitContext.Reached {
			fmt.Printf("Rate limiter is reached %s\n", key)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limiter breached GLOBAL rate limiter."})
			return
		}
		c.Next()
	}
}

func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimiterUrlPath(urlPath)
		if limitContext.Reached {
			fmt.Println("Rate limiter is reached")

		}
	}
}

func (rl *RateLimiter) UserPrivateAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (rl *RateLimiter) filterLimiterUrlPath(urlPath string) *limiter.Limiter {
	if urlPath == "v1/2024/user/login" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "v1/2024/user/info" {
		return rl.userPrivateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}
