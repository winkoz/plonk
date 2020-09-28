package scaffolding

import (
	"testing"

	"github.com/winkoz/plonk/internal/io"
)

func Test_scaffolder_Install(t *testing.T) {
	type fields struct {
		targetPath               string
		customTemplatePath       string
		templateReader           *templateReaderMock
		duplicator               *io.DuplicatorMock
		templatePaths            []string
		destinationDeployDirName string
		destinationVariablesPath string
	}
	type args struct {
		name string
	}
	type wantDuplicatorMock struct {
		shouldTest bool
		err        error
	}
	type wantTemplateReaderMock struct {
		shouldTest   bool
		templatedata TemplateData
		err          error
	}
	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		wantErr                bool
		wantTemplateReaderMock wantTemplateReaderMock
		wantDuplicatorMock     wantDuplicatorMock
	}{
		{
			name: "succesfully scaffolds the default template",
			fields: fields{
				targetPath:               "/tmp/plonk/tests/scripts",
				customTemplatePath:       "../fixtures/scripts",
				templateReader:           new(templateReaderMock),
				duplicator:               new(io.DuplicatorMock),
				destinationDeployDirName: "",
				destinationVariablesPath: "",
			},
			wantTemplateReaderMock: wantTemplateReaderMock{
				shouldTest: true,
				templatedata: TemplateData{
					Name:              "",
					Files:             []string{},
					FilesLocation:     []io.FileLocation{},
					VariablesFileName: "",
					VariablesContents: "",
					Manifests:         []string{},
				},
			},
			wantDuplicatorMock: wantDuplicatorMock{
				shouldTest: true,
				err:        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scaffolder{
				targetPath:               tt.fields.targetPath,
				customTemplatePath:       tt.fields.customTemplatePath,
				templateReader:           tt.fields.templateReader,
				duplicator:               tt.fields.duplicator,
				destinationDeployDirName: tt.fields.destinationDeployDirName,
				destinationVariablesPath: tt.fields.destinationVariablesPath,
			}
			if tt.wantTemplateReaderMock.shouldTest {
				tt.fields.templateReader.On(
					"Read",
					tt.args.name,
				).Return(
					tt.wantTemplateReaderMock.templatedata, tt.wantTemplateReaderMock.err,
				)
			}
			if tt.wantDuplicatorMock.shouldTest {
				tt.fields.duplicator.On(
					"CopyMultiple",
					tt.fields.targetPath,
					tt.wantTemplateReaderMock.templatedata.FilesLocation,
				).Return(
					tt.wantDuplicatorMock.err,
				)
			}
			if err := s.Install(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("scaffolder.Install() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
