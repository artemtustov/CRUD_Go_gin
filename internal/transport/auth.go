package transport

import (
	"CRUD_Go_gin/internal/domain"
	"CRUD_Go_gin/internal/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	logr "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func handleSingUp(c *gin.Context) {
	var u domain.User

	logr.Info("Parsing started")
	if err := c.BindJSON(&u); err != nil {
		logr.Info("Cant parse request")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := services.SingUp(&u); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	return
}

func handleSingIn(c *gin.Context) {
	var inp domain.UserRq
	if err := c.BindJSON(&inp); err != nil {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := services.SingIn(inp.Email, inp.PasswordEnc)
	if err != nil {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(response)
	return
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			logr.Info("No token")
			c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
			return
		}

		err := services.ParseToken(headerParts[1])
		if err != nil {
			logr.Info("Bad token")
			c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
			return
		}

		c.Next()
		return
	}
}
