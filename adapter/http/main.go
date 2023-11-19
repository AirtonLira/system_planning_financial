package http

import (
	"fmt"
	"net/http"

	"github.com/AirtonLira/system_planning_financial.git/adapter/http/infra"
	"github.com/AirtonLira/system_planning_financial.git/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/", transaction.GetTransaction)
	http.HandleFunc("/create", transaction.CreateTransaction)
	http.HandleFunc("/health", infra.Health)

	fmt.Println("Listening 8090.......")
	http.ListenAndServe(":8090", nil)
}
