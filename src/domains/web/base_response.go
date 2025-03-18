package web

type baseResponse struct {
	CallID any `json:"call_id,omitempty"`
	Data   any `json:"data"`
}

func NewBaseResponse(callID any, data any) baseResponse {
	return baseResponse{
		CallID: callID,
		Data:   data,
	}
}
