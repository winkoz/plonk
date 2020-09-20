package scaffolding

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/winkoz/plonk/internal/io"
)

func Test_templateReader_Read(t *testing.T) {
	defaultTemplatePath := "../fixtures/templateReader"
	customTemplatePath := "../fixtures/templateReader/customTemplate"
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
			name: "successfully loads a template data file located in the default template folder into a TemplateData structure",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "base",
			},
			want: TemplateData{
				Name: "base",
				Origin: []string{
					defaultTemplatePath + "/base/plonk.yaml",
				},
				Variables: []string{
					defaultTemplatePath + "/base/base.yaml",
				},
				Manifests: []string{
					defaultTemplatePath + "/base/ingress.yaml",
				},
			},
			wantErr: nil,
		},
		{
			name: "successfully loads a template data file located in the custom template folder into a TemplateData structure",
			fields: fields{
				defaultTemplatePath: customTemplatePath,
				customTemplatePath:  defaultTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "custom",
			},
			want: TemplateData{
				Name: "custom",
				Origin: []string{
					customTemplatePath + "/custom/plonk.yaml",
				},
				Variables: []string{
					customTemplatePath + "/custom/base.yaml",
				},
				Manifests: []string{
					customTemplatePath + "/custom/ingress.yaml",
				},
			},
			wantErr: nil,
		},
		{
			name: "successfully loads a template data file with files from custom & default template folders into a TemplateData structure",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "base_custom",
			},
			want: TemplateData{
				Name: "base_custom",
				Origin: []string{
					defaultTemplatePath + "/base/plonk.yaml",
				},
				Variables: []string{
					customTemplatePath + "/custom/base.yaml",
				},
				Manifests: []string{
					defaultTemplatePath + "/base/ingress.yaml",
				},
			},
			wantErr: nil,
		},
		{
			name: "returns an error when the configuration file cannot be located",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "non-existent-config-file",
			},
			want:    TemplateData{},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found non-existent-config-file.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
		},
		{
			name: "returns an error when the configuration file is invalid",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "invalidYaml",
			},
			want:    TemplateData{},
			wantErr: io.NewParseYamlError(fmt.Sprintf("Unable to parse %s", defaultTemplatePath+"/invalidYaml.yaml")),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within origin",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "missingOriginFiles",
			},
			want: TemplateData{
				Name:   "base",
				Origin: []string{},
				Variables: []string{
					"base/base.yaml",
				},
				Manifests: []string{
					"base/ingress.yaml",
				},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found base/missingOriginFile.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within variables",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "missingVariablesFiles",
			},
			want: TemplateData{
				Name: "base",
				Origin: []string{
					defaultTemplatePath + "/base/base.yaml",
				},
				Variables: []string{},
				Manifests: []string{
					defaultTemplatePath + "/base/ingress.yaml",
				},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found base/missingVariableFile.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within scripts",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				configurationFileName: "missingScriptsFiles",
			},
			want: TemplateData{
				Name: "base",
				Origin: []string{
					defaultTemplatePath + "/base/ingress.yaml",
				},
				Variables: []string{
					"base/base.yaml",
				},
				Manifests: []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found base/missingScriptFile.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
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
