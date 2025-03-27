package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func sendNotfiyEmail(taskName string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "task-schedule@gmail.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Task Time!")
	message := fmt.Sprintf("hi bro =) time to complate %s task", taskName)
	m.SetBody("text/plain", message)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("MAIL_HOST")
	password := os.Getenv("MAIL_PASSWORD")
	username := os.Getenv("MAIL_USERNAME")

	d := gomail.NewDialer(host, 25, username, password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
}
