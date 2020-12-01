package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_secretReader_GetSecretFromFile(t *testing.T) {
	fixturesDir := "../fixtures/secrets"
	service := NewService()
	yamlReader := NewYamlReader(service)
	type args struct {
		projectName string
		env         string
	}
	type fields struct {
		path         string
		baseFileName string
		yamlReader   YamlReader
		service      Service
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    DeploySecrets
		wantErr ErrorCode
	}{
		{
			name: "return a FileNotFound error when the base file doesn't exist",
			fields: fields{
				path:         fixturesDir,
				baseFileName: "notFound",
				yamlReader:   yamlReader,
				service:      service,
			},
			args: args{
				projectName: "plonk-test",
				env:         "production",
			},
			want:    DeploySecrets{},
			wantErr: FileNotFoundError,
		},
		{
			name: "returns ParseSecretError when the yaml file is invalid",
			fields: fields{
				path:         fixturesDir,
				baseFileName: "invalidYaml",
				yamlReader:   yamlReader,
				service:      service,
			},
			args: args{
				projectName: "plonk-test",
				env:         "production",
			},
			want:    DeploySecrets{},
			wantErr: ParseSecretError,
		},
		{
			name: "uses the value of the env file when the same key is present in the base",
			fields: fields{
				path:         fixturesDir,
				baseFileName: "base",
				yamlReader:   yamlReader,
				service:      service,
			},
			args: args{
				projectName: "plonk-test",
				env:         "production",
			},
			want: DeploySecrets{
				Secret: map[string]string{
					"DOMAIN":     "production.example.com",
					"FILES_PATH": "/tmp/files",
				},
			},
			wantErr: NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := secretReader{
				path:         tt.fields.path,
				baseFileName: tt.fields.baseFileName,
				yamlReader:   tt.fields.yamlReader,
				service:      tt.fields.service,
			}
			got, err := sut.GetSecretFromFile(tt.args.projectName, tt.args.env)
			if (tt.wantErr != NoError && err == nil) || (tt.wantErr == NoError && err != nil) {
				t.Errorf("secretReader.GetSecretFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != NoError && err != nil && err.(*Error).Code() != tt.wantErr {
				t.Errorf("secretReader.GetSecretFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
