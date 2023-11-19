package transaction

import "time"

// Transaction - Struct about transaction informations and what will return in the new request
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int16     `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions - Collection of the transaction
type Transactions []Transaction
