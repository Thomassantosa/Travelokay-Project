package controllers

import (
	"log"
	"net/http"
)

func UpdateUsers(w http.ResponseWriter, r *http.Request) {

	// connect to database
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
	address := r.Form.Get("address")

	// Query
	query := "UPDATE users SET"

	if fullname != "" {
		query += " fullname='" + fullname + "',"
	}
	if username != "" {
		query += " username='" + username + "',"
	}
	if email != "" {
		query += " email='" + email + ","
	}
	if password != "" {
		query += " password=" + password + ","
	}
	if address != "" {
		query += " address=" + address + ","
	}
	queryNew := query[:len(query)-1] // Delete last coma
	queryNew += " WHERE email=" + email + "'"

	result, errQuery := db.Exec(queryNew)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			SendErrorResponse(w, 400)
		} else {
			SendSuccessResponse(w)
			log.Println(email)
		}
	} else {
		SendErrorResponse(w, 400)
	}
}
