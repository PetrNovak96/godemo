package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("Ooops"))
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}
		log.Printf("Data: %s\n", d)
		fmt.Fprintf(rw, "Hello, Your data was: %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good Bye World")
	})

	http.ListenAndServe(":9090", nil)
}
