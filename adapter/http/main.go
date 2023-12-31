package http

import (
	"fmt"
	"net/http"

	"github.com/AirtonLira/system_planning_financial/adapter/http/infra"
	"github.com/AirtonLira/system_planning_financial/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init - function initial application
func Init() {
	http.HandleFunc("/", transaction.GetTransaction)
	http.HandleFunc("/create", transaction.CreateTransaction)
	http.HandleFunc("/health", infra.Health)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Listening 8090.......")
	http.ListenAndServe(":8090", nil)
}
