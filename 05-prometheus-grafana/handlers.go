package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Increment counter
    requestsProcessed.Inc()
    
    fmt.Fprintf(w, "Hello! This request has been counted.\n")
}