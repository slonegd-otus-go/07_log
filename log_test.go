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
		{
			name: "from readme 1",
			event: log.HwSubmitted{
				Id:      3456,
				Comment: "please take a look at my homework",
			},
			want: `2019-01-01 submitted 3456 "please take a look at my homework"`,
		},
		{
			name: "from readme 2",
			event: log.HwAccepted{
				Id:    3456,
				Grade: 4,
			},
			want: `2019-01-01 accepted 3456 4`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			log.LogOtusEvent(tt.event, writer)
			assert.Equal(t, tt.want, writer.String())
		})
	}
}
