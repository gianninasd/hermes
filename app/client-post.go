package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	guuid "github.com/google/uuid"
)

func main() {
	log.Println("HTTP test client |", "PID=", os.Getpid())

	type BillingDetails struct {
		Zip string `json:"zip"`
	}

	type CardExpiry struct {
		Month string `json:"month"`
		Year  string `json:"year"`
	}

	type Card struct {
		CardNum    string     `json:"cardNum"`
		CardExpiry CardExpiry `json:"cardExpiry"`
	}

	type CardRequest struct {
		MerchantRefNum string         `json:"merchantRefNum"`
		Amount         int            `json:"amount"`
		SettleWithAuth bool           `json:"settleWithAuth"`
		Card           Card           `json:"card"`
		BillingDetails BillingDetails `json:"billingDetails"`
	}

	bill := BillingDetails{Zip: "H8P1K4"}
	exp := CardExpiry{Month: "10", Year: "2022"}
	card := Card{CardNum: "4111111111111111", CardExpiry: exp}
	cardRequest := CardRequest{
		MerchantRefNum: guuid.New().String(),
		Amount:         14000,
		SettleWithAuth: true,
		Card:           card,
		BillingDetails: bill,
	}

	log.Println("Sending ...", cardRequest)
	body, err := json.Marshal(cardRequest)

	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", "https://api.test.paysafe.com/cardpayments/v1/accounts/1001289630/auths", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic dGVzdF9hc3NsMTpCLXFhMi0wLTViZTg4MzJjLTAtMzAyYzAyMTQ2Y2Q4ZDUyZGRjY2E4ZWU4Y2U1Nzg0NTUwNWNlODBjZmNhYjIzYzYyMDIxNDBmYjAzMDBiMGJmOWE4Y2M2M2ZjMGI3ZDU4ZTJjMGMxYjY3MjQxMzA=")

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Completed! ...")
	log.Println(string(body2))
}
