package log_v1

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

type ErrOffsetOutOfRange struct {
	Offset uint64
}

func (e ErrOffsetOutOfRange) GRPCStatus() *status.Status {
	s := status.New(404, fmt.Sprintf("offset out of range: %d", e.Offset))
	msg := fmt.Sprintf("the requested offset is outside the log's range: %d", e.Offset)
	d := &errdetails.LocalizedMessage{
		Locale:  "en-US",
		Message: msg,
	}
	std, err := s.WithDetails(d)
	if err != nil {
		return s
	}
	return std
}

func (e ErrOffsetOutOfRange) Error() string {
	return e.GRPCStatus().Err().Error()
}
