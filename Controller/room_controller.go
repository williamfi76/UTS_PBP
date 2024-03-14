package controller

import (
	"fmt"
	"log"
	"net/http"
	m "uts_pbp/Model"
)

func GetAllRoomsForGame(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * from rooms WHERE id_game = ?"

	id := r.FormValue("id")

	if id == "" {
		sendErrorResponse(w, "Data yang diberikan tidak lengkap")
		return
	}

	rows, err := db.Query(query, id)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Data gagal di temukan")
		return
	}

	var room m.Room
	var rooms []m.Room

	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.Room_name, &room.Id_game); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	if len(rooms) == 0 {
		sendErrorResponse(w, "Error: Data kosong")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	sendRoomsSuccessResponse(w, rooms, "Data berhasil ditemukan")
}

func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	inputRoomId := r.FormValue("id")

	if inputRoomId == "" {
		sendErrorResponse(w, "Error: Data belum Lengkap")
		return
	}

	query := "SELECT r.id, r.room_name, r.id_game, p.id, p.id_room, p.id_account, a.id, a.username from rooms r INNER JOIN participants p ON r.id = p.id_room INNER JOIN accounts a ON p.id_account = a.id where r.id = ?"

	rows, err := db.Query(query, inputRoomId)

	if err != nil {
		sendErrorResponse(w, "Error: Data gagal di eksekusi")
	}

	var roomDesc m.RoomDesc
	var detRoom m.DetailRoom
	var detPart m.DetailParticipant
	dummy := "Hello"
	for rows.Next() {

		if err := rows.Scan(&detRoom.Id, &detRoom.Name, &dummy, &dummy, &dummy, &dummy, &detPart.IdAccount, &detPart.Username); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Tidak dapat mengaplikasikan data kedalam JSON")
			return
		} else {

			detRoom.Participants = append(detRoom.Participants, detPart)
		}
	}

	roomDesc.Room = detRoom

	if len(roomDesc.Room.Participants) == 0 {
		sendErrorResponse(w, "Error: Data not Found")
		return
	}

	sendRoomDescSuccessResponse(w, roomDesc, "Data berhasil ditemukan")

}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	idRoom := r.FormValue("id_room")
	idAcc := r.FormValue("id_acc")

	maxPart, err := db.Query("SELECT g.max_player FROM rooms r INNER JOIN games g ON r.id_game = g.id INNER JOIN participants p ON p.id_room = r.id WHERE r.id = ? LIMIT 1", idRoom)
	if err != nil {
		log.Println("Error: Cant Find MaxpLayer")
		sendErrorResponse(w, "Tidak bisa menemukan Max Player")
		return
	}
	currPart, err := db.Query("SELECT count(*) FROM rooms r INNER JOIN games g ON r.id_game = g.id INNER JOIN participants p ON p.id_room = r.id WHERE r.id = ?", idRoom)
	if err != nil {
		log.Println("Error parsing form data:", err)
		http.Error(w, "Gagal mengambil Data", http.StatusInternalServerError)
		return
	}

	var count = 0
	var maxPlayer = 0

	for currPart.Next() {
		if err := currPart.Scan(&count); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Gagal mengambil jumlah Partisipan saat ini")
			return
		}
	}

	for maxPart.Next() {
		if err := maxPart.Scan(&maxPlayer); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Gagal mengambil jumlah maksimal Player")
			return
		}
	}

	fmt.Println(count, " ", maxPlayer)

	if count == maxPlayer {
		sendErrorResponse(w, "Tidak bisa join room sudah memenuhi kapasitas")
		return
	}

	query := "INSERT INTO participants(id_room, id_account) VALUES (?, ?)"

	result, err := db.Query(query, idRoom, idAcc)
	if result.Err() != nil {
		sendErrorResponse(w, "Gagal memeasukan account ke dalam partisipan")
	}

	sendRoomsSuccessResponse(w, nil, "Accountberhasil dimasukan kedalam Partisipan")

}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

}
