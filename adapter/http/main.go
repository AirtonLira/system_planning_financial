package http

import (
	"fmt"
	"net/http"

	"github.com/AirtonLira/system_planning_financial.git/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/", transaction.GetTransaction)
	http.HandleFunc("/create", transaction.CreateTransaction)

	fmt.Println("Listening 8090.......")
	http.ListenAndServe(":8090", nil)
}
