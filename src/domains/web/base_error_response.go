package web

type baseErrorResponse struct {
	CallID any `json:"call_id"`
	Errors any `json:"errors"`
}

func NewBaseErrorResponse(callID any, errors any) baseErrorResponse {
	return baseErrorResponse{
		CallID: callID,
		Errors: errors,
	}
}
