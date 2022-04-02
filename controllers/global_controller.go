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
func InsertUsers(w http.ResponseWriter, r *http.Request) {

	// connect to database
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	fullname := r.Form.Get("fullname")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	address := r.Form.Get("address")
	user_type := r.Form.Get("user_type")
	partner_type := r.Form.Get("partner_type")
	company_name := r.Form.Get("company_name")
	date_created := r.Form.Get("date_created")

	_, errQuery := db.Exec("INSERT INTO users(fullname,username,email,password,address,user_type,partner_type,company_name,date_created) values (?,?,?,?,?,?,?,?,?)", fullname, username, email, password, address, user_type, partner_type, company_name, date_created)

	if errQuery == nil {
		SendSuccessResponse(w)
	} else {
		SendErrorResponse(w, 400)
	}

	db.Close()
}
func UpdateUsers(w http.ResponseWriter, r *http.Request) {

	// connect to database
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	fullname := r.Form.Get("fullname")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	address := r.Form.Get("address")
	user_type := r.Form.Get("user_type")
	partner_type := r.Form.Get("partner_type")
	company_name := r.Form.Get("company_name")
	date_created := r.Form.Get("date_created")

	result, errQuery := db.Exec("UPDATE users SET fullname=?, username=?, password=?, address=?, user_type=?, partner_type=?, company_name=?, date_created=? WHERE email=?", fullname, username, password, address, user_type, partner_type, company_name, date_created, email)

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

	db.Close()
}
