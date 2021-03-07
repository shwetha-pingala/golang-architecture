package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Person ...
type Person struct{
	First string
}


func main(){

	p1:= Person{
		First:"Sia",
	}

	p2:= Person {
		First: "Dia",
	}

	xp := []Person{p1,p2}

	bs,err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Print JSON",string(bs))

	xp2 := []Person{}
	err = json.Unmarshal(bs,&xp2)
	if err != nil {

		log.Panic(err)
	}

	fmt.Println("Back into a go datastructure",xp2)

}