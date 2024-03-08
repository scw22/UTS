package main

import (
	"UTS/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/Rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/users", controllers.InsertRoom).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
