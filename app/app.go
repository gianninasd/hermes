package main

import (
	"dg/client"
	"log"
	"os"
	"runtime"
	"time"

	guuid "github.com/google/uuid"
	"github.com/magiconair/properties"
)

//var wg sync.WaitGroup

func processRequest(p *properties.Properties, cardRequest client.CardRequest, resp chan string) {
	//defer wg.Done()
	resp2 := client.SendPurchase(p, cardRequest)
	resp <- resp2 // send response to the channel
	//wg.Done()
}

// Main function executed to start the program
func main() {
	log.Println("Go", runtime.Version(), "File Processor running on", runtime.GOOS, "- Process ID", os.Getpid())
	log.Println("Number of CPUs:", runtime.NumCPU())
	//runtime.GOMAXPROCS(2)

	// load config from the property file
	p := properties.MustLoadFile("config.properties", properties.UTF8)

	bill := client.BillingDetails{Zip: "H8P1K4"}
	exp := client.CardExpiry{Month: "10", Year: "2022"}
	card := client.Card{CardNum: "4111111111111111", CardExpiry: exp}
	cardRequest := client.CardRequest{
		MerchantRefNum: "",
		Amount:         14000,
		SettleWithAuth: true,
		Card:           card,
		BillingDetails: bill,
	}

	// some dummy data, only amounts
	var records = [...]int{14000, 13, 5, 1000, 7723, 15000, 9956, 91, 14000, 13, 5, 1000, 7723, 15000, 9956, 91, 14000, 13, 5, 1000, 7723, 15000, 9956, 91}
	//var records = [...]int{14000, 13, 5, 1000, 7723}
	var totalRecords = len(records)

	resp := make(chan string, 6) // unbuffured string to receive response
	startTime := time.Now()

	for i := 0; i < totalRecords; i++ {
		//wg.Add(1)
		cardRequest.MerchantRefNum = guuid.New().String()
		cardRequest.Amount = records[i]

		log.Println("Sending reference", cardRequest.MerchantRefNum, "with amount", cardRequest.Amount)

		go processRequest(p, cardRequest, resp)

		//r := <-resp // get response from channel
		//log.Println(r)
	}

	//wg.Wait()

	// loop and print all responses
	for j := 0; j < totalRecords; j++ {
		log.Println(<-resp)
	}

	endTime := time.Now()
	log.Println("Done processing", totalRecords, "records in", endTime.Sub(startTime))
}
