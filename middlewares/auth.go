package middlewares

import (
	"net/http"
	"strings"

	"github.com/adrianramadhan/final-task-pbi-rakamin-fullstack-AdrianRamadhan/helpers"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": "Failed", "Message": "Token missing"})
		return
	}

	accessToken := strings.Split(authHeader, " ")[1]
	claims, err := helpers.ReadToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	c.Set("reqID", claims.ID)

	c.Next()
}