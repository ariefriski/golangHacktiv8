package main

import (
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


type Buku struct{
	ID int `gorm:"primaryKey"`
	Name_Book string `gorm:"not null;type:varchar(255)"`
	Author string `gorm:"not null;type:varchar(250)"`
	Created_at time.Time `gorm:"type:timestamp"`
	Updated_at time.Time `gorm:"type:timestamp"`
}
// var mapBook = make(map[int]Book,0)
// var counter int

var db *gorm.DB

func main(){
	g:= gin.Default()
	var err error
	db,err=gorm.Open(postgres.Open("host= localhost port=5432 user=postgres password=123123 dbname=postgres sslmode=disable"))
	if err != nil{
		panic(err)
	}
	sqlDB,err:=db.DB()
	if err!= nil{
		panic(err)
	}
	err= sqlDB.Ping()
	if err !=nil{
		panic(err)
	}

	db.AutoMigrate(Buku{})

	g.GET("/book",getAllBookHandler)
	g.POST("/book",addBookHandler)
	g.DELETE("/book/:id",deleteBookHandler)
	g.GET("/book/:id",getIdBookHandler)
	g.PUT("/book/:id",updateBookHandler)
	g.Run(":8080")
}

func getAllBookHandler(ctx *gin.Context){
	//  books:= make([]Book,0)
	//  for _,v := range mapBook{
	// 	books = append(books, v)
	//  }
	var books = make([]Buku,0)
	//  ctx.JSON(http.StatusOK,books)
	tx:=db.Find(&books)
	if tx.Error != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error":tx.Error.Error(),
		})
		return
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

	var selectedBook Buku
	selectedBook.ID = id
	tx:=db.Find(&selectedBook)
	if tx.Error !=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error":tx.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK,selectedBook)

}

func addBookHandler(ctx *gin.Context){
	var newBook Buku
	err:=ctx.ShouldBindJSON(&newBook)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return
	}

	tx:=db.Create(&newBook)
	if tx.Error != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error":tx.Error.Error(),
		})
		return
	}
	row:=tx.Row()


	//row:=db.QueryRow(query,newBook.Title,newBook.Author,newBook.Description)

	err = row.Scan(&newBook.ID,&newBook.Name_Book,&newBook.Author,&newBook.Created_at,&newBook.Updated_at)
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

	var updateBook Buku

	err=ctx.ShouldBindJSON(&updateBook)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return
	}

	tx:=db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{
				Name:"id",
			},
			
		},
	}).Where("id=?",id).Updates(&updateBook)
	if tx.Error!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error":tx.Error.Error(),
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
			"err":err.Error(),
		})
		return
	}
	var deletedBook Buku
	deletedBook.ID = id


	tx:=db.Delete(&deletedBook)
	if tx.Error != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error":tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"message":"Book Deleted Successfully",
	})

}
