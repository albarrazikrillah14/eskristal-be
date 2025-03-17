package exceptions

type logBody struct {
	TraceID any `json:"trace_id"`
	Body    any `json:"body"`
}

func NewLogBody(traceID any, body any) logBody {
	return logBody{
		TraceID: traceID,
		Body:    body,
	}
}
