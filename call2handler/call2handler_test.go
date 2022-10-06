package call2handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		URL string
	}{
		{"https://club.directum.ru/robots.txt"},
	}

	var prevURL string
	for _, test := range tests {
		if test.URL != prevURL {
			fmt.Printf("\n%s\n", test.URL)
			prevURL = test.URL
		}
	}
	resp, err := http.Get(prevURL)
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

func BenchmarkCallHandler(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 5; i++ {
		var tests = []struct {
			URL string
		}{
			{"https://club.directum.ru/robots.txt"},
		}

		var prevURL string
		for _, test := range tests {
			if test.URL != prevURL {
				fmt.Printf("\n%s\n", test.URL)
				prevURL = test.URL
			}
		}
		resp, err := http.Get(prevURL)
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
}
