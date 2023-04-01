package controller

import (
	"Project_2_cleanArchitecture/models/entity"
	"Project_2_cleanArchitecture/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookControllerImpl struct {
	BookRepository repository.BookRepository
}

func NewBookController(bookRepository repository.BookRepository)BookController{
	return &BookControllerImpl{
		BookRepository: bookRepository,
	}
}


func (controller *BookControllerImpl)GetAllBook(ctx *gin.Context){
	result,err:=controller.BookRepository.GetAllBook()
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,result)
}
func (controller *BookControllerImpl)GetIdBook(ctx *gin.Context){
	id:=ctx.Param("id")
	
	idBook,err:=strconv.Atoi(id)
	if err !=nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	result,err:=controller.BookRepository.GetIdBook(uint(idBook))
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,result)
}
func (controller *BookControllerImpl)CreateBook(ctx *gin.Context){
	CreateBook := entity.Book{}

	err:=ctx.ShouldBindJSON(&CreateBook)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Erorr":err.Error(),
		})
		return
	}

	result,err:=controller.BookRepository.CreateBook(CreateBook)
	if err !=nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,result)
}
func (controller *BookControllerImpl)UpdateBook(ctx *gin.Context){
	id := ctx.Param("id")
	updateBook := entity.Book{}
	idBook,err:=strconv.Atoi(id)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	err= ctx.ShouldBindJSON(&updateBook)
	if err !=nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	result,err:=controller.BookRepository.UpdateBook(uint(idBook),updateBook)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,result)


}
func (controller *BookControllerImpl)DeleteBook(ctx *gin.Context){
	id:= ctx.Param("id")
	idBook,err:=strconv.Atoi(id)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	result,err:=controller.BookRepository.DeleteBook(uint(idBook))
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":result,
	})
}
