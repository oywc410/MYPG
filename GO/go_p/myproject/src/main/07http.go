package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")

	//http://127.0.0.1:3001/?id=1
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err == nil {
		fmt.Fprint(w, id)
	} else {
		log.Fatal(err)
	}

}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "POST" {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		//http 201
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/person", PersonHandler)
	http.ListenAndServe(":3001", nil)
}
