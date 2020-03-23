package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Counter struct {
	Id    int
	Count int
	Time  time.Time
}

var CounterById map[int]*Counter

func main() {
	var counter Counter
	CounterById = make(map[int]*Counter)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if CounterById[1] == nil {
			counter = Counter{Id: 1, Count: 1, Time: time.Now()}
			CounterById[counter.Id] = &counter
		} else {
			log.Printf("%v", CounterById[1].Count)
			CounterById[1].Count = CounterById[1].Count + 1
			CounterById[1].Time = time.Now()
		}
		log.Print("Accepted")
		message := r.URL.Path
		message = "This is " + strconv.Itoa(CounterById[1].Count) + "'s Count!\n"
		message += "Recent Request is " + CounterById[1].Time.String()
		t, _ := template.ParseFiles("template/index.html")
		t.Execute(w, message)
	})
	log.Print("starting web server")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
