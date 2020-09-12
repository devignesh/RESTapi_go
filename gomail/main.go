package main

import (
	"fmt"

	"github.com/mailgun/mailgun-go"
)

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun("gopherhut.com", "pubkey-ecc733d482a41f1a5aa11d0fb6698ebb")
	m := mg.NewMessage(
		"Mailgun Sandbox <postmaster@sandbox80d9ed4a4bae4afab1cc80ea5117d881.mailgun.org>",
		"Hello Vignesh kumar",
		"Testing mail with go",
		"vigneshkumar.mca2016@adhiyamaan.in",
	)
	_, id, err := mg.Send(m)
	return id, err
}

func main() {
	fmt.Println("Go with mail")
	SendSimpleMessage("gopherhut.com", "pubkey-ecc733d482a41f1a5aa11d0fb6698ebb")
}
