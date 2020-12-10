package scaffolding

import (
	"fmt"
	"testing"

	"github.com/winkoz/plonk/internal/config"
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
	type mockIOServiceExpectation struct {
		createPathReturn      error
		directoryExistsReturn bool
		walkReturn            error
		appendReturn          error
		paramFullPath         string
	}
	type wantIOServiceMock struct {
		shouldTest   bool
		expectations []mockIOServiceExpectation
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
					Name:          "",
					Files:         []string{},
					FilesLocation: []io.FileLocation{},
					Manifests:     []string{},
				},
			},
			wantDuplicatorMock: wantDuplicatorMock{
				shouldTest: true,
				err:        nil,
			},
			wantIOServiceMock: wantIOServiceMock{
				shouldTest: true,
				expectations: []mockIOServiceExpectation{
					{
						directoryExistsReturn: true,
						createPathReturn:      nil,
						walkReturn:            nil,
						paramFullPath:         "/tmp/plonk/tests/scripts",
					},
					{
						directoryExistsReturn: true,
						createPathReturn:      nil,
						walkReturn:            nil,
						paramFullPath:         "/tmp/plonk/tests/scripts/deploy/variables",
					},
				},
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
					Name:          "",
					Files:         []string{},
					FilesLocation: []io.FileLocation{},
					Manifests:     []string{},
				},
			},
			wantDuplicatorMock: wantDuplicatorMock{
				shouldTest: false,
			},
			wantIOServiceMock: wantIOServiceMock{
				shouldTest: true,
				expectations: []mockIOServiceExpectation{
					{
						directoryExistsReturn: true,
						createPathReturn:      nil,
						walkReturn:            nil,
						paramFullPath:         "/tmp/plonk/tests/scripts",
					},
					{
						paramFullPath: "/tmp/plonk/tests/scripts/deploy/variables",
						appendReturn:  nil,
						walkReturn:    nil,
					},
				},
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
					Name:          "",
					Files:         []string{},
					FilesLocation: []io.FileLocation{},
					Manifests:     []string{},
				},
			},
			wantDuplicatorMock: wantDuplicatorMock{
				shouldTest: false,
			},
			wantIOServiceMock: wantIOServiceMock{
				shouldTest: true,
				expectations: []mockIOServiceExpectation{
					{
						directoryExistsReturn: false,
						createPathReturn:      fmt.Errorf("Failed to create path"),
						walkReturn:            nil,
						paramFullPath:         "/tmp/plonk/tests/scripts/deploy/variables",
					},
				},
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
					Manifests: []string{},
				},
			},
			wantDuplicatorMock: wantDuplicatorMock{
				shouldTest: true,
				err:        fmt.Errorf("Failed CopyMultiple"),
			},
			wantIOServiceMock: wantIOServiceMock{
				shouldTest: true,
				expectations: []mockIOServiceExpectation{
					{
						directoryExistsReturn: true,
						createPathReturn:      nil,
						walkReturn:            nil,
						paramFullPath:         "/tmp/plonk/tests/scripts",
					},
					{
						directoryExistsReturn: true,
						paramFullPath:         "/tmp/plonk/tests/scripts/deploy/variables",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scaffolder{
				ctx: config.Context{
					TargetPath:          tt.fields.targetPath,
					CustomTemplatesPath: tt.fields.customTemplatePath,
					DeployVariablesPath: tt.fields.destinationVariablesPath,
				},
				templateReader: tt.fields.templateReader,
				duplicator:     tt.fields.duplicator,

				ioService: tt.fields.ioService,
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
				for _, expectation := range tt.wantIOServiceMock.expectations {
					tt.fields.ioService.On(
						"DirectoryExists",
						expectation.paramFullPath,
					).
						Once().
						Return(
							expectation.directoryExistsReturn,
						)
					tt.fields.ioService.On(
						"CreatePath",
						expectation.paramFullPath,
					).
						Once().
						Return(
							expectation.createPathReturn,
						)
				}
			}
			if err := s.Install(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("scaffolder.Install() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
