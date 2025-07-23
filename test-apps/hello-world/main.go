package main

import (
    "log"
    "fmt"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        log.Printf("Request reieved from %s\n", r.RemoteAddr);
        fmt.Fprintf(w, "Updated Hit Hello World!\n")
    })

    log.Fatal(http.ListenAndServe(":8080", nil));
}
