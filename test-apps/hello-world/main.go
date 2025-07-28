package main

import (
    "log"
    "fmt"
    "net/http"
   "github.com/prometheus/client_golang/prometheus"
)

var httpRequestsCounter prometheus.Counter

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        httpRequestsCounter.Inc()
        log.Printf("Request reieved from %s\n", r.RemoteAddr);
        fmt.Fprintf(w, "Updated Hit Hello World!\n")
    })

    log.Fatal(http.ListenAndServe(":8080", nil));
}

func init() {

    httpRequestsCounter = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "http-requests-total",
        Help: "Total number of http requests handled.",
    });

    prometheus.MustRegister(httpRequestsCounter)
}
