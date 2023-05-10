package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dilyara4949/pq_daq/app/services"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		}
		if authHeader == "" {
			response := Response {
				Status:  false,
				Message: "empty",
				Error:   nil,
				Data:    "headre is empty",
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err, token)
			response := Response {
				Status:  false,
				Message: "errror",
				Error:   nil,
				Data:    err.Error(),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	})
}
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}
