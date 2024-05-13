package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	mux       *http.ServeMux
	phonebook phonebook
}

func main() {
	app := &application{
		phonebook: phonebook{},
		mux:       http.NewServeMux(),
	}

	app.bindRoutes()

	go func() {
		http.ListenAndServe("localhost:7331", app.mux)
	}()

	for {
		fmt.Println("Enter your command:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		parseInput(scanner.Text(), &app.phonebook)
	}
}
