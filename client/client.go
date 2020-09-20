package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/magiconair/properties"
)

// BillingDetails customer address
type BillingDetails struct {
	Zip string `json:"zip"`
}

// CardExpiry card expiry date
type CardExpiry struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

// Card card details
type Card struct {
	CardNum    string     `json:"cardNum"`
	CardExpiry CardExpiry `json:"cardExpiry"`
}

// CardRequest the full request
type CardRequest struct {
	MerchantRefNum string         `json:"merchantRefNum"`
	Amount         int            `json:"amount"`
	SettleWithAuth bool           `json:"settleWithAuth"`
	Card           Card           `json:"card"`
	BillingDetails BillingDetails `json:"billingDetails"`
}

// SendPurchase call the remote REST API to perform a purchase request
func SendPurchase(p *properties.Properties, cardRequest CardRequest) string {
	body, err := json.Marshal(cardRequest)

	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", p.MustGetString("url"), bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+p.MustGetString("apikey"))

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

	var respBody = string(body2)
	return respBody
}
