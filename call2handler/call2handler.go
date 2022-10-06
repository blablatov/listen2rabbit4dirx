// Demo method call Directum RX via formed HyperLink.
// Вызов метода Directum RX через сформированный HyperLink.
package call2handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	//url = "http://directum.integration/Sungero?module.id=1a2eb24e-6ad5-45e1-9e7f-d75aad92a869&module.func=rabbit_handler"
	url = "https://club.directum.ru/robots.txt" // demo link
)

func CallHadler() {
	// Call of method Directum RX via formed HyperLink.
	// Вызов обработчика сообщений RabbitMQ из Directum RX, через сформированную HyperLink.
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	// Handling response of server (optional). Обработка ответа сервера.
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: чтение %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}
