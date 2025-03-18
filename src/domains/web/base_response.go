package web

type baseResponse struct {
	CallID string `json:"call_id"`
	Data   any    `json:"data"`
}

func NewBaseResponse(callID string, data any) baseResponse {
	return baseResponse{
		CallID: callID,
		Data:   data,
	}
}
