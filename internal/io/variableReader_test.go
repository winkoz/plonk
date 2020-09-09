package io

import (
	"reflect"
	"testing"
)

func Test_variableReader_GetVariables(t *testing.T) {
	sut := NewVariableReader()
	type args struct {
		stackName string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr error
	}{
		{
			name: "File doesn't exist should return a ParseVariableError",
			args: args{
				stackName: "production",
			},
			want:    nil,
			wantErr: NewParseVariableError("File not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sut.GetVariables(tt.args.stackName)
			if err != tt.wantErr {
				t.Errorf("variableReader.GetVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("variableReader.GetVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}
