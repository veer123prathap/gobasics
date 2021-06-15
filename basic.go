package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://159.65.148.87:9090/metrics")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Fatal(err)
	}

	io.WriteString(w, string(body))
}

func main() {
	http.HandleFunc("/vee", index)
	http.ListenAndServe(":8080", nil)
}
