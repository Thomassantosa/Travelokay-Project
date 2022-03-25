package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	models "github.com/Travelokay-Project/model"
	// "github.com/Travelokay-Project/models"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// get data
	err := r.ParseForm()
	if err != nil {
		// error response
		SendErrorResponse(w, 500)
		log.Println(err)
		return
	}

	email := r.FormValue("email")
	password := r.Form.Get("password")

	query := "SELECT * FROM users WHERE email = '" + email + "' AND password = '" + password + "'"
	rows, err := db.Query(query)
	if err != nil {
		// error response
		SendErrorResponse(w, 500)
		log.Println(err)
		return
	}

	// if !rows.Next() {
	// 	// error response
	// 	SendErrorResponse(w, 400)
	// 	log.Println(err)
	// 	return
	// }

	var user models.User
	var users []models.User
	// var user User
	// var users []User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password)
		if err != nil {
			// error response
			SendErrorResponse(w, 500)
			log.Println(err)
			return
		} else {
			users = append(users, user)
		}
	}

	// success response
	var response models.UsersResponse
	response.Status = 200
	response.Message = "Request success"
	response.Data = users

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("Login Success")
}
