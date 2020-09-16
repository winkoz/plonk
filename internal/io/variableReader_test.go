package io

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_variableReader_GetVariables(t *testing.T) {
	fixturesDir := "../fixtures"
	yamlReader := NewYamlReader()
	type args struct {
		stackName string
	}
	type fields struct {
		path           string
		baseFileName   string
		customFileName string
		yamlReader     YamlReader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]string
		wantErr error
	}{
		{
			name: "return a ParseVariableError with 'file not found' when file doesn't exist",
			fields: fields{
				path:         fixturesDir,
				baseFileName: "notFound",
				yamlReader:   yamlReader,
			},
			args: args{
				stackName: "production",
			},
			want:    nil,
			wantErr: NewParseVariableError(fmt.Sprintf("notFound.yaml not found at location: %s", fixturesDir)),
		},
		{
			name: "returns ParseVariableError with 'unable to parse file' when yaml file is wrong",
			fields: fields{
				path:           fixturesDir,
				baseFileName:   "base",
				customFileName: "invalidYaml",
				yamlReader:     yamlReader,
			},
			args: args{
				stackName: "production",
			},
			want:    nil,
			wantErr: NewParseVariableError(fmt.Sprintf("Unable to parse %s/invalidYaml.yaml", fixturesDir)),
		},
		{
			name: "uses the value of the custom variables when the same key is present in the base and the stack configurations",
			fields: fields{
				path:           fixturesDir,
				baseFileName:   "base",
				customFileName: "test",
				yamlReader:     yamlReader,
			},
			args: args{
				stackName: "producution",
			},
			want: map[string]string{
				"profile_pictures": "profile-pictures-dev",
				"second_val":       "merged_value",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := variableReader{
				path:           tt.fields.path,
				customFileName: tt.fields.customFileName,
				baseFileName:   tt.fields.baseFileName,
				yamlReader:     tt.fields.yamlReader,
			}
			got, err := sut.GetVariables(tt.args.stackName)
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
