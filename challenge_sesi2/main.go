package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type Book struct{
	ID int
	Title string
	Author string
	Desc string
}
var mapBook = make(map[int]Book,0)
var counter int

func main(){

	g := gin.Default()
	g.GET("/book",getAllBookHandler)
	g.GET("/book/:id",getIdBookHandler)
	g.POST("/book",addBookHandler)
	g.DELETE("/book/:id",deleteBookHandler)
	g.PUT("/book/:id",updateBookHandler)
	g.Run(":8080")
}

func getAllBookHandler(ctx *gin.Context){
	 books:= make([]Book,0)
	 for _,v := range mapBook{
		books = append(books, v)
	 }

	 ctx.JSON(http.StatusOK,books)
}

func getIdBookHandler(ctx *gin.Context){
	idString := ctx.Param("id")

	id,err:= strconv.Atoi(idString)
	if err !=nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
		return
	}

	v,ok:= mapBook[id]
	if !ok{
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error":"Not Found",
		})
	}

	ctx.JSON(http.StatusOK,v)

}

func addBookHandler(ctx *gin.Context){
	var newBook Book
	err:=ctx.ShouldBindJSON(&newBook)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return
	}
	newBook.ID = counter
	mapBook[counter] = newBook
	counter++
	
	ctx.JSON(http.StatusOK,"Created")

}

func updateBookHandler(ctx *gin.Context){
	idString:=ctx.Param("id")
	id,err:=strconv.Atoi(idString)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"err":err,
		})
		return
	}

	v,ok := mapBook[id]
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error":"Not Found",
		})
		return
	}
	delete(mapBook,id)
	ctx.JSON(http.StatusOK,v)

	var updateBook Book

	err=ctx.ShouldBindJSON(&updateBook)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return
	}
	updateBook.ID = v.ID
	mapBook[id]=updateBook

	ctx.JSON(http.StatusOK,"Updated")
}

func deleteBookHandler(ctx *gin.Context){
	idString:=ctx.Param("id")
	id,err:=strconv.Atoi(idString)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"err":err,
		})
		return
	}

	_,ok := mapBook[id]
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error":"Not Found",
		})
		return
	}
	delete(mapBook,id)
	ctx.JSON(http.StatusOK,"Deleted")

}
