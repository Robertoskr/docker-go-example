package main

import (
	"log"
	"fmt"
	"net/http"
)

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome to the go web api\n")
	fmt.Println("Endpoint hit: HOMEPAGE")	
}

func aboutMe(response http.ResponseWriter, request *http.Request){
	who := "Roberto Kalinovskyy"
	fmt.Fprintf(response, "a little bit about Roberto Kalinovskyy")
	fmt.Println("Endpoint hit: ", who)
}

func main(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	log.Fatal(http.ListenAndServe(":8000", nil))	
}