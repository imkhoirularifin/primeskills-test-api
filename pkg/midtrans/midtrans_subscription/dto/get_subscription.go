package dto

type GetSubscriptionResponse struct {
	Response
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	Amount          string            `json:"amount"`
	Currency        string            `json:"currency"`
	CreatedAt       string            `json:"created_at"`
	Schedule        ScheduleResponse  `json:"schedule"`
	Status          string            `json:"status"`
	Token           string            `json:"token"`
	PaymentType     string            `json:"payment_type"`
	TransactionIds  *[]string         `json:"transaction_ids,omitempty"`
	Metadata        map[string]string `json:"metadata"`
	CustomerDetails CustomerDetails   `json:"customer_details"`
	Gopay           *Gopay            `json:"gopay,omitempty"`
}
