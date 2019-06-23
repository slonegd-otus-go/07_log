package log

import (
	"io"
)

type HwAccepted struct {
	Id    int
	Grade int
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

type OtusEvent interface {
}

func LogOtusEvent(e OtusEvent, w io.Writer) {

}
