package accountTransactionHandler

type EventRequest struct {
	EventType     string  `json:"type"`
	FromAccountID string  `json:"origin"`
	ToAccountID   string  `json:"destination"`
	Amount        float64 `json:"amount"`
}
