package repository

import "Project_2_cleanArchitecture/models/entity"

type BookRepository interface {
	GetAllBook() ([]entity.Book, error)
	GetIdBook(id uint) (entity.Book,error)
	CreateBook(book entity.Book) (entity.Book,error)
	UpdateBook(id uint,book entity.Book) (entity.Book,error)
	DeleteBook(id uint) (string,error)
}