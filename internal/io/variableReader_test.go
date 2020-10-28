package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_variableReader_GetVariablesFromFile(t *testing.T) {
	fixturesDir := "../fixtures/variables"
	service := NewService()
	yamlReader := NewYamlReader(service)
	interpolator := NewInterpolator()
	type args struct {
		projectName string
		env         string
	}
	type fields struct {
		path         string
		baseFileName string
		yamlReader   YamlReader
		interpolator Interpolator
		service      Service
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    DeployVariables
		wantErr ErrorCode
	}{
		{
			name: "return a FileNotFound error when the base file doesn't exist",
			fields: fields{
				path:         fixturesDir,
				baseFileName: "notFound",
				yamlReader:   yamlReader,
				interpolator: interpolator,
				service:      service,
			},
			args: args{
				projectName: "plonk-test",
				env:         "production",
			},
			want:    DeployVariables{},
			wantErr: FileNotFoundError,
		},
		{
			name: "returns ParseVariableError when the yaml file is invalid",
			fields: fields{
				path:         fixturesDir,
				baseFileName: "invalidYaml",
				yamlReader:   yamlReader,
				interpolator: interpolator,
				service:      service,
			},
			args: args{
				projectName: "plonk-test",
				env:         "production",
			},
			want:    DeployVariables{},
			wantErr: ParseVariableError,
		},
		{
			name: "uses the value of the env file when the same key is present in the base",
			fields: fields{
				path:         fixturesDir,
				baseFileName: "base",
				yamlReader:   yamlReader,
				interpolator: interpolator,
				service:      service,
			},
			args: args{
				projectName: "plonk-test",
				env:         "production",
			},
			want: DeployVariables{
				Build: map[string]string{
					"USE_LOAD_BALANCER": "true",
				},
				Environment: map[string]string{
					"DOMAIN":     "production.example.com",
					"FILES_PATH": "/tmp/files",
				},
			},
			wantErr: NoError,
		},
		{
			name: "uses the base file with interpolated variables when the env files doesn't exists",
			fields: fields{
				path:         fixturesDir,
				baseFileName: "base",
				yamlReader:   yamlReader,
				interpolator: interpolator,
				service:      service,
			},
			args: args{
				projectName: "plonk-test",
				env:         "staging",
			},
			want: DeployVariables{
				Build: map[string]string{
					"USE_LOAD_BALANCER": "true",
				},
				Environment: map[string]string{
					"DOMAIN":     "staging.plonk-test.example.com",
					"FILES_PATH": "/tmp/files",
				},
			},
			wantErr: NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := variableReader{
				path:         tt.fields.path,
				baseFileName: tt.fields.baseFileName,
				yamlReader:   tt.fields.yamlReader,
				interpolator: tt.fields.interpolator,
				service:      tt.fields.service,
			}
			got, err := sut.GetVariablesFromFile(tt.args.projectName, tt.args.env)
			if (tt.wantErr != NoError && err == nil) || (tt.wantErr == NoError && err != nil) {
				t.Errorf("variableReader.GetVariablesFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != NoError && err != nil && err.(*Error).Code() != tt.wantErr {
				t.Errorf("variableReader.GetVariablesFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
