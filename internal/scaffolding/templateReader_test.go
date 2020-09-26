package scaffolding

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/winkoz/plonk/internal/io"
)

func Test_templateReader_Read(t *testing.T) {
	defaultTemplatePath := "../fixtures/templateReader/defaultTemplates"
	customTemplatePath := "../fixtures/templateReader/customTemplates"
	yamlReader := io.NewYamlReader()
	type fields struct {
		defaultTemplatePath string
		customTemplatePath  string
		yamlReader          io.YamlReader
	}
	type args struct {
		templateName string
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
				templateName: "default",
			},
			want: TemplateData{
				Name: "default-template",
				FilesLocation: []io.FileLocation{
					{
						ResolvedFilePath: defaultTemplatePath + "/default/plonk.yaml",
						OriginalFilePath: "plonk.yaml",
					},
				},
				Files:             []string{"plonk.yaml"},
				VariablesFileName: defaultTemplatePath + "/default/vars.yaml",
				VariablesContents: `plonk:
  - NAME: "$NAME"
  - APP_ENV: "$APP_ENV"
  - NAMESPACE: "$STACK-$NAME"`,
				Manifests: []string{},
			},
			wantErr: nil,
		},
		{
			name: "successfully loads a template data file located in the custom template folder into a TemplateData structure",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				templateName: "custom",
			},
			want: TemplateData{
				Name: "custom-template",
				FilesLocation: []io.FileLocation{
					{
						ResolvedFilePath: customTemplatePath + "/custom/plonk.yaml",
						OriginalFilePath: "plonk.yaml",
					},
				},
				Files:             []string{"plonk.yaml"},
				VariablesFileName: customTemplatePath + "/custom/vars.yaml",
				VariablesContents: `plonk:
  - NAME: "custom-name"
  - APP_ENV: "$APP_ENV"
  - NAMESPACE: "$STACK-$NAME"`,
				Manifests: []string{},
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
				templateName: "mixed",
			},
			want: TemplateData{
				Name:              "custom-mixed-template",
				FilesLocation:     []io.FileLocation{},
				Files:             []string{},
				VariablesFileName: defaultTemplatePath + "/mixed/vars.yaml",
				VariablesContents: `plonk:
  - NAME: "mixed-name"
  - APP_ENV: "$APP_ENV"
  - NAMESPACE: "$STACK-$NAME"`,
				Manifests: []string{
					customTemplatePath + "/mixed/manifest3.yaml",
					defaultTemplatePath + "/mixed/manifest2.yaml",
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
				templateName: "non-existent-config-file",
			},
			want: TemplateData{
				FilesLocation: []io.FileLocation{},
				Files:         []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found template-definition.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
		},
		{
			name: "returns an error when the configuration file is invalid",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				templateName: "invalid",
			},
			want: TemplateData{
				FilesLocation: []io.FileLocation{},
				Files:         []string{},
			},
			wantErr: io.NewParseYamlError(fmt.Sprintf("Unable to parse %s", defaultTemplatePath+"/invalid/template-definition.yaml")),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within the default path",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				templateName: "missingFiles",
			},
			want: TemplateData{
				FilesLocation: []io.FileLocation{},
				Files:         []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found template-definition.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within variables",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				templateName: "missingVariables",
			},
			want: TemplateData{
				Name:              "missing-variables",
				FilesLocation:     []io.FileLocation{},
				Files:             []string{},
				VariablesFileName: "",
				Manifests:         []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found vars.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within scripts",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				templateName: "missingManifests",
			},
			want: TemplateData{
				Name:              "missing-manifests",
				FilesLocation:     []io.FileLocation{},
				Files:             []string{},
				VariablesFileName: "",
				Manifests:         []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found missingFile.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within scripts",
			fields: fields{
				defaultTemplatePath: defaultTemplatePath,
				customTemplatePath:  customTemplatePath,
				yamlReader:          yamlReader,
			},
			args: args{
				templateName: "missingManifests",
			},
			want: TemplateData{
				Name:              "missing-manifests",
				FilesLocation:     []io.FileLocation{},
				Files:             []string{},
				VariablesFileName: "",
				Manifests:         []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found missingFile.yaml. Locations [%s, %s]", customTemplatePath, defaultTemplatePath)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := templateReader{
				defaultTemplatePath: tt.fields.defaultTemplatePath,
				customTemplatePath:  tt.fields.customTemplatePath,
				yamlReader:          tt.fields.yamlReader,
			}
			got, err := tf.Read(tt.args.templateName)
			if (tt.wantErr == nil && err != nil) || (tt.wantErr != nil && err == nil) {
				t.Errorf("templateReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (tt.wantErr != nil && err != nil) && tt.wantErr.Error() != err.Error() {
				t.Errorf("templateReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templateReader.Read() =\n\t%+v,\nwant\n\t%+v", got, tt.want)
			}
		})
	}
}
