package main

import "fmt"

func countMap(kata string) map[string]int {
	hitung := make(map[string]int)
	for _, v := range kata {
		_, ok := hitung[string(v)]
		if ok {
			hitung[string(v)] += 1
		} else {
			hitung[string(v)] = 1
		}
	}
	return hitung
}

func LoopingAndCount (word string){
	for _,v := range word {
		fmt.Printf("%c\n",v)
	}
	map1 := countMap(word)
	fmt.Println(map1)
}

func main() {
	const kata = "selamat malam"

	LoopingAndCount(kata)
}