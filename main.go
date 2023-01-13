package main

import (
	"crypto/rand"
	"fmt"
	"net/smtp"

	// "strings"

	"math/big"
	"strconv"
)

func main() {
	otp, err := getRandNum()
	if err != nil {
		fmt.Println(err)
		return
	}
	sendMail(otp)
}

func sendMail(otp string) {

	// Sender data.
	from := "<sender-email>"
	password := "<password>"

	// Receiver email address.
	to := []string{
		"<reciever-email>",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(otp+" is your otp"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func getRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}
