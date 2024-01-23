package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rifansyah/go-crud/app/initializers"
	"github.com/rifansyah/go-crud/app/models"
)

func RequireAuth() gin.HandlerFunc {
	return func (c *gin.Context)  {
		versioningPrefix := "/api/v1"
		excludedPaths := []string{"/signup", "/login"}
		for _, path := range excludedPaths {
			if strings.HasPrefix(c.Request.URL.Path, fmt.Sprintf("%s%s", versioningPrefix, path)) {
				c.Next()
				return;
			}
		}

		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or missing bearer token",
			})
			return;
		}

		tokenStr := strings.TrimPrefix(bearerToken, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", t.Header["Authorization"])
			}
			return []byte(os.Getenv("TOKEN_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return;
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if isExpired := float64(time.Now().Unix()) > claims["exp"].(float64); isExpired {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Token expired",
				})
				return;
			}

			uid := claims["uid"]
			var user models.User
			if result := initializers.DB.First(&user, uid); result.Error != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return;
		}
	}
}
