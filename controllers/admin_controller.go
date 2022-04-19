package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	models "github.com/Travelokay-Project/models"
	"github.com/gorilla/mux"
)

func GetRefundList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Query
	rows, errQuery := db.Query("SELECT * FROM orders WHERE order_status = 'refund'")

	if errQuery != nil {
		log.Println(errQuery)
		return
	}

	var order models.Order
	var orders []models.Order

	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.UserID, &order.SeatID, &order.RoomID, &order.TourScheduleID, &order.OrderDate, &order.OrderStatus, &order.TransactionType); err != nil {
			log.Println(err.Error())
		} else {
			orders = append(orders, order)
		}
	}

	// Response
	var response models.OrdersResponse
	if len(orders) == 0 {
		SendErrorResponse(w, 204)
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = orders
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func ApproveRefund(w http.ResponseWriter, r *http.Request) {

	//Connect database via Gorm
	db := ConnectGorm()

	//Connect database via MySql
	dbSql := Connect()
	defer dbSql.Close()

	//Get value from mux.Vars
	vars := mux.Vars(r)
	orderId := vars["orderId"]

	//Query
	row := dbSql.QueryRow("SELECT order_status FROM orders WHERE order_id = ?", orderId)
	var orderStatus string
	if err := row.Scan(&orderStatus); err != nil {
		log.Print("(ERROR)\t", err)
		SendErrorResponse(w, 500)
		return
	}

	if orderStatus != "refund" {
		SendErrorResponse(w, 400)
		log.Println("a")
		return
	}

	//Query using gorm
	db.Delete(&models.Orders{}, orderId)

	// Response
	SendSuccessResponse(w)
}
