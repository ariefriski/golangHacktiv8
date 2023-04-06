package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/novalagung/gubrak"
)

var baseURL = "https://jsonplaceholder.typicode.com"

type Value struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Tittle string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	for true {
		water, wind := RandomNumber()
		PostValue(water, wind)
		time.Sleep(15 * time.Second)
	}
}

func RandomNumber() (water, wind string) {
	randomWater := gubrak.RandomInt(1, 100)
	randomWind := gubrak.RandomInt(1, 100)

	water = strconv.Itoa(randomWater)
	wind = strconv.Itoa(randomWind)

	return
}

func PostValue(water, wind string) {
	newData := Post{
		Tittle: water,
		Body:   wind,
	}

	dataRequestJson, err := json.Marshal(newData)
	if err != nil {
		log.Fatal(err.Error())
	}

	request, err := http.NewRequest("POST", baseURL+"/posts", bytes.NewBuffer(dataRequestJson))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var client = &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer response.Body.Close()

	value := ParseAndConvertJson(response.Body)
	valueJson, _ := json.MarshalIndent(value, "", "   ")
	statusWater, statusWind := SetStatus(value.Water, value.Wind)

	fmt.Println(string(valueJson))
	fmt.Println(fmt.Sprintf("status water: %s", statusWater))
	fmt.Println(fmt.Sprintf("status wind: %s", statusWind))
}

func ParseAndConvertJson(body io.ReadCloser) Value {
	var post Post
	err := json.NewDecoder(body).Decode(&post)
	if err != nil {
		log.Fatal(err.Error())
	}
	waterRes, _ := strconv.Atoi(post.Tittle)
	windRes, _ := strconv.Atoi(post.Body)

	value := Value{
		Water: waterRes,
		Wind:  windRes,
	}

	return value
}

func SetStatus(waterInt, windInt int) (water, wind string) {
	switch {
	case waterInt <= 5:
		water = "aman"
	case waterInt >= 6 && waterInt <= 8:
		water = "siaga"
	case waterInt > 8:
		water = "bahaya"
	}

	switch {
	case windInt <= 6:
		wind = "aman"
	case windInt >= 7 && windInt <= 15:
		wind = "siaga"
	case windInt > 15:
		wind = "bahaya"
	}

	return
}