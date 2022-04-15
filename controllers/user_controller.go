package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	models "github.com/Travelokay-Project/models"
	gomail "gopkg.in/gomail.v2"
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
		log.Println(errQuery1)
		SendErrorResponse(w, 400)
		return
	} else {
		log.Println(errQuery2)
		SendErrorResponse(w, 400)
		return
	}
}

func AddNewFlightOrder(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(w, 500)
		log.Println(err)
		return
	}
	userId := GetIdFromCookie(r)
	seatId := r.Form.Get("seatId")
	transactionType := r.Form.Get("transactionType")

	// Check seat_id
	row := db.QueryRow("SELECT seat_status FROM seats WHERE seat_id=?", seatId)
	var seatType int
	if err := row.Scan(&seatType); err != nil {
		SendErrorResponse(w, 400)
		log.Print(err)
		return
	}

	if seatType == 0 {

		// Query order & update seat_status
		_, errQuery1 := db.Exec("INSERT INTO orders(user_id, seat_id, transaction_type) values (?,?,?)", userId, seatId, transactionType)
		_, errQuery2 := db.Exec("UPDATE seats SET seat_status = 1 WHERE seat_id = ?", seatId)

		if errQuery1 == nil && errQuery2 == nil {
			SendSuccessResponse(w)
		} else if errQuery1 != nil {
			log.Println(errQuery1)
			SendErrorResponse(w, 400)
			return
		} else {
			log.Println(errQuery2)
			SendErrorResponse(w, 400)
			return
		}
	} else {
		SendMessageOnlyResponse(w, "Seat already booked")
		return
	}
}

func GetUserOrder(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from cookie
	userId := GetIdFromCookie(r)

	// Query
	rows, errQuery := db.Query("SELECT * FROM orders WHERE user_id=?", userId)
	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 500)
		return
	}

	// Set value
	var order models.Orders
	var orders []models.Orders

	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.UserID, &order.SeatID, &order.RoomID, &order.TourScheduleID, &order.OrderDate, &order.OrderStatus, &order.TransactionType); err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}
		orders = append(orders, order)
	}

	// Response
	var response models.OrdersResponse

	if len(orders) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 204)
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = orders

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get user order request")
	}
}

func RequestRefund(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	orderId := r.FormValue("orderId")

	// Get value from cookie
	userId := GetIdFromCookie(r)

	// Query
	_, errQuery := db.Exec("UPDATE orders SET order_status = 'refund' WHERE order_id = ? AND user_id = ?", orderId, userId)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
	} else {
		SendSuccessResponse(w)
		log.Println("(SUCCESS)\t", "Refund request")
	}
}

func SendReceipt() {

	m := gomail.NewMessage()

	// Get value from env
	emailSender := LoadEnv("EMAIL_SENDER")
	emailReceiver := LoadEnv("EMAIL_RECEIVER")
	emailPassword := LoadEnv("EMAIL_PASS")

	// Set email content
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", "Travelokay Order Receipt")

	text := "<h1>Your Purchase Receipt</h1></br>" +
		"<p>You have made a purchase via Traveloka app with the following details:</p>"
	m.SetBody("text/html", text)

	d := gomail.NewDialer("smtp.gmail.com", 465, emailSender, emailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}
