package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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


	http.HandleFunc("/encode",foo)
	http.HandleFunc("/decode",bar)
	http.ListenAndServe(":8080",nil)

}

func foo(w http.ResponseWriter,r *http.Request){


	p3 := Person{
		First: "Sia",
	}

	err := json.NewEncoder(w).Encode(p3)
	if err!= nil {
		log.Println("Encoded into JSON data",err)
	}

}

func bar(w http.ResponseWriter,r *http.Request){


	var p3 Person
	err := json.NewDecoder(r.Body).Decode(&p3)
	if err!=nil{
		log.Println("Decoded bad data",err)
	}

	log.Println("Person",p3)
}