package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var baseURL = "http://localhost:8080"

type Book struct{
	ID int `json:"id"`
	NameBook string `json:"name_book"`
	Author string `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	//get()
	post()
}

func get() {
	client := http.Client{}
	url:= fmt.Sprintf("%s/books",baseURL)

	request,err:=http.NewRequest(http.MethodGet,url,http.NoBody)
	if err != nil{
		panic(err)
	}

	response,err:=client.Do(request)
	if err !=nil{
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK{
		fmt.Println("Status Not OK")
		return
	}

	data,err:=io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	var books = make([]Book,0)

	err = json.Unmarshal(data,&books)
	if err != nil{
		panic( err)
	}

	fmt.Println(books)
	fmt.Println(len(books))
	fmt.Println(books[0])

}

func post() {
	var newBook = Book{
		NameBook: "Buku ajaib XXX",
		Author: "Arief",
	}

	data,err := json.Marshal(newBook)
	if err != nil{
		panic(err)
	}
	body:=bytes.NewBuffer(data)

	client := http.Client{}

	url:= fmt.Sprintf("%s/books",baseURL)
	

	request,err := http.NewRequest(http.MethodPost,url,body)
	if err != nil{
		panic(err)
	}
	request.Header.Set("Content-Type","application/json")

	response,err:=client.Do(request)
	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK{
		fmt.Println("Error")
		return
	}

	dataResponse,err:=io.ReadAll(response.Body)
	if err !=nil{
		panic(err)
	}

	var createdBook Book
	err= json.Unmarshal(dataResponse,&createdBook)
	if err !=nil{
		panic(err)
	}

	fmt.Println(createdBook)

}