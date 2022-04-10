package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		return
	}
	fullname := r.Form.Get("fullname")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	// encrypt password
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	address := r.Form.Get("address")

	// Query
	query := "UPDATE users SET"

	if fullname != "" {
		query += " fullname = '" + fullname + "',"
	}
	if username != "" {
		query += " username = '" + username + "',"
	}
	if email != "" {
		query += " email = '" + email + "',"
	}
	if encryptedPassword != "" {
		query += " password = '" + encryptedPassword + "',"
	}
	if address != "" {
		query += " address = '" + address + "',"
	}
	queryNew := query[:len(query)-1] // Delete last coma
	userId := GetIdFromCookie(r)
	queryNew += " WHERE user_id = " + strconv.Itoa(userId)

	_, errQuery := db.Exec(queryNew)

	if errQuery != nil {
		SendErrorResponse(w, 400)
	} else {
		SendSuccessResponse(w)
	}
}

func AddNewBusOrder(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		return
	}
}

func AddNewTrainOrder(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		return
	}
}

func AddNewFlightOrder(w http.ResponseWriter, r *http.Request) {

	log.Println("A")

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		log.Println("B")
		SendErrorResponse(w, 500)
		log.Println(err)
		return
	}
	userId := GetIdFromCookie(r)
	seatId := r.Form.Get("seatId")
	transactionType := r.Form.Get("transactionType")

	// Query order & update seat_status
	_, errQuery1 := db.Exec("INSERT INTO orders(user_id, seat_id, transaction_type) values (?,?,?)", userId, seatId, transactionType)
	_, errQuery2 := db.Exec("UPDATE seats SET seat_status = 1 WHERE seat_id = ", seatId)

	if errQuery1 == nil && errQuery2 == nil {
		SendSuccessResponse(w)
	} else if errQuery1 != nil {
		log.Println("C")
		log.Println(errQuery1)
		SendErrorResponse(w, 400)
		return
	} else {
		log.Println("D")
		log.Println(errQuery2)
		SendErrorResponse(w, 400)
		return

	}

}
