package dto

type CreateSubscriptionRequest struct {
	Name            string            `json:"name"`
	Amount          string            `json:"amount"`
	Currency        string            `json:"currency"`
	PaymentType     string            `json:"payment_type"`
	Token           string            `json:"token"`
	Schedule        ScheduleRequest   `json:"schedule"`
	RetrySchedule   *RetrySchedule    `json:"retry_schedule,omitempty"`
	Metadata        map[string]string `json:"metadata"`
	CustomerDetails CustomerDetails   `json:"customer_details"`
	Gopay           *Gopay            `json:"gopay,omitempty"`
}

type CreateSubscriptionResponse struct {
	Response
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	Amount          string            `json:"amount"`
	Currency        string            `json:"currency"`
	CreatedAt       string            `json:"created_at"`
	Schedule        ScheduleResponse  `json:"schedule"`
	RetrySchedule   *RetrySchedule    `json:"retry_schedule,omitempty"`
	Status          string            `json:"status"`
	Token           string            `json:"token"`
	PaymentType     string            `json:"payment_type"`
	Metadata        map[string]string `json:"metadata"`
	CustomerDetails CustomerDetails   `json:"customer_details"`
	Gopay           *Gopay            `json:"gopay,omitempty"`
}
