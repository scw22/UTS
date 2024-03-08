package controllers

import (
	m "UTS/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DataResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	var response m.Response
	response.Status = status
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func MessageResponse(w http.ResponseWriter, status int, message string) {
	var response m.ErrorResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM rooms"

	name := r.URL.Query()["name"]
	id_game := r.URL.Query()["id_game"]
	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name='" + name[0] + "'"
	}
	if id_game != nil {
		if name != nil {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " id_game='" + id_game[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		MessageResponse(w, 400, "Something went wrong, please try again.")
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.Name, &room.IdGame); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	if len(rooms) < 5 {
		DataResponse(w, 200, "Success", rooms)
	} else {
		MessageResponse(w, 400, "Error, Incorrect Array Size")
	}
}
func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM rooms"

	name := r.URL.Query()["name"]
	id_game := r.URL.Query()["id_game"]
	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name='" + name[0] + "'"
	}
	if id_game != nil {
		if name != nil {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " id_game='" + id_game[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		MessageResponse(w, 400, "Something went wrong, please try again.")
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.Name, &room.IdGame); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	if len(rooms) < 5 {
		DataResponse(w, 200, "Success", rooms)
	} else {
		MessageResponse(w, 400, "Error, Incorrect Array Size")
	}
}
func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		MessageResponse(w, 400, "failed")
	}
	name := r.Form.Get("id")
	IdGame := r.Form.Get("id_game")

	_, errQuery := db.Exec("INSERT INTO users (name, id_game) values (?,?)",
		name,
		IdGame,
	)

	if errQuery == nil {
		MessageResponse(w, 200, "Success")
	} else {
		fmt.Println(errQuery)
		MessageResponse(w, 400, "Insert Failed")
	}
}

func DeleteRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	IdRoom := vars["room_id"]

	_, errQuery := db.Exec("DELETE FROM rooms WHERE id=?", IdRoom)

	if errQuery == nil {
		MessageResponse(w, 200, "Success")
	} else {
		fmt.Println(errQuery)
		MessageResponse(w, 400, "Delete Failed")
	}
}
