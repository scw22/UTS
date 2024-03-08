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

	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/rooms", controllers.GetDetailRoom).Methods("GET")
	router.HandleFunc("/rooms", controllers.InsertRoom).Methods("POST")
	router.HandleFunc("/rooms", controllers.DeleteRooms).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
