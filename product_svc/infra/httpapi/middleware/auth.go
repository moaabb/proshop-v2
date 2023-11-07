package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthResult struct {
	UserId string `json:"userId"`
}

func Authenticate(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, _ := http.NewRequest("POST", "http://auth:8080/v1/auth", nil)
		cookie, err := c.Cookie("jwt")
		if err != nil {
			l.Error("could not retrieve auth token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials",
			})
			return
		}

		req.Header.Set("Authorization", cookie)
		l.Info(fmt.Sprintf("making request to auth Service: %v, %v, %v", req.Method, req.Body, req.URL))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			l.Error("error authenticating", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "could not reach auth service",
			})
			return
		}

		var a AuthResult
		json.NewDecoder(resp.Body).Decode(&a)
		l.Info(fmt.Sprintf("response from auth Service: {statusCode: %v, body: %v}", resp.StatusCode, a))
		if resp.StatusCode != 200 {
			c.AbortWithStatusJSON(resp.StatusCode, gin.H{
				"error": "invalid credentials",
			})
			return
		}

		c.Set("userId", a.UserId)

		c.Next()
	}
}
