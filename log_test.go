package log_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	log "github.com/slonegd-otus-go/07_log"
)

func TestLogOtusEvent(t *testing.T) {
	tests := []struct {
		name  string
		event log.OtusEvent
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			log.LogOtusEvent(tt.event, writer)
			assert.Equal(t, tt.want, writer.String())
		})
	}
}
