package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"
)

func UpdatePartner(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	fullname := r.FormValue("fullname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	address := r.FormValue("address")
	partnerType := r.FormValue("partnerType")
	companyName := r.FormValue("companyName")

	// encrypt password
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

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
	if partnerType != "" {
		query += " address = '" + partnerType + "',"
	}
	if companyName != "" {
		query += " address = '" + companyName + "',"
	}

	queryNew := query[:len(query)-1] // Delete last coma
	partnerId := GetIdFromCookie(r)
	queryNew += " WHERE user_id = " + strconv.Itoa(partnerId)

	_, errQuery := db.Exec(queryNew)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
	} else {
		SendSuccessResponse(w)
		log.Println("(SUCCESS)\t", "Update partner request")
	}
}
