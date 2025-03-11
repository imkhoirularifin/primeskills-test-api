package dto

type UpdateSubscriptionRequest struct {
	Response
	Name          string        `json:"name"`
	Amount        string        `json:"amount"`
	Currency      string        `json:"currency"`
	Token         string        `json:"token"`
	Schedule      Schedule      `json:"schedule"`
	RetrySchedule RetrySchedule `json:"retry_schedule"`
	Gopay         *Gopay        `json:"gopay,omitempty"`
}
