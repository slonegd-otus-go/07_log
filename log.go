package otus

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

type Event interface {
}

func Log(event Event, writer io.Writer) {

}
