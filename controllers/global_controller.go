package controllers

import (
	"log"
	"net/http"
	"os"

	models "github.com/Travelokay-Project/models"
	"github.com/joho/godotenv"
	// "github.com/Travelokay-Project/models"
)

func LoadEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Login(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	email := r.FormValue("Email")
	password := r.FormValue("Password")

	row := db.QueryRow("SELECT * FROM users WHERE email=? AND password=?", email, password)
	var user models.User
	if err := row.Scan(&user.ID, &user.Fullname, &user.Username, &user.Email, &user.Password, &user.Address, &user.UserType, &user.PartnerType, &user.CompanyName, &user.DateCreated); err != nil {
		log.Println(row)
		SendErrorResponse(w, 400)
		log.Print(err)
	} else {
		GenerateToken(w, user.ID, user.Username, user.UserType)
		SendSuccessResponse(w)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	ResetUserToken(w)
	SendSuccessResponse(w)
}
