package scaffolding

import (
	"fmt"
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
		appendReturn          error
		paramFullPath         string
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
			name: "succesfully install the default template",
			args: args{
				name: "test",
			},
			wantErr: false,
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
			},
		},
		{
			name: "successfully append variable contents to file",
			args: args{
				name: "test",
			},
			wantErr: false,
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
					VariablesContents: "variable contents",
					Manifests:         []string{},
				},
			},
			wantDuplicatorMock: wantDuplicatorMock{
				shouldTest: false,
			},
			wantIOServiceMock: wantIOServiceMock{
				shouldTest:    true,
				paramFullPath: "/tmp/plonk/tests/scripts/deploy/variables",
				appendReturn:  nil,
				walkReturn:    nil,
			},
		},
		{
			name: "fails when directory creation fails",
			args: args{
				name: "test",
			},
			wantErr: true,
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
				shouldTest: false,
			},
			wantIOServiceMock: wantIOServiceMock{
				shouldTest:            true,
				directoryExistsReturn: false,
				createPathReturn:      fmt.Errorf("Failed to create path"),
				walkReturn:            nil,
				paramFullPath:         "/tmp/plonk/tests/scripts/deploy/variables",
			},
		},
		{
			name: "fails when duplicator copy multiple fails",
			args: args{
				name: "test",
			},
			wantErr: true,
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
					Name:  "",
					Files: []string{},
					FilesLocation: []io.FileLocation{
						{
							OriginalFilePath: "",
							ResolvedFilePath: "",
						},
					},
					VariablesFileName: "",
					VariablesContents: "",
					Manifests:         []string{},
				},
			},
			wantDuplicatorMock: wantDuplicatorMock{
				shouldTest: true,
				err:        fmt.Errorf("Failed CopyMultiple"),
			},
			wantIOServiceMock: wantIOServiceMock{
				shouldTest:            true,
				directoryExistsReturn: true,
				paramFullPath:         "/tmp/plonk/tests/scripts/deploy/variables",
			},
		},
		{
			name: "fails when unable to append variables",
			args: args{
				name: "test",
			},
			wantErr: true,
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
				shouldTest: false,
			},
			wantIOServiceMock: wantIOServiceMock{
				shouldTest:    true,
				paramFullPath: "/tmp/plonk/tests/scripts/deploy/variables",
				walkReturn:    fmt.Errorf("Failed while walking"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
					tt.wantIOServiceMock.paramFullPath,
				).Return(
					tt.wantIOServiceMock.createPathReturn,
				)
				tt.fields.ioService.On(
					"Walk",
					tt.wantIOServiceMock.paramFullPath,
				).Return(
					tt.wantIOServiceMock.walkReturn,
				)
				tt.fields.ioService.On(
					"Append",
					tt.wantIOServiceMock.paramFullPath,
					"\n"+tt.wantTemplateReaderMock.templatedata.VariablesContents,
				).Return(
					tt.wantIOServiceMock.appendReturn,
				)
			}
			if err := s.Install(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("scaffolder.Install() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
