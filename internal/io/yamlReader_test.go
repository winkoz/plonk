package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_yamlReader_Parse(t *testing.T) {
	service := NewService()
	output := []map[string]string{}
	type fields struct {
		service Service
	}
	type args struct {
		data   []byte
		output []map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []map[string]string
		wantErr bool
	}{
		{
			name: "parses array of maps",
			fields: fields{
				service: service,
			},
			args: args{
				data:   []byte("[{key1: test, key2: test, key3: test}, {key1: test, key2: test, key3: test}, {key1: test, key2: test, key3: test}]"),
				output: output,
			},
			want: []map[string]string{
				{
					"key1": "test",
					"key2": "test",
					"key3": "test",
				},
				{
					"key1": "test",
					"key2": "test",
					"key3": "test",
				},
				{
					"key1": "test",
					"key2": "test",
					"key3": "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yr := yamlReader{
				service: tt.fields.service,
			}
			err := yr.Parse(tt.args.data, &tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlReader.Parse() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(t, tt.want, tt.args.output)
			}
		})
	}
}
