package transport

import (
	"CRUD_Go_gin/internal/domain"
	"CRUD_Go_gin/internal/services"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"strconv"
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
		g.POST("/add", handleCreateBook)
		g.POST("/update", handleUpdateBook)
		g.GET("/:id", handleGetBook)
		g.DELETE("/delete/:id", handleDeleteBook)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func Run(r *gin.Engine) error {
	return r.Run("localhost:8080")
}

// @Summary Add book in database
// @Accept json
// @Produce json
// @Success 200
// @Router /add [post]
func handleCreateBook(c *gin.Context) {
	var b domain.Book

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateBook(&b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}

// @Summary Get book from database
// @Produce json
// @Param        id   path      string  true  "Account ID"
// @Success      200  {object}  domain.Book
// @Failure      204  header  string  "No such book"
// @Failure      400  header  string  "Incorrect request"
// @Failure      404  header  string  "Page not found"
// @Router /{id} [get]
func handleGetBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := services.GetBook(idInt)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": book})
	return
}

func handleUpdateBook(c *gin.Context) {
	var b domain.Book

	if err := c.ShouldBindJSON(&b); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := services.UpdateBook(&b)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	return
}

func handleDeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.DeleteBook(idInt)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	return
}
