package otus_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	otus "github.com/slonegd-otus-go/07_log"
)

func TestLogOtusEvent(t *testing.T) {
	defer otus.SetRealNow()

	tests := []struct {
		name  string
		now   string
		event otus.Event
		want  string
	}{
		{
			name: "from readme 1",
			now:  "2019-01-01",
			event: otus.HwSubmitted{
				Id:      3456,
				Comment: "please take a look at my homework",
			},
			want: "2019-01-01 submitted 3456 \"please take a look at my homework\"\n",
		},
		{
			name: "from readme 2",
			now:  "2019-01-01",
			event: otus.HwAccepted{
				Id:    3456,
				Grade: 4,
			},
			want: "2019-01-01 accepted 3456 4\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			otus.SetNow(tt.now)
			writer := &bytes.Buffer{}
			otus.Log(tt.event, writer)
			assert.Equal(t, tt.want, writer.String())
		})
	}
}
