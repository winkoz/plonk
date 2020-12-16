package io

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlainRenderer_RenderComponents(t *testing.T) {
	output := []byte("This string should go directly to log.StdOut\n")
	var stdout bytes.Buffer
	type args struct {
		output []byte
	}
	tests := []struct {
		name string
		pr   PlainRenderer
		args args
	}{
		{
			name: "RenderComponents passes the `output` to log as stdin",
			pr: PlainRenderer{
				log: log.New(&stdout, "", 0),
			},
			args: args{
				output: output,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pr.RenderComponents(tt.args.output)
			expected := string(output)
			got := stdout.String()
			assert.Equal(t, expected, got)
		})
	}
}
