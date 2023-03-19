package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct{
	Name string
	Alamat string
	Pekerjaan string
	Alasan string
}

func main(){

	//persons:= make([]Person,0)

	persons :=[]Person {
		Person{
			Name:"Arief",
			Alamat:"Bekasi",
			Pekerjaan: "Programmer",
			Alasan: "Ingin menambah skill baru",
		},
		Person{
			Name:"Riski",
			Alamat:"Jakarta",
			Pekerjaan: "Accounting",
			Alasan: "Ingin Belajar",
		},
		Person{
			Name:"Indra",
			Alamat:"Ambon",
			Pekerjaan: "Koki",
			Alasan: "Ingin mengisi waktu luang",
		},
		Person{
			Name:"Pratama",
			Alamat:"Aceh",
			Pekerjaan: "Polisi",
			Alasan: "Ingin mengasah logika",
		},
		Person{
			Name:"Rudi",
			Alamat:"Karawang",
			Pekerjaan: "Satpam",
			Alasan: "Ingin pindah karir",
		},
		Person{
			Name:"Nurwanto",
			Alamat:"Purwakarta",
			Pekerjaan: "Admin",
			Alasan: "Ingin pintar",
		},
		Person{
			Name:"James",
			Alamat:"Australia",
			Pekerjaan: "Gamer",
			Alasan: "Ingin membuat game",
		},
		Person{
			Name:"Arthur",
			Alamat:"Inggris",
			Pekerjaan: "Singer",
			Alasan: "Ingin sing a song",
		},
		Person{
			Name:"Morgan",
			Alamat:"Africa",
			Pekerjaan: "Slave",
			Alasan: "Ingin bebas",
		},
		Person{
			Name:"Freeman",
			Alamat:"Norwegia",
			Pekerjaan: "Bos",
			Alasan: "Ingin mengajar",
		},
	}

	printData(persons)

}

func printData(persons []Person){
	var sum int64 = 1
	for i:= 1;i<len(os.Args);i++{
		n,_ := strconv.ParseInt(os.Args[i],10,0)
		sum += n
	}

	if sum == 0|| sum ==1{
		fmt.Println(persons[0])
	}else if sum >= int64(len(persons)){
		fmt.Println("Data Tidak ada")
	}else{
		fmt.Println(persons[sum])
	}
}


