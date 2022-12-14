// Demo listener for data of queue Rabbit. If data is, it call method Directum RX via formed HyperLink.
// Демо "прослушиватель" данных очереди Rabbit. Если данные появились, вызывается метод Directum RX через сформированную гиперссылку.

package main

import (
	"log"
	"sync"

	"github.com/blablatov/listen2rabbit4dirx/call2handler"
	"github.com/blablatov/tlsgorabbit"
	"github.com/streadway/amqp"
)

func main() {
	// Create infinite chan for listen queue. Цикличный канал для прослушивания очереди.
	forever := make(chan bool)
	empty := make(chan int)
	close(empty)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		mq, err := gorabbit.New(
			"amqp://guest:guest@localhost:5672/dirx",
			"QueueDirx",
			"ExchangeDirx",
		)
		if err != nil {
			log.Fatalf("Error of new: %v", err)
		}

		// Start connection.
		err = mq.Connect()
		if err != nil {
			log.Fatalf("Error of conn: %v", err)
		}

		// Binding all routing key. Привязка всех ключей маршрутизации.
		err = mq.Bind([]string{"SAP_A", "SAP_B", "SAP_C"})
		if err != nil {
			log.Fatalf("Error of bind: %v", err)
		}

		deliveries, err := mq.Consume()
		if err != nil {
			log.Fatalf("Error of consume: %v", err)
		}

		//log.Println("Waiting for messages")
		log.Println("Ожидание сообщений")

		for q, d := range deliveries {
			go mq.HandleConsumedDeliveries(q, d, handleConsume)
		}
	}()
	<-forever
	go func() {
		defer wg.Done()
		mqt, err := gorabbit.New(
			"amqps://guest:guest@localhost:5671/dirx",
			"QueueDirx",
			"ExchangeDirx",
		)
		if err != nil {
			log.Fatalf("Error of new: %v", err)
		}

		// Start TLS connection.
		err = mqt.ConnectTLS()
		if err != nil {
			log.Fatalf("Error of conn: %v", err)
		}

		// Binding all routing key. Привязка всех ключей маршрутизации.
		err = mqt.Bind([]string{"SAP_A", "SAP_B", "SAP_C"})
		if err != nil {
			log.Fatalf("Error of bind: %v", err)
		}

		deliveries, err := mqt.Consume()
		if err != nil {
			log.Fatalf("Error of consume: %v", err)
		}

		//log.Println("Waiting for messages TLS")
		log.Println("Ожидание сообщений TLS")

		for q, d := range deliveries {
			go mqt.HandleConsumedDeliveries(q, d, handleConsume)
		}
	}()
	<-forever
	// Waits of counter. Ожидание счетчика.
	go func() {
		wg.Wait()
		close(forever)
	}()
}

// Handling messages of queue. Обработчик сообщений очереди.
func handleConsume(mq gorabbit.RabbitMQ, queue string, deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		switch d.RoutingKey {
		case "SAP_A":
			//log.Println("Message come from SAP_A")
			log.Println("Сообщение пришло от SAP_A")
		case "SAP_B":
			//log.Println("Message come from SAP_A")
			log.Println("Сообщение пришло от SAP_B")
		case "SAP_C":
			//log.Println("Message come from SAP_A")
			log.Println("Сообщение пришло от SAP_C")
		}
		//log.Println("Call of method Directum RX via formed HyperLink")
		log.Println("Вызов обработчика сообщений RabbitMQ из Directum RX, \nчерез сформированную гиперссылку: `https://club.directum.ru/robots.txt`")
		go call2handler.CallHadler()
	}
}
