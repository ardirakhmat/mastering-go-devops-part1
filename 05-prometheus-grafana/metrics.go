package main

import "github.com/prometheus/client_golang/prometheus"

var requestsProcessed = prometheus.NewCounter(
    prometheus.CounterOpts{
        Name: "myapp_requests_processed_total",
        Help: "The total number of processed requests",
    })

func init() {
    prometheus.MustRegister(requestsProcessed)
}