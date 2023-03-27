package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)


type Book struct{
	ID int
	Title string
	Author string
	Description string
}
var mapBook = make(map[int]Book,0)
var counter int

var db *sql.DB

func main(){
	var err error
	db,err=sql.Open("postgres","host= localhost port=5432 user=postgres password=123123 dbname=postgres sslmode=disable")
	if err != nil{
		panic(err)
	}

	err = db.Ping()
	if err!= nil{
		panic(err)
	}
	fmt.Println(db)
	g := gin.Default()
	g.GET("/book",getAllBookHandler)
	g.GET("/book/:id",getIdBookHandler)
	 g.POST("/book",addBookHandler)
	g.DELETE("/book/:id",deleteBookHandler)
	g.PUT("/book/:id",updateBookHandler)
	g.Run(":8080")
}

func getAllBookHandler(ctx *gin.Context){
	//  books:= make([]Book,0)
	//  for _,v := range mapBook{
	// 	books = append(books, v)
	//  }

	//  ctx.JSON(http.StatusOK,books)
	query := "SELECT * FROM book"
	rows,err := db.Query(query)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	var books = make([]Book,0)
	for rows.Next(){
		var book Book

		err:=rows.Scan(&book.ID,&book.Title,&book.Author,&book.Description)
		if err !=nil{
			ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		books = append(books, book)
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

	// v,ok:= mapBook[id]
	// if !ok{
	// 	ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
	// 		"error":"Not Found",
	// 	})
	// }
	var selectedBook Book

	query:= "select * from book where id=$1" // harus pakai angka
	rows:=db.QueryRow(query,id)
	err=rows.Scan(&selectedBook.ID,&selectedBook.Title,&selectedBook.Author,&selectedBook.Description)
    if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error":err,
		})
		return
	}

	ctx.JSON(http.StatusOK,selectedBook)

}

func addBookHandler(ctx *gin.Context){
	var newBook Book
	err:=ctx.ShouldBindJSON(&newBook)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return
	}
	// newBook.ID = counter
	// mapBook[counter] = newBook
	// counter++
	query := "insert into book (title,author,description) VALUES($1,$2,$3) returning *"
	
	// result,err:=db.Exec(query,newBook.Title,newBook.Author,newBook.Description)
	// if err != nil{
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
	// 		"error":err.Error(),
	// 	})
	// 	return
	// }
	row:=db.QueryRow(query,newBook.Title,newBook.Author,newBook.Description)
	
	err = row.Scan(&newBook.ID,&newBook.Title,&newBook.Author,&newBook.Description)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error" : err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK,newBook)

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

	var updateBook Book

	err=ctx.ShouldBindJSON(&updateBook)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return
	}

	query:= "update book set title=$1, author=$2, description=$3 WHERE id=$4 returning id"
	rows:=db.QueryRow(query,updateBook.Title,updateBook.Author,updateBook.Description,id)
	err=rows.Scan(&updateBook.ID)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error:":err,
		})
		return
	}

	ctx.JSON(http.StatusOK,updateBook)
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
	var deletedBook Book

	// _,ok := mapBook[id]
	// if !ok {
	// 	ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
	// 		"error":"Not Found",
	// 	})
	// 	return
	// }
	// delete(mapBook,id)
	//Newbook.ID = id
	query := "delete from book where id=$1 returning *" //tidak boleh pake i
	rows:=db.QueryRow(query,id)
	// if err!= nil {
	// 	ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
	// 		"error":"Not Found",
	// 	})
	// 	return
	// }
	err= rows.Scan(&deletedBook.ID,&deletedBook.Title,&deletedBook.Author,&deletedBook.Description)
	if err!= nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error":err,
		})
		return
	}
	ctx.JSON(http.StatusOK,deletedBook)

}
