package main

import (
    "fmt"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // App endpoint
    http.HandleFunc("/hello", helloHandler)
    
    // Metrics endpoint
    http.Handle("/metrics", promhttp.Handler())
    
    fmt.Println("Server running on :8080")
    fmt.Println("Try: http://localhost:8080/hello")
    fmt.Println("Metrics: http://localhost:8080/metrics")
    
    http.ListenAndServe(":8080", nil)
}