package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AirtonLira/system_planning_financial.git/adapter/http/util/time"
	"github.com/AirtonLira/system_planning_financial.git/model/transaction"
)

// GetTransaction - Function of the return transaction in method GET
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    15000.0,
			Type:      0,
			CreatedAt: time.ConvertStringToTime("2023-11-19T10:52:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

// CreateTransaction - Function of the create a new transaction and print console application
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
