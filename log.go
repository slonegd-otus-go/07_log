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
	GetID() int
	GetComment() string
	GetGrade() int
}

func (h HwAccepted) GetID() int {
	return h.Id
}

func (h HwAccepted) GetComment() string {
	return ""
}

func (h HwAccepted) GetGrade() int {
	return h.Grade
}

func (h HwSubmitted) GetID() int {
	return h.Id
}

func (h HwSubmitted) GetComment() string {
	return h.Comment
}

func (h HwSubmitted) GetGrade() int {
	return 0
}

// for tests
var now = time.Now

func Log(event Event, writer io.Writer) {
	switch event.(type) {
	case HwSubmitted:
		fmt.Fprintln(writer, now().Format("2006-01-02"), "submitted", event.GetID(), `"`+event.GetComment()+`"`)
	case HwAccepted:
		fmt.Fprintln(writer, now().Format("2006-01-02"), "accepted", event.GetID(), event.GetGrade())
	}
}
