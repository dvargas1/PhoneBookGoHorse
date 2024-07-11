package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type application struct {
	counter500 prometheus.Counter
	mux        *http.ServeMux
	phonebook  phonebook
}

func main() {
	mychan := make(chan int)
	app := &application{
		phonebook: phonebook{},
		mux:       http.NewServeMux(),
		counter500: promauto.NewCounter(prometheus.CounterOpts{
			Name: "myapp_processed_500_total",
			Help: "The total number of 500 responses",
		}),
	}

	app.bindRoutes()

	go func() {
		fmt.Println("starting server on localhost:7331!")
		app.mux.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe("localhost:7331", app.mux)
		mychan <- 0
	}()

	<-mychan

	// for {
	// 	fmt.Println("Enter your command:")
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	err := scanner.Err()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	parseInput(scanner.Text(), &app.phonebook)
	// }
}
