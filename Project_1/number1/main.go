package main

import (
	"fmt"
	"sync"
)

func main() {
var data1 interface{}
var data2 interface{}

data1 = []string{"coba1","coba2","coba3"}

data2 = []string{"bisa1","bisa2","bisa3"} 

var wg sync.WaitGroup
//var m sync.Mutex

wg.Add(8)

for i:=1;i<=4;i++{
	
	go cobaPrint1(data1,&wg,i)
	go cobaPrint2(data2,&wg,i)
	
}

wg.Wait()

}

func cobaPrint1(data1 interface{},wg *sync.WaitGroup,i int){
	fmt.Println(data1,i)
	wg.Done()
}

func cobaPrint2(data2 interface{},wg *sync.WaitGroup,i int){
	fmt.Println(data2,i)
	wg.Done()
}


