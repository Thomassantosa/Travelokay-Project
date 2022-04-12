package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	models "github.com/Travelokay-Project/models"
)

func GetRefundList(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	orderStatus := r.Form.Get("order_status")

	rows, errQuery := db.Query("SELECT order_id,user_id,order_date,order_status,transaction_type FROM orders WHERE order_status=?", orderStatus)

	var order models.Orders
	var orders []models.Orders

	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.OrderStatus, &order.TransactionType); err != nil {
			log.Println(err.Error())
		} else {
			orders = append(orders, order)
		}
	}

	var response models.OrdersResponse
	if errQuery == nil {
		if len(orders) == 0 {
			SendErrorResponse(w, 400)
		} else {
			response.Status = 200
			response.Message = "Success Get Data"
			response.Data = orders
		}
	} else {
		SendErrorResponse(w, 400)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	db.Close()
}
