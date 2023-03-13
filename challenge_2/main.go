package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ", i)
		if i > 3 {
			for j := 0; j <= 10; j++ {
				if j == 5 {
					const karakter = "САШАРВО"
					for index, runeValue := range karakter {
						fmt.Printf("character %#U starts at byte position %d\n", runeValue, index)
					}
					continue
				}
				fmt.Println("Nilai j = ", j)
			}
			break
		}
	}
}
