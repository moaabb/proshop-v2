package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moaabb/proshop-v2/asset_svc/infra/config"
	"go.uber.org/zap"
)

type AuthResult struct {
	UserId  uint `json:"userId"`
	IsAdmin bool `json:"isAdmin"`
}

type AuthMiddleware struct {
	l   *zap.Logger
	cfg *config.Config
}

func NewAuthMiddleware(l *zap.Logger, cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		l,
		cfg,
	}
}

func (am *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s/v1/auth", am.cfg.AuthSvcUrl), nil)
		cookie, err := c.Cookie("jwt")
		if err != nil {
			am.l.Error("could not retrieve auth token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials",
			})
			return
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cookie))
		am.l.Info(fmt.Sprintf("making request to auth Service: %v, %v, %v", req.Method, req.Body, req.URL))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			am.l.Error("error authenticating", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "could not reach auth service",
			})
			return
		}

		var a AuthResult
		json.NewDecoder(resp.Body).Decode(&a)
		am.l.Info(fmt.Sprintf("response from auth Service: {statusCode: %v, body: %v}", resp.StatusCode, a))
		if resp.StatusCode != 200 {
			c.AbortWithStatusJSON(resp.StatusCode, gin.H{
				"error": "invalid credentials",
			})
			return
		}

		c.Set("userId", a.UserId)
		c.Set("isAdmin", a.IsAdmin)

		c.Next()
	}
}

func (am *AuthMiddleware) Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetBool("isAdmin") {
			c.Next()
			return
		}

		am.l.Error("user is not an admin")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "user cannot access this resource",
		})
	}
}
