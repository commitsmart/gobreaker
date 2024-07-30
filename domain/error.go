package domain

type ErrorMessageBody struct {
	Code          int       `json:"-"`
	CorrelationID string    `json:"correlationId"`
	FunctionCode  string    `json:"functionCode"`
	Messages      []Message `json:"messages"`
}

type Message struct {
	Label        string `json:"label"`
	FormProperty string `json:"formProperty"`
}
