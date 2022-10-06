// Demo listener for data of queue Rabbit. If data is, it call method Directum RX via formed HyperLink.
// Демо "прослушиватель" данных очереди Rabbit. Если данные появились, вызывается метод Directum RX через сформированную гиперссылку.
package main

import (
	"log"

	"github.com/blablatov/listen2rabbit4dirx/call2handler"
	"github.com/pandeptwidyaop/gorabbit"
	"github.com/streadway/amqp"
)

func main() {
	// Create infinite chan for listen queue. Цикличный канал для прослушивания очереди.
	forever := make(chan bool)

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
	<-forever
}

// Handling messages of queue. Обработчик сообщений очереди.
func handleConsume(mq gorabbit.RabbitMQ, queue string, deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		switch d.RoutingKey {
		case "SAP_A":
			//log.Println("message come from SAP_A")
			log.Println("Сообщение пришло от SAP_A")
		case "SAP_B":
			//log.Println("message come from SAP_A")
			log.Println("Сообщение пришло от SAP_B")
		case "SAP_C":
			//log.Println("message come from SAP_A")
			log.Println("Сообщение пришло от SAP_C")
		}
		//log.Println("Call of method Directum RX via formed HyperLink")
		log.Println("Вызов обработчика сообщений RabbitMQ из Directum RX, через сформированную гиперссылку.")
		go call2handler.CallHadler()
	}
}
