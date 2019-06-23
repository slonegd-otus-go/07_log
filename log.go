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
	return fmt.Sprintf("%v accepted %v %v", now().Format("2006-01-02"), h.Id, h.Grade)
}

func (h HwSubmitted) String() string {
	return fmt.Sprintf("%v submitted %v \"%v\"", now().Format("2006-01-02"), h.Id, h.Comment)
}

// for tests
var now = time.Now

func Log(event Event, writer io.Writer) {
	fmt.Fprintf(writer, "%v\n", event)
}
