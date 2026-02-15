package model

type Response struct {
	Payload any    `json:"payload"`
	Success bool   `json:"success"`
	TraceID string `json:"traceId"`
}

func EmptyResponse(ctx Context) *Response {
	return &Response{
		Success: true,
		TraceID: requestUUID(ctx),
	}
}

func BadResponse(ctx Context, problemDetail *ProblemDetail) *Response {
	return &Response{
		Payload: problemDetail,
		Success: false,
		TraceID: requestUUID(ctx),
	}
}

// FieldError Описание ошибки заполнения поля объекта
type FieldError struct {
	Item          string `json:"item"`
	Field         string `json:"field"`
	Message       string `json:"message"`
	RejectedValue any    `json:"rejectedValue"`
}

// ProblemDetail Описание возникшей ошибки
type ProblemDetail struct {
	Type     string        `json:"type"`
	Instance string        `json:"instance"`
	Title    string        `json:"title"`
	Detail   string        `json:"detail"`
	Errors   []*FieldError `json:"errors"`
}
