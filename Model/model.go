package model

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type AccountResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Account `json:"data"`
}

type AccountsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Account `json:"data"`
}

type Game struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}

type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Game   `json:"data"`
}

type GamesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Game `json:"data"`
}

type Room struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
	Id_game   int    `json:"id_game"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type Participant struct {
	ID         int `json:"id"`
	Id_room    int `json:"id_room"`
	Id_account int `json:"id_account"`
}

type ParticipantResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Participant `json:"data"`
}

type ParticipantsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Participant `json:"data"`
}

type DetailRoom struct {
	Id           int                 `json:"id"`
	Name         string              `json:"room_name"`
	Participants []DetailParticipant `json:"participants"`
}

type RoomDesc struct {
	Room DetailRoom `json:"room"`
}

type RoomDescResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    RoomDesc `json:"data"`
}

type DetailParticipant struct {
	Id        int    `json:"id"`
	IdAccount int    `json:"id_account"`
	Username  string `json:"username"`
}
