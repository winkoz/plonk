package scaffolding

import (
	"reflect"
	"testing"

	"github.com/winkoz/plonk/internal/io"
)

func Test_templateReader_Read(t *testing.T) {
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
		want    []string
		wantErr bool
	}{
		{
			name: "successfully loads a template data file into ",
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
			if (err != nil) != tt.wantErr {
				t.Errorf("templateReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templateReader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
