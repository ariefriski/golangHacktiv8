package repository

import (
	"Project_2_cleanArchitecture/models/entity"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}
 
func NewBookRepository(db *gorm.DB) BookRepository{
	return &BookRepositoryImpl{
		DB: db,
	}
}

func (repository *BookRepositoryImpl) GetAllBook() ([]entity.Book,error){
	listBook:= []entity.Book{}
	tx:=repository.DB.Find(&listBook)
	if tx.Error != nil{
		return  nil,tx.Error
	}	
	return listBook,nil
}

func (repository *BookRepositoryImpl) GetIdBook(id uint) (entity.Book,error){
	book:= entity.Book{}
	tx:=repository.DB.Find(&book,"id =?",id)
	if tx.Error != nil{
		return entity.Book{},tx.Error
	}
	return book,nil
}

func (repository *BookRepositoryImpl) CreateBook(book entity.Book) (entity.Book,error){
	NewBook:= entity.Book{
		Title: book.Title,
		Author: book.Author,
	}

	tx:=repository.DB.Create(&NewBook)
	if tx.Error!=nil{
		return entity.Book{},tx.Error
	}
	return NewBook,nil
}

func (repository *BookRepositoryImpl) UpdateBook(id uint,book entity.Book) (entity.Book,error){
	UpdateBook := entity.Book{}

	tx:=repository.DB.First(&UpdateBook,"id=?",id)
	if tx.Error !=nil{
		return entity.Book{},tx.Error
	}

	tx=repository.DB.Model(&UpdateBook).Updates(entity.Book{Title: book.Title,Author: book.Author})
	if tx.Error != nil{
		return entity.Book{},tx.Error
	}
	return UpdateBook,nil


}

func (repository *BookRepositoryImpl) DeleteBook(id uint) (string,error){
	DeleteBook := entity.Book{}
	tx:=repository.DB.Delete(&DeleteBook,id)
	if tx.Error != nil{
		return "Data gagal dihapus",nil
	}
	return "Data Berhasil Dihapus",nil
}