package main

import (
	"fmt"
	"os"

	"github.com/torukita/go-fitbit/fitbit"
)

func main() {
	expiry := "2592000" // one month expiry
	access_token := os.Getenv("TEST_TOKEN")
	id := os.Getenv("TEST_CLIENT_ID")
	secret := os.Getenv("TEST_CLIENT_SECRET")
	redirect := os.Getenv("TEST_CALLBACK_URL")
	if access_token == "" || id == "" || secret == "" || redirect == "" {
		os.Exit(1)
	}
	cfg := &fitbit.AuthConfig{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  redirect,
	}
	url := fitbit.ImplicitURL(cfg, expiry)
	fmt.Println("==== you can get access_token by accessing the following url (implicit grant)====")
	fmt.Println(url)
}
