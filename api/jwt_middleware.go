package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/shipments-service/config"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				return nil, fmt.Errorf("Invalid token", token.Header["alg"])
			}
			return []byte(config.ServerConf.Jwt.Key), nil
		})
		if err != nil {
			log.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
