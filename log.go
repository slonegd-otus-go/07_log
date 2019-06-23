package otus

import (
	"fmt"
	"io"
	"time"
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
	String() string
}

func (h HwAccepted) String() string {
	return fmt.Sprintf("accepted %v %v", h.Id, h.Grade)
}

func (h HwSubmitted) String() string {
	return fmt.Sprintf("submitted %v \"%v\"", h.Id, h.Comment)
}

// for tests
var now = time.Now

func Log(event Event, writer io.Writer) {
	fmt.Fprintf(writer, "%v %v\n", now().Format("2006-01-02"), event)
}
