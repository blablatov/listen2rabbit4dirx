// Demo publisher for data send to queue Rabbit.
// Демо публикатор для отправки данных в очередь Rabbit.
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pandeptwidyaop/gorabbit"
)

type Message struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	for _, router := range os.Args[1:] {
		m := Message{
			Name:    "directum",
			Address: "integration",
		}

		jsonMessage, err := json.Marshal(m)
		if err != nil {
			log.Fatalf("Error of marshal: %v", err)
		}

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

		// Bublish data to event of queue. Опубликовать данные в событии очереди.
		err = mq.Publish(router, "application/json", jsonMessage)
		if err != nil {
			log.Fatalf("Error of publish: %v", err)
		}
	}
}
