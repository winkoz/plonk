package scaffolding

import (
	"testing"

	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/sharedtesting"
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
		ioService                *sharedtesting.IOServiceMock
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
	type wantIOServiceMock struct {
		shouldTest            bool
		createPathReturn      error
		directoryExistsReturn bool
		walkReturn            error
		paramFullPath         string
		err                   error
	}

	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		wantErr                bool
		wantTemplateReaderMock wantTemplateReaderMock
		wantDuplicatorMock     wantDuplicatorMock
		wantIOServiceMock      wantIOServiceMock
	}{
		{
			name: "succesfully scaffolds the default template",
			fields: fields{
				targetPath:               "/tmp/plonk/tests/scripts",
				customTemplatePath:       "../fixtures/scripts",
				templateReader:           new(templateReaderMock),
				duplicator:               new(io.DuplicatorMock),
				destinationDeployDirName: "deploy",
				destinationVariablesPath: "deploy/variables",
				ioService:                new(sharedtesting.IOServiceMock),
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
			wantIOServiceMock: wantIOServiceMock{
				shouldTest:            true,
				directoryExistsReturn: true,
				createPathReturn:      nil,
				walkReturn:            nil,
				paramFullPath:         "/tmp/plonk/tests/scripts/deploy/variables",
				err:                   nil,
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
				ioService:                tt.fields.ioService,
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
			if tt.wantIOServiceMock.shouldTest {
				tt.fields.ioService.On(
					"DirectoryExists",
					tt.wantIOServiceMock.paramFullPath,
				).Return(
					tt.wantIOServiceMock.directoryExistsReturn,
				)
				tt.fields.ioService.On(
					"CreatePath",
					tt.fields.targetPath,
				).Return(
					tt.wantIOServiceMock.createPathReturn,
				)
				tt.fields.ioService.On(
					"Walk",
					tt.wantIOServiceMock.paramFullPath,
				).Return(
					tt.wantIOServiceMock.walkReturn,
				)
			}
			if err := s.Install(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("scaffolder.Install() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
