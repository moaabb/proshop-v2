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
		req.Header.Set("Authorization", c.GetHeader("Authorization"))
		l.Info(fmt.Sprintf("making request to auth Service: %v, %v, %v", req.Method, req.Body, req.URL))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			l.Error("error authenticating", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "could not reach auth service",
			})
			return
		}

		if resp.StatusCode != 200 {
			l.Info(fmt.Sprintf("response from auth Service: {statusCode: %v, body: %v}", resp.StatusCode, resp.Body))
			c.AbortWithStatusJSON(resp.StatusCode, gin.H{
				"error": "invalid credentials",
			})
			return
		}

		var a AuthResult
		json.NewDecoder(resp.Body).Decode(&a)

		c.Set("userId", a.UserId)

		c.Next()
	}
}
