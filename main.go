package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	setupRoutes()
	config, err := ParseConfig("config")
	if err != nil{
		log.Fatal(err)
	}
	addr, ok := config["server"]
	if !ok{
		log.Fatal("Echec de la configuration du serveur")
	}
	fmt.Println("I am serving on port ",addr)
	if err := http.ListenAndServe(addr, nil); err != nil{
		fmt.Fprintln(os.Stderr,"Failed to start server:", err)
	}
}

func setupRoutes(){
	http.HandleFunc("/", homePage)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome!")
}
