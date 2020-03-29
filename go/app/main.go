package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Counter struct {
	Count int
	Time  time.Time
}

var counter Counter

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if counter.Count == 0 {
			counter = Counter{Count: 1, Time: time.Now()}
		} else {
			log.Printf("%v", counter.Count)
			counter.Count = counter.Count + 1
			counter.Time = time.Now()
		}
		log.Print("Accepted")
		message := r.URL.Path
		message = "This is " + strconv.Itoa(counter.Count) + "'s Count!\n"
		message += "Recent Request is " + counter.Time.String()
		t, _ := template.ParseFiles("template/index.html")
		t.Execute(w, message)
	})
	log.Print("starting web server")
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
