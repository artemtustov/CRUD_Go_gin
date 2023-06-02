package transport

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Books API
// @version 1.0
// @description REST API for Books App

// @host localhost:8080
// @BasePath /book

func NewServer() *gin.Engine {
	r := gin.New()
	g := r.Group("/book")
	{
		g.Use(authMiddleware())

		g.POST("/add", handleCreateBook)
		g.POST("/update", handleUpdateBook)
		g.GET("/:id", handleGetBook)
		g.DELETE("/delete/:id", handleDeleteBook)
	}

	g2 := r.Group("/auth")
	{
		g2.POST("/singin", handleSingIn)
		g2.POST("/singup", handleSingUp)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func Run(r *gin.Engine) error {
	return r.Run("localhost:8080")
}
