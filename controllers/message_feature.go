package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/Travelokay-Project/models"
	"github.com/go-co-op/gocron"
	"gopkg.in/gomail.v2"
)

func SendReceipt(emailReceiver string, newOrder models.Order, price int) {

	m := gomail.NewMessage()

	// Get value from env
	emailSender := LoadEnv("EMAIL_SENDER")
	emailPassword := LoadEnv("EMAIL_PASS")

	// Set email content
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", "Travelokay Order Receipt")

	text := `<h1>Your Purchase Receipt</h1></br>
		<p>You have made a purchase via Traveloka app with the following details:</p>
		<table>
		<tr>
			<td><b>Order ID</b></td>
			<td>: ` + strconv.Itoa(newOrder.ID) + `</td>
		</tr>
		<tr>
			<td><b>Order date</b></td>
			<td>: ` + newOrder.OrderDate + `</td>
		</tr>
		<tr>
			<td><b>Order status</b></td>
			<td>: ` + newOrder.OrderStatus + `</td>
		</tr>
		<tr>
			<td><b>Transaction type</b></td>
			<td>: ` + newOrder.TransactionType + `</td>
		</tr>
		<tr>
			<td><b>Price</b></td>
			<td>: ` + strconv.Itoa(price) + `</td>
		</tr>
		</table>`

	m.SetBody("text/html", text)

	d := gomail.NewDialer("smtp.gmail.com", 465, emailSender, emailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}
func OfferMail() {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from env
	emailSender := LoadEnv("EMAIL_SENDER")
	emailPassword := LoadEnv("EMAIL_PASS")
	rows, errQuery := db.Query("SELECT email FROM users")
	if errQuery != nil {
		log.Fatal(errQuery)
		return
	}
	var email string
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&email); err != nil {
			log.Fatal(errQuery)
			return
		}

		m := gomail.NewMessage()

		// Set email content
		m.SetHeader("From", emailSender)
		m.SetHeader("To", email)
		m.SetHeader("Subject", "Flash Sale Promo")

		text := "<h1>Here Is Your Best Deal Offer</h1></br>" +
			"<p><a href='https://github.com/nicholasrussel/PBP-API-Tools-1120003-1120011-1120027-1120037-1120043/blob/master/Photo%201.jpeg'>click here</a> to see your deal</p>"
		m.SetBody("text/html", text)

		d := gomail.NewDialer("smtp.gmail.com", 465, emailSender, emailPassword)

		// Send the email
		if err := d.DialAndSend(m); err != nil {
			log.Println(err)
		}
	}
}

func GocronEvent() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(OfferMail)
	s.Every(1).MonthLastDay().Do(OfferMail)

	// starts the scheduler asynchronously
	s.StartAsync()
}
