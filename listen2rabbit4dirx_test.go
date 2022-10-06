package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/blablatov/listen2rabbit4dirx/call2handler"
	"github.com/pandeptwidyaop/gorabbit"
)

func TestConn(t *testing.T) {
	var ctests = []struct {
		URL      string
		Queue    string
		Exchange string
	}{
		{"amqp://guest:guest@localhost:5672/dirx", "QueueDirx", "ExchangeDirx"},
		//{"http://guest:guest@localhost:5672/dirx", "qwerty", "qwerty"},
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
	deliveries := map[string]int{
		"SAP_A": 1,
		"SAP_B": 2,
		"SAP_C": 3,
	}
	for RoutingKey := range deliveries {
		switch RoutingKey {
		case "SAP_A":
			log.Println("Сообщение пришло от SAP_A")
		case "SAP_B":
			log.Println("Сообщение пришло от SAP_B")
		case "SAP_C":
			log.Println("Сообщение пришло от SAP_C")
		}
		//log.Println("Call of method Directum RX via formed HyperLink")
		log.Println("Вызов обработчика сообщений RabbitMQ из Directum RX, через сформированную гиперссылку.")
		go call2handler.CallHadler()
	}
}

func BenchmarkConnect(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 10; i++ {
		var ctests = []struct {
			URL      string
			Queue    string
			Exchange string
		}{
			{"amqp://guest:guest@localhost:5672/dirx", "QueueDirx", "ExchangeDirx"},
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
	}
}

func BenchmarkCallHandler(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 5; i++ {
		go call2handler.CallHadler()
	}
}
