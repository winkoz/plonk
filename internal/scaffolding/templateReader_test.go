package scaffolding

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/winkoz/plonk/internal/io"
)

func Test_templateReader_Read(t *testing.T) {
	fixturesPath := "../fixtures/templateReader"
	yamlReader := io.NewYamlReader()
	type fields struct {
		defaultTemplatePath string
		customTemplatePath  string
		yamlReader          io.YamlReader
	}
	type args struct {
		configurationFileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    TemplateData
		wantErr error
	}{
		{
			name: "successfully loads a template data file into a TemplateData structure",
			fields: fields{
				defaultTemplatePath: fixturesPath,
				customTemplatePath:  fixturesPath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "base",
			},
			want: TemplateData{
				Name: "base",
				Origin: []string{
					fixturesPath + "/base/plonk.yaml",
				},
				Variables: []string{
					fixturesPath + "/base/base.yaml",
				},
				Scripts: []string{
					fixturesPath + "/base/ingress.yaml",
				},
			},
			wantErr: nil,
		},
		{
			name: "returns an error when the configuration file cannot be located",
			fields: fields{
				defaultTemplatePath: fixturesPath,
				customTemplatePath:  fixturesPath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "non-existent-config-file",
			},
			want:    TemplateData{},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found non-existent-config-file.yaml. Locations [%s, %s]", fixturesPath, fixturesPath)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := templateReader{
				defaultTemplatePath: tt.fields.defaultTemplatePath,
				customTemplatePath:  tt.fields.customTemplatePath,
				yamlReader:          tt.fields.yamlReader,
			}
			got, err := tf.Read(tt.args.configurationFileName)
			if (tt.wantErr == nil && err != nil) || (tt.wantErr != nil && err == nil) {
				t.Errorf("templateReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (tt.wantErr != nil && err != nil) && tt.wantErr.Error() != err.Error() {
				t.Errorf("templateReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templateReader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
