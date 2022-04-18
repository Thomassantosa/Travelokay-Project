package controllers

import (
	"log"
	"time"

	"github.com/Travelokay-Project/models"
	"github.com/go-co-op/gocron"
	"gopkg.in/gomail.v2"
)

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
func OfferMail(hari int) {
	// Connect to database
	db := Connect()
	defer db.Close()

	m := gomail.NewMessage()

	// Get value from env
	emailSender := LoadEnv("EMAIL_SENDER")
	emailPassword := LoadEnv("EMAIL_PASS")
	rows, errQuery := db.Query("SELECT fullname, email FROM users")
	if errQuery != nil {
		log.Fatal(errQuery)
		return
	}
	var user models.User
	var emailReceiver []string
	i := 0
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.Fullname, &user.Email); err != nil {
			log.Fatal(errQuery)
			return
		}
		emailReceiver[i] = m.FormatAddress(user.Email, user.Fullname)
		i++
	}
	// Set email content
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailReceiver...)
	if hari == 0 {
		m.SetHeader("Subject", "test aja")
	}
	if hari == 1 {
		m.SetHeader("Subject", "Idul Fitri Promotion Offer")
	} else if hari == 2 {
		m.SetHeader("Subject", "Christmast Promotion Offer")
	} else if hari == 3 {
		m.SetHeader("Subject", "New Year Promotion Offer")
	}

	text := "<h1>Here Is Your Best Deal Offer</h1></br>" +
		"<p><a href='#'>click here</a> to see your deal</p>"
	m.SetBody("text/html", text)

	d := gomail.NewDialer("smtp.gmail.com", 465, emailSender, emailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}

func GocronEvent() {
	s := gocron.NewScheduler(time.UTC)
	s.Cron("*/5 * * * * *").Do(func() { log.Println("test") }) // every 5 sec
	s.Cron("* * * /2 /5 *").Do(OfferMail, 1)                   // every idul fitri
	s.Cron("* * * /25 /12 *").Do(OfferMail, 2)                 // every christmast
	s.Cron("* * * /1 /1 *").Do(OfferMail, 3)                   // every new year

	// starts the scheduler asynchronously
	s.StartAsync()
	// starts the scheduler and blocks current execution path
	// s.StartBlocking()
}
