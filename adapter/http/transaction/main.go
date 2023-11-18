package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AirtonLira/system_planning_financial.git/model/transaction"
)

var layout = "2006-01-02T15:04:05"

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	salaryReceived, _ := time.Parse(layout, "2023-11-18T16:41:00")
	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    15000.0,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

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
