package domain

type ResultCode string
type TaskType string

type Task interface {
	Execute(data []byte) (ResultCode, error)
}

const (
	ResultCodeSuccess       ResultCode = "success"
	ResultCodeFailure       ResultCode = "failure"
	ResultCodeInternalError ResultCode = "internal_error"
)
