package scaffolding

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/network"
)

func Test_templateReader_Read(t *testing.T) {
	defaultTemplatePath := io.BinaryFile + "/templates"
	customTemplatePath := "../fixtures/templateReader/customTemplates"
	ctx := config.Context{
		CustomTemplatesPath: customTemplatePath,
	}
	invalidTemplateData, _ := ioutil.ReadFile(customTemplatePath + "/invalid/template-definition.yaml")
	invalidTemplate := string(invalidTemplateData)
	ioService := io.NewService()
	networkService := network.NewService()
	yamlReader := io.NewYamlReader(ioService)
	type fields struct {
		ctx            config.Context
		yamlReader     io.YamlReader
		ioService      io.Service
		networkService network.Service
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
		// TODO: https://github.com/winkoz/plonk/issues/57 (jseravalli): Add failing files test for default template
		{
			name: "successfully loads a template data file located in the default template folder into a TemplateData structure",
			fields: fields{
				ctx:            ctx,
				ioService:      ioService,
				networkService: networkService,
				yamlReader:     yamlReader,
			},
			args: args{
				templateName: "default",
			},
			want: TemplateData{
				Name: "default",
				FilesLocation: []io.FileLocation{
					{
						OriginalFilePath: "plonk.yaml",
						ResolvedFilePath: defaultTemplatePath + "/default/plonk.yaml",
					},
					{
						OriginalFilePath: "deploy/variables/production.yaml",
						ResolvedFilePath: defaultTemplatePath + "/default/deploy/variables/production.yaml",
					},
					{
						OriginalFilePath: "deploy/variables/base.yaml",
						ResolvedFilePath: defaultTemplatePath + "/default/deploy/variables/base.yaml",
					},
					{
						OriginalFilePath: "deploy/secrets/production.yaml",
						ResolvedFilePath: defaultTemplatePath + "/default/deploy/secrets/production.yaml",
					},
					{
						OriginalFilePath: "deploy/secrets/base.yaml",
						ResolvedFilePath: defaultTemplatePath + "/default/deploy/secrets/base.yaml",
					},
				},
				Files: []string{
					"plonk.yaml",
					"deploy/variables/production.yaml",
					"deploy/variables/base.yaml",
					"deploy/secrets/production.yaml",
					"deploy/secrets/base.yaml",
				},
				Manifests: []string{},
			},
			wantErr: nil,
		},
		{
			name: "successfully loads a template data file located in the custom template folder into a TemplateData structure",
			fields: fields{
				ctx:            ctx,
				ioService:      ioService,
				networkService: networkService,
				yamlReader:     yamlReader,
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
				Files:     []string{"plonk.yaml"},
				Manifests: []string{},
				DefaultVariables: struct {
					Build       map[string]string `yaml:"build,omitempty"`
					Environment map[string]string `yaml:"environment,omitempty"`
				}{
					Build: map[string]string{
						"TEST_BUILD_VAR": "custom-template-build",
					},
					Environment: map[string]string{
						"TEST_ENV_VAR": "custom-template-env",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "returns an error when the configuration file cannot be located",
			fields: fields{
				ctx:            ctx,
				ioService:      ioService,
				networkService: networkService,
				yamlReader:     yamlReader,
			},
			args: args{
				templateName: "non-existent-config-file",
			},
			want: TemplateData{
				FilesLocation: []io.FileLocation{},
				Files:         []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found template-definition.yaml. Locations [%s]", customTemplatePath)),
		},
		{
			name: "returns an error when the configuration file is invalid",
			fields: fields{
				ctx:            ctx,
				ioService:      ioService,
				networkService: networkService,
				yamlReader:     yamlReader,
			},
			args: args{
				templateName: "invalid",
			},
			want: TemplateData{
				FilesLocation: []io.FileLocation{},
				Files:         []string{},
			},
			wantErr: io.NewParseYamlError(fmt.Sprintf("Unable to parse %s", invalidTemplate)),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within the default path",
			fields: fields{
				ctx:            ctx,
				ioService:      ioService,
				networkService: networkService,
				yamlReader:     yamlReader,
			},
			args: args{
				templateName: "missingFiles",
			},
			want: TemplateData{
				FilesLocation: []io.FileLocation{},
				Files:         []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found template-definition.yaml. Locations [%s]", customTemplatePath)),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within scripts",
			fields: fields{
				ctx:            ctx,
				ioService:      ioService,
				networkService: networkService,
				yamlReader:     yamlReader,
			},
			args: args{
				templateName: "missingManifests",
			},
			want: TemplateData{
				Name:          "missing-manifests",
				FilesLocation: []io.FileLocation{},
				Files:         []string{},
				Manifests:     []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found missingFile.yaml. Locations [%s]", customTemplatePath)),
		},
		{
			name: "returns an error when the configuration file points to a non-existent file within scripts",
			fields: fields{
				ctx:            ctx,
				ioService:      ioService,
				networkService: networkService,
				yamlReader:     yamlReader,
			},
			args: args{
				templateName: "missingManifests",
			},
			want: TemplateData{
				Name:          "missing-manifests",
				FilesLocation: []io.FileLocation{},
				Files:         []string{},
				Manifests:     []string{},
			},
			wantErr: NewScaffolderFileNotFound(fmt.Sprintf("Template not found missingFile.yaml. Locations [%s]", customTemplatePath)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := templateReader{
				ctx:            ctx,
				service:        tt.fields.ioService,
				networkService: tt.fields.networkService,
				yamlReader:     tt.fields.yamlReader,
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
