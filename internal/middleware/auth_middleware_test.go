package middleware

//
//import (
//	"context"
//	"errors"
//	"github.com/gin-gonic/gin"
//	"github.com/haodam/user-backend-golang/utils/auth"
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//// Mocking the auth package methods for testing purposes
//var (
//	validToken       = "valid-token"
//	invalidToken     = "invalid-token"
//	expectedSubject  = "test-subject"
//	extractedTokenFn = auth.ExtractBearerToken
//	verifyTokenFn    = auth.VerifyTokenSubject
//)
//
//func mockExtractBearerToken(c *gin.Context) (string, bool) {
//	authHeader := c.GetHeader("Authorization")
//	if authHeader == "Bearer "+validToken {
//		return validToken, true
//	}
//	return "", false
//}
//
//func mockVerifyTokenSubject(token string) (*auth.Claims, error) {
//	if token == validToken {
//		// Return a mocked claim with the test subject
//		return &auth.Claims{Subject: expectedSubject}, nil
//	}
//	return nil, errors.New("invalid token")
//}
//
//func TestAuthedMiddleware(t *testing.T) {
//	// Replace the auth functions with mocks during the test
//	auth.ExtractBearerToken = mockExtractBearerToken
//	auth.VerifyTokenSubject = mockVerifyTokenSubject
//	// Restore the original functions after the test
//	defer func() {
//		auth.ExtractBearerToken = extractedTokenFn
//		auth.VerifyTokenSubject = verifyTokenFn
//	}()
//
//	tests := []struct {
//		name       string
//		headerKey  string
//		headerVal  string
//		wantStatus int
//		wantBody   string
//	}{
//		{
//			name:       "Valid token provided",
//			headerKey:  "Authorization",
//			headerVal:  "Bearer valid-token",
//			wantStatus: http.StatusOK,
//			wantBody:   "Success", // Our test endpoint response represents "Success".
//		},
//		{
//			name:       "Missing token",
//			headerKey:  "",
//			headerVal:  "",
//			wantStatus: http.StatusUnauthorized,
//			wantBody:   `{"code":40001,"description":"","err":"Unauthorized"}`,
//		},
//		{
//			name:       "Invalid token provided",
//			headerKey:  "Authorization",
//			headerVal:  "Bearer invalid-token",
//			wantStatus: http.StatusUnauthorized,
//			wantBody:   `{"code":40001,"description":"","err":"invalid token"}`,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Setup Gin
//			gin.SetMode(gin.TestMode)
//			r := gin.New()
//			r.Use(AuthedMiddleware())
//
//			// Add a simple test route
//			r.GET("/test", func(c *gin.Context) {
//				// Extract the context value set by the middleware
//				subjectUUID := c.Request.Context().Value("subjectUUID").(string)
//				assert.Equal(t, expectedSubject, subjectUUID, "subjectUUID should match expectedSubject")
//				c.String(http.StatusOK, "Success")
//			})
//
//			// Create the HTTP request
//			req := httptest.NewRequest(http.MethodGet, "/test", nil)
//			if tt.headerKey != "" {
//				req.Header.Set(tt.headerKey, tt.headerVal)
//			}
//			// Create a response recorder
//			w := httptest.NewRecorder()
//
//			// Perform the request
//			r.ServeHTTP(w, req)
//
//			// Assert the expected response
//			assert.Equal(t, tt.wantStatus, w.Code)
//			assert.JSONEq(t, tt.wantBody, w.Body.String(), "Response body should match")
//		})
//	}
//}
