package controllers

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
type Games struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	MaxPlayer int    `json:"max_player"`
}
type Rooms struct {
	Id     int    `json:"id"`
	Name   string `json:"room_name"`
	IdGame int    `json:"id_game"`
}
type Participants struct {
	ID        int `json:"id"`
	IdRoom    int `json:"id_room"`
	IdAccount int `json:"id_account"`
}
type Response struct {
	Message string      `json:"error"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}
type ErrorResponse struct {
	Message string `json:"error"`
	Status  int    `json:"status"`
}
