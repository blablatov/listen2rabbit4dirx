package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/blablatov/tlsgorabbit"
)

func Test(t *testing.T) {
	var tests = []struct {
		Name    string
		Address string
	}{
		{"admin", "qwerty"},
		{"Directum", "sap"},
		{"\t,..,", "NaN\null\n"},
		{"Data for test", "Number 99999 to data test"},
		{"Yes, no", "No, or, yes"},
	}

	var prevName string
	for _, test := range tests {
		if test.Name != prevName {
			fmt.Printf("\n%s\n", test.Name)
			prevName = test.Name
		}
	}

	var prevAddress string
	for _, test := range tests {
		if test.Address != prevAddress {
			fmt.Printf("\n%s\n", test.Address)
			prevAddress = test.Address
		}
	}
}

func TestConnPublish(t *testing.T) {
	var ctests = []struct {
		URL      string
		Queue    string
		Exchange string
		router   string
	}{
		{"amqp://guest:guest@localhost:5672/dirx", "QueueDirx", "ExchangeDirx", "router"},
	}

	var prevURL string
	for _, test := range ctests {
		if test.URL != prevURL {
			fmt.Printf("\n%s\n", test.URL)
			prevURL = test.URL
		}
	}

	var prevQueue string
	for _, test := range ctests {
		if test.Queue != prevQueue {
			fmt.Printf("\n%s\n", test.Queue)
			prevQueue = test.Queue
		}
	}

	var prevExchange string
	for _, test := range ctests {
		if test.Exchange != prevExchange {
			fmt.Printf("\n%s\n", test.Exchange)
			prevExchange = test.Exchange
		}
	}

	var prevrouter string
	for _, test := range ctests {
		if test.router != prevrouter {
			fmt.Printf("\n%s\n", test.router)
			prevrouter = test.router
		}
	}

	mq, err := gorabbit.New(
		prevURL,
		prevQueue,
		prevExchange,
	)
	if err != nil {
		log.Fatalf("Error of new: %v", err)
	}

	// Start connection.
	err = mq.Connect()
	err = nil
	if err != nil {
		log.Fatalf("Error of conn: %v", err)
	}
	log.Println("Connect yes: ", err)

	m := Message{
		Name:    prevQueue,
		Address: prevExchange,
	}
	jsonMessage, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error of marshal: %v", err)
	}
	fmt.Printf("Data of queue:%s", jsonMessage)

	// Bublish data to event of queue. Опубликовать данные в событии очереди.
	err = mq.Publish(prevrouter, "application/json", jsonMessage)
	if err != nil {
		log.Fatalf("Error of publish: %v", err)
	}
}

func TestConnPublishTLS(t *testing.T) {
	var ctests = []struct {
		URL      string
		Queue    string
		Exchange string
		router   string
	}{
		{"amqps://guest:guest@localhost:5671/dirx", "QueueDirx", "ExchangeDirx", "router"},
	}

	var prevURL string
	for _, test := range ctests {
		if test.URL != prevURL {
			fmt.Printf("\n%s\n", test.URL)
			prevURL = test.URL
		}
	}

	var prevQueue string
	for _, test := range ctests {
		if test.Queue != prevQueue {
			fmt.Printf("\n%s\n", test.Queue)
			prevQueue = test.Queue
		}
	}

	var prevExchange string
	for _, test := range ctests {
		if test.Exchange != prevExchange {
			fmt.Printf("\n%s\n", test.Exchange)
			prevExchange = test.Exchange
		}
	}

	var prevrouter string
	for _, test := range ctests {
		if test.router != prevrouter {
			fmt.Printf("\n%s\n", test.router)
			prevrouter = test.router
		}
	}

	mq, err := gorabbit.New(
		prevURL,
		prevQueue,
		prevExchange,
	)
	if err != nil {
		log.Fatalf("Error of new: %v", err)
	}

	// Start tls connection.
	err = mq.ConnectTLS()
	err = nil
	if err != nil {
		log.Fatalf("Error of conn: %v", err)
	}
	log.Println("ConnectTLS yes: ", err)

	m := Message{
		Name:    prevQueue,
		Address: prevExchange,
	}

	jsonMessage, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error of marshal: %v", err)
	}
	fmt.Printf("Data of queue:%s", jsonMessage)

	// Bublish data to event of queue. Опубликовать данные в событии очереди.
	/*err = mq.Publish(prevrouter, "application/json", jsonMessage)
	if err != nil {
		log.Fatalf("Error of publish: %v", err)
	}*/
}

func BenchmarkConnPublish(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 5; i++ {
		var ctests = []struct {
			URL      string
			Queue    string
			Exchange string
			router   string
		}{
			{"amqp://guest:guest@localhost:5672/dirx", "QueueDirx", "ExchangeDirx", "router"},
		}

		var prevURL string
		for _, test := range ctests {
			if test.URL != prevURL {
				fmt.Printf("\n%s\n", test.URL)
				prevURL = test.URL
			}
		}

		var prevQueue string
		for _, test := range ctests {
			if test.Queue != prevQueue {
				fmt.Printf("\n%s\n", test.Queue)
				prevQueue = test.Queue
			}
		}

		var prevExchange string
		for _, test := range ctests {
			if test.Exchange != prevExchange {
				fmt.Printf("\n%s\n", test.Exchange)
				prevExchange = test.Exchange
			}
		}

		var prevrouter string
		for _, test := range ctests {
			if test.router != prevrouter {
				fmt.Printf("\n%s\n", test.router)
				prevrouter = test.router
			}
		}

		mq, err := gorabbit.New(
			prevURL,
			prevQueue,
			prevExchange,
		)
		if err != nil {
			log.Fatalf("Error of new: %v", err)
		}

		// Start connection.
		err = mq.Connect()
		if err != nil {
			log.Fatalf("Error of conn: %v", err)
		}

		m := Message{
			Name:    prevQueue,
			Address: prevExchange,
		}
		jsonMessage, err := json.Marshal(m)
		if err != nil {
			log.Fatalf("Error of marshal: %v", err)
		}
		fmt.Printf("Data of queue:%s", jsonMessage)

		// Bublish data to event of queue. Опубликовать данные в событии очереди.
		err = mq.Publish(prevrouter, "application/json", jsonMessage)
		if err != nil {
			log.Fatalf("Error of publish: %v", err)
		}
	}
}

func BenchmarkConnPublishTLS(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 5; i++ {
		var ctests = []struct {
			URL      string
			Queue    string
			Exchange string
			router   string
		}{
			{"amqps://guest:guest@localhost:5671/dirx", "QueueDirx", "ExchangeDirx", "router"},
		}

		var prevURL string
		for _, test := range ctests {
			if test.URL != prevURL {
				fmt.Printf("\n%s\n", test.URL)
				prevURL = test.URL
			}
		}

		var prevQueue string
		for _, test := range ctests {
			if test.Queue != prevQueue {
				fmt.Printf("\n%s\n", test.Queue)
				prevQueue = test.Queue
			}
		}

		var prevExchange string
		for _, test := range ctests {
			if test.Exchange != prevExchange {
				fmt.Printf("\n%s\n", test.Exchange)
				prevExchange = test.Exchange
			}
		}

		var prevrouter string
		for _, test := range ctests {
			if test.router != prevrouter {
				fmt.Printf("\n%s\n", test.router)
				prevrouter = test.router
			}
		}

		mq, err := gorabbit.New(
			prevURL,
			prevQueue,
			prevExchange,
		)
		if err != nil {
			log.Fatalf("Error of new: %v", err)
		}

		// Start tls connection.
		err = mq.ConnectTLS()
		err = nil
		if err != nil {
			log.Fatalf("Error of conn: %v", err)
		}
		log.Println("ConnectTLS yes: ", err)

		m := Message{
			Name:    prevQueue,
			Address: prevExchange,
		}
		jsonMessage, err := json.Marshal(m)
		if err != nil {
			log.Fatalf("Error of marshal: %v", err)
		}
		fmt.Printf("Data of queue:%s", jsonMessage)

		// Bublish data to event of queue. Опубликовать данные в событии очереди.
		/*err = mq.Publish(prevrouter, "application/json", jsonMessage)
		if err != nil {
			log.Fatalf("Error of publish: %v", err)
		}*/
	}
}
