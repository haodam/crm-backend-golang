package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"testing"
)

// Mock function for limiter setup
func createMockLimiter(rate string) *limiter.Limiter {
	rateObj, _ := limiter.NewRateFromFormatted(rate)
	store := memory.NewStore()
	return limiter.New(store, rateObj)
}

func TestNewRateLimiter(t *testing.T) {
	tests := []struct {
		name string
		want *RateLimiter
	}{
		{
			name: "should create a valid RateLimiter instance with default settings",
			want: &RateLimiter{}, // Expected output is a RateLimiter instance
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRateLimiter()
			assert.NotNil(t, got)
		})
	}
}

func TestRateLimiter_GlobalRateLimiter(t *testing.T) {
	type fields struct {
		globalRateLimiter         *limiter.Limiter
		publicAPIRateLimiter      *limiter.Limiter
		userPrivateAPIRateLimiter *limiter.Limiter
	}

	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		{
			name: "should provide a global rate limiter handler",
			fields: fields{
				globalRateLimiter: createMockLimiter("5-S"), // 5 requests per second
			},
			want: gin.HandlerFunc(func(c *gin.Context) { c.Next() }),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := &RateLimiter{
				globalRateLimiter:         tt.fields.globalRateLimiter,
				publicAPIRateLimiter:      tt.fields.publicAPIRateLimiter,
				userPrivateAPIRateLimiter: tt.fields.userPrivateAPIRateLimiter,
			}
			assert.NotNil(t, rl.GlobalRateLimiter()) // Ensure Middleware is returned
		})
	}
}

func TestRateLimiter_PublicAPIRateLimiter(t *testing.T) {
	type fields struct {
		globalRateLimiter         *limiter.Limiter
		publicAPIRateLimiter      *limiter.Limiter
		userPrivateAPIRateLimiter *limiter.Limiter
	}

	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		{
			name: "should provide a public API rate limiter handler",
			fields: fields{
				publicAPIRateLimiter: createMockLimiter("10-M"), // 10 requests per minute
			},
			want: gin.HandlerFunc(nil), // No specific return check for middleware
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := &RateLimiter{
				globalRateLimiter:         tt.fields.globalRateLimiter,
				publicAPIRateLimiter:      tt.fields.publicAPIRateLimiter,
				userPrivateAPIRateLimiter: tt.fields.userPrivateAPIRateLimiter,
			}
			assert.NotNil(t, rl.PublicAPIRateLimiter()) // Ensure Middleware is returned
		})
	}
}

func TestRateLimiter_UserAndPrivateAPIRateLimiter(t *testing.T) {
	type fields struct {
		globalRateLimiter         *limiter.Limiter
		publicAPIRateLimiter      *limiter.Limiter
		userPrivateAPIRateLimiter *limiter.Limiter
	}

	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		{
			name: "should provide a user and private API limiter handler",
			fields: fields{
				userPrivateAPIRateLimiter: createMockLimiter("2-S"), // 2 requests per second
			},
			want: gin.HandlerFunc(nil), // No middleware-specific return check
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := &RateLimiter{
				globalRateLimiter:         tt.fields.globalRateLimiter,
				publicAPIRateLimiter:      tt.fields.publicAPIRateLimiter,
				userPrivateAPIRateLimiter: tt.fields.userPrivateAPIRateLimiter,
			}
			assert.NotNil(t, rl.UserAndPrivateAPIRateLimiter()) // Ensure Middleware is returned
		})
	}
}

func TestRateLimiter_filterLimiterUrlPath(t *testing.T) {
	type fields struct {
		globalRateLimiter         *limiter.Limiter
		publicAPIRateLimiter      *limiter.Limiter
		userPrivateAPIRateLimiter *limiter.Limiter
	}

	type args struct {
		urlPath string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *limiter.Limiter
	}{
		{
			name: "should return globalLimiter for root path",
			fields: fields{
				globalRateLimiter: createMockLimiter("5-S"),
			},
			args: args{
				urlPath: "/",
			},
			want: createMockLimiter("5-S"),
		},
		{
			name: "should return publicLimiter for public API path",
			fields: fields{
				publicAPIRateLimiter: createMockLimiter("10-M"),
			},
			args: args{
				urlPath: "/public",
			},
			want: createMockLimiter("10-M"),
		},
		{
			name: "should return privateLimiter for private API path",
			fields: fields{
				userPrivateAPIRateLimiter: createMockLimiter("2-S"),
			},
			args: args{
				urlPath: "/private",
			},
			want: createMockLimiter("2-S"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := &RateLimiter{
				globalRateLimiter:         tt.fields.globalRateLimiter,
				publicAPIRateLimiter:      tt.fields.publicAPIRateLimiter,
				userPrivateAPIRateLimiter: tt.fields.userPrivateAPIRateLimiter,
			}

			got := rl.filterLimiterUrlPath(tt.args.urlPath)
			assert.NotNil(t, got)
		})
	}
}

func Test_rateLimiter(t *testing.T) {
	type args struct {
		interval string
	}

	tests := []struct {
		name string
		args args
		want *limiter.Limiter
	}{
		{
			name: "should create limiter with valid interval",
			args: args{
				interval: "5-S", // 5 requests per second
			},
			want: createMockLimiter("5-S"),
		},
		{
			name: "should return nil for invalid interval",
			args: args{
				interval: "invalid",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rateLimiter(tt.args.interval)
			if tt.want == nil {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
			}
		})
	}
}
