package io

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_variableReader_GetVariables(t *testing.T) {
	fixturesDir := "../fixtures"
	type args struct {
		stackName string
	}
	tests := []struct {
		name    string
		sut     VariableReader
		args    args
		want    map[string]string
		wantErr error
	}{
		{
			name: "File doesn't exist should return a ParseVariableError with file not found",
			sut:  variableReader{path: fixturesDir, baseFileName: "notFound"},
			args: args{
				stackName: "production",
			},
			want:    nil,
			wantErr: NewParseVariableError(fmt.Sprintf("notFound.yaml not found at location: %s", fixturesDir)),
		},
		{
			name: "When yaml file is wrong returns ParseVariableError with unable to parse file",
			sut: variableReader{
				path:           fixturesDir,
				baseFileName:   "base",
				customFileName: "invalidYAML",
			},
			args: args{
				stackName: "production",
			},
			want:    nil,
			wantErr: NewParseVariableError(fmt.Sprintf("Unable to parse %s/invalidYAML.yaml", fixturesDir)),
		},
		// TODO: Add ALL the missing test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sut.GetVariables(tt.args.stackName)
			if (tt.wantErr != nil && err == nil) || (tt.wantErr == nil && err != nil) {
				t.Errorf("variableReader.GetVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != nil && err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("variableReader.GetVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("variableReader.GetVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}
