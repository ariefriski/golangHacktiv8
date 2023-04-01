package route

import (
	"Project_2_cleanArchitecture/controller"
	"Project_2_cleanArchitecture/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookRoute(router *gin.Engine, db *gorm.DB) {
	bookRepository := repository.NewBookRepository(db)
	bookController := controller.NewBookController(bookRepository)

	router.POST("/books", bookController.CreateBook)
	router.GET("/books", bookController.GetAllBook)
	router.GET("/books/:id", bookController.GetIdBook)
	router.PUT("/books/:id", bookController.UpdateBook)
	router.DELETE("/books/:id", bookController.DeleteBook)
}