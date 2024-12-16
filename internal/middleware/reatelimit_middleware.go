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

// PublicAPIRateLimiter applies the public API rate-limiting logic
func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {

		urlPath := c.Request.URL.Path
		// Get the appropriate limiter based on the URL path
		limiterInstance := rl.filterLimiterUrlPath(urlPath)
		if limiterInstance != nil {
			log.Printf("No rate limiter configured for path: %s\n", urlPath)
			c.Next()
			return
		}

		// Use the client's IP as the key
		clientIP := c.ClientIP() // Optionally, customize the key
		log.Printf("Applying Public API rate limiting for path: %s and client IP: %s\n", urlPath, clientIP)

		limitContext, err := limiterInstance.Get(c, clientIP)
		if err != nil {
			fmt.Println("Failed to check rate limiter for PUBLIC API.", err)
			c.Next() // Allow proceeding if there is an error
			return
		}

		if limitContext.Reached {
			log.Printf("Public API rate limit reached for path: %s\n", urlPath)
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				gin.H{"error": "Rate limit exceeded for public API."})
			return
		}

		// Allow the request if rate limit is not exceeded
		c.Next()
	}
}

func (rl *RateLimiter) UserAndPrivateAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {

		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimiterUrlPath(urlPath)
		if rateLimitPath != nil {
			userID := 1001 //context.GetUserIdFromUUID()
			key := fmt.Sprintf("%d-%s", userID, urlPath)
			limitContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limiter for USER API.", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limiter is reached %s\n", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limiter breached USER API."})
				return
			}
		}
	}
}

// filterLimiterUrlPath selects the appropriate rate limiter based on URL path
func (rl *RateLimiter) filterLimiterUrlPath(urlPath string) *limiter.Limiter {

	// Normalize the URL path to ensure consistent comparison
	switch urlPath {
	case "/v1/2024/user/login": // Ensure path starts with "/" for proper comparison
		return rl.publicAPIRateLimiter

	case "/v1/2024/user/info":
		return rl.userPrivateAPIRateLimiter

	default:
		// For any other unhandled routes, use the global rate limiter
		return rl.globalRateLimiter
	}
}
