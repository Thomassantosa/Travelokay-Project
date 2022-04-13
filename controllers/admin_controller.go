package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	models "github.com/Travelokay-Project/models"
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

	var order models.Orders
	var orders []models.Orders

	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.UserID, &order.SeatID, &order.RoomID, &order.TourScheduleID, &order.OrderDate, &order.OrderStatus, &order.TransactionType); err != nil {
			log.Println(err.Error())
		} else {
			orders = append(orders, order)
		}
	}

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
