package dto

type ScheduleRequest struct {
	Interval     int    `json:"interval"`
	IntervalUnit string `json:"interval_unit"`
	MaxInterval  int    `json:"max-interval"`
	StartTime    string `json:"start_time"`
}

type ScheduleResponse struct {
	Interval            int    `json:"interval"`
	IntervalUnit        string `json:"interval_unit"`
	StartTime           string `json:"start_time"`
	PreviousExecutionAt string `json:"previous_execution_at"`
	NextExecutionAt     string `json:"next_execution_at"`
}

type Schedule struct {
	Interval int `json:"interval"`
}

type RetrySchedule struct {
	Interval     int    `json:"interval"`
	IntervalUnit string `json:"interval_unit"`
	MaxInterval  int    `json:"max_interval"`
}

type CustomerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Gopay struct {
	AccountId string `json:"account_id"`
}

type Response struct {
	StatusMessage      *string   `json:"status_message"`
	ValidationMessages *[]string `json:"validation_messages,omitempty"`
}
