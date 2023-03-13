package main

import (
	"fmt"
)

func main() {
	var i int = 21
	var j bool = true
	russia := 'Я'
	var k float32 = 123.456
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println("%")
	fmt.Println(j)
	fmt.Printf("%b\n", russia)
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%U\n", russia)
	fmt.Printf("%f\n", 123.456)
	fmt.Printf("%e\n", k)

}
