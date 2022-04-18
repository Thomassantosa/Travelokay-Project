package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	models "github.com/Travelokay-Project/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value form cookie
	userId := GetIdFromCookie(r)

	// Get value from form
	fullname := r.FormValue("fullname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	address := r.FormValue("address")

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

	queryNew := query[:len(query)-1] // Delete last coma
	queryNew += " WHERE user_id = " + strconv.Itoa(userId)

	_, errQuery := db.Exec(queryNew)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
	} else {
		SendSuccessResponse(w)
		log.Println("(SUCCESS)\t", "Update partner request")
	}
}

func AddNewOrder(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	userId := GetIdFromCookie(r)
	seatId := r.FormValue("seatId")
	roomId := r.FormValue("roomId")
	tourScheduleId := r.FormValue("tourScheduleId")
	transactionType := r.FormValue("transactionType")

	// Order flight, traintrip, or bustrip
	if seatId != "" {

		// Check seat_id
		row := db.QueryRow("SELECT seat_status FROM seats WHERE seat_id = ?", seatId)
		var seatStatus int
		if err := row.Scan(&seatStatus); err != nil {
			log.Print("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}

		if seatStatus == 0 {

			// Query order & update seat_status
			_, errQuery1 := db.Exec("INSERT INTO orders(user_id, seat_id, transaction_type) values (?,?,?)", userId, seatId, transactionType)
			_, errQuery2 := db.Exec("UPDATE seats SET seat_status = 1 WHERE seat_id = ?", seatId)

			if errQuery1 != nil {
				log.Println("(ERROR)\t", errQuery1)
				SendErrorResponse(w, 500)
				return
			}
			if errQuery2 != nil {
				log.Println("(ERROR)\t", errQuery2)
				SendErrorResponse(w, 500)
				return
			}
			SendSuccessResponse(w)
			log.Println("(SUCCESS)\t", "Add new order success")
			return

		} else {
			SendMessageOnlyResponse(w, "Seat already booked")
			return
		}
	}

	// Order room
	if roomId != "" {

		// Check room_id
		row := db.QueryRow("SELECT room_status FROM rooms WHERE room_id = ?", roomId)
		var roomStatus int
		if err := row.Scan(&roomStatus); err != nil {
			log.Print("(ERROR)\t", err)
			SendErrorResponse(w, 400)
			return
		}

		if roomStatus == 0 {

			// Query order & update seat_status
			_, errQuery1 := db.Exec("INSERT INTO orders(user_id, room_id, transaction_type) values (?,?,?)", userId, roomId, transactionType)
			_, errQuery2 := db.Exec("UPDATE rooms SET room_status = 1 WHERE room_id = ?", roomId)

			if errQuery1 != nil {
				log.Println("(ERROR)\t", errQuery1)
				SendErrorResponse(w, 500)
				return
			}
			if errQuery2 != nil {
				log.Println("(ERROR)\t", errQuery2)
				SendErrorResponse(w, 500)
				return
			}
			SendSuccessResponse(w)
			log.Println("(SUCCESS)\t", "Add new order success")
			return
		} else {
			SendMessageOnlyResponse(w, "Room already booked")
			return
		}
	}

	// Order tour
	if tourScheduleId != "" {

		// Query order
		_, errQuery := db.Exec("INSERT INTO orders(user_id, tourschedule_id, transaction_type) values (?,?,?)", userId, tourScheduleId, transactionType)

		if errQuery != nil {
			log.Println("(ERROR)\t", errQuery)
			SendErrorResponse(w, 400)
			return
		} else {
			SendSuccessResponse(w)
			log.Println("(SUCCESS)\t", "Add new order success")
			return
		}
	}

	if seatId == "" && roomId == "" && tourScheduleId == "" {
		SendErrorResponse(w, 400)
		return
	}
}

// func AddNewBusOrder(w http.ResponseWriter, r *http.Request) {
// 	// Connect to database
// 	db := Connect()
// 	defer db.Close()

// 	// Get value from form
// 	err := r.ParseForm()
// 	if err != nil {
// 		return
// 	}
// }

// func AddNewTrainOrder(w http.ResponseWriter, r *http.Request) {
// 	// Connect to database
// 	db := Connect()
// 	defer db.Close()

// 	// Get value from form
// 	err := r.ParseForm()
// 	if err != nil {
// 		SendErrorResponse(w, 500)
// 		log.Println(err)
// 		return
// 	}
// 	userId := GetIdFromCookie(r)
// 	seatId := r.Form.Get("seatId")
// 	transactionType := r.Form.Get("transactionType")
// 	// Query order & update seat_status
// 	_, errQuery1 := db.Exec("INSERT INTO orders(user_id, seat_id, transaction_type) values (?,?,?)", userId, seatId, transactionType)
// 	_, errQuery2 := db.Exec("UPDATE seats SET seat_status = 1 WHERE seat_id = ", seatId)

// 	if errQuery1 == nil && errQuery2 == nil {
// 		SendSuccessResponse(w)
// 	} else if errQuery1 != nil {
// 		log.Println(errQuery1)
// 		SendErrorResponse(w, 400)
// 		return
// 	} else {
// 		log.Println(errQuery2)
// 		SendErrorResponse(w, 400)
// 		return
// 	}
// }

// func AddNewFlightOrder(w http.ResponseWriter, r *http.Request) {

// 	// Connect to database
// 	db := Connect()
// 	defer db.Close()

// 	// Get value from form
// 	err := r.ParseForm()
// 	if err != nil {
// 		SendErrorResponse(w, 500)
// 		log.Println(err)
// 		return
// 	}
// 	userId := GetIdFromCookie(r)
// 	seatId := r.Form.Get("seatId")
// 	transactionType := r.Form.Get("transactionType")

// 	// Check seat_id
// 	row := db.QueryRow("SELECT seat_status FROM seats WHERE seat_id=?", seatId)
// 	var seatType int
// 	if err := row.Scan(&seatType); err != nil {
// 		SendErrorResponse(w, 400)
// 		log.Print(err)
// 		return
// 	}

// 	if seatType == 0 {

// 		// Query order & update seat_status
// 		_, errQuery1 := db.Exec("INSERT INTO orders(user_id, seat_id, transaction_type) values (?,?,?)", userId, seatId, transactionType)
// 		_, errQuery2 := db.Exec("UPDATE seats SET seat_status = 1 WHERE seat_id = ?", seatId)

// 		if errQuery1 == nil && errQuery2 == nil {
// 			SendSuccessResponse(w)
// 		} else if errQuery1 != nil {
// 			log.Println(errQuery1)
// 			SendErrorResponse(w, 400)
// 			return
// 		} else {
// 			log.Println(errQuery2)
// 			SendErrorResponse(w, 400)
// 			return
// 		}
// 	} else {
// 		SendMessageOnlyResponse(w, "Seat already booked")
// 		return
// 	}
// }

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
		if err := rows.Scan(&order.Order_id, &order.UserID, &order.SeatID, &order.RoomID, &order.TourScheduleID, &order.OrderDate, &order.OrderStatus, &order.TransactionType); err != nil {
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
