package scaffolding

import (
	"testing"

	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

func Test_scriptsGenerator_InitProject(t *testing.T) {
	sourcePath := "../fixtures/scripts"
	targetPath := "/tmp/plonk/tests/scripts"
	projectName := "plonkTests"
	projectDefinition := BaseProjectFiles

	type fields struct {
		sourcePath   string
		targetPath   string
		duplicator   *io.DuplicatorMock
		interpolator *sharedtesting.InterpolatorMock
		stitcher     *sharedtesting.StitcherMock
	}
	type args struct {
		projectName  string
		templateName string
	}
	type wantInterpolator struct {
		testInterpolator   bool
		source             map[string]string
		target             string
		interpolatedResult string
		returnErr          error
	}
	type wantDuplicator struct {
		testDuplicator bool
		targetPath     string
		sources        []string
		returnErr      error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantErr          bool
		wantInterpolator wantInterpolator
		wantDuplicator   wantDuplicator
	}{
		{
			name: "succeeds by generating the boilerplate file structure",
			fields: fields{
				sourcePath:   sourcePath,
				targetPath:   targetPath,
				duplicator:   new(io.DuplicatorMock),
				interpolator: new(sharedtesting.InterpolatorMock),
			},
			args: args{
				projectName:  projectName,
				templateName: "base",
			},
			wantErr: false,
			wantInterpolator: wantInterpolator{
				testInterpolator: true,
				source:           map[string]string{"PROJECT_NAME": projectName},
				target:           targetPath,
				returnErr:        nil,
			},
			wantDuplicator: wantDuplicator{
				testDuplicator: true,
				targetPath:     targetPath,
				sources:        projectDefinition,
				returnErr:      nil,
			},
		},
	}
	for _, tt := range tests {
		sharedtesting.DeletePath(targetPath)
		sharedtesting.CreatePath(targetPath)

		t.Run(tt.name, func(t *testing.T) {
			s := scriptsGenerator{
				customTemplatesPath: tt.fields.sourcePath,
				targetPath:          tt.fields.targetPath,
				interpolator:        tt.fields.interpolator,
				duplicator:          tt.fields.duplicator,
			}

			if tt.wantInterpolator.testInterpolator {
				for _, source := range tt.args.templateName {
					tt.fields.interpolator.On(
						"SubstituteValues",
						tt.wantInterpolator.source,
						source,
					).Return(
						source, tt.wantInterpolator.returnErr,
					)
				}
			}

			if tt.wantDuplicator.testDuplicator {
				tt.fields.duplicator.On(
					"CopyMultiple",
					tt.wantDuplicator.targetPath,
					tt.wantDuplicator.sources,
				).Return(
					tt.wantDuplicator.returnErr,
				)
			}

			if err := s.ScaffoldTemplate(tt.args.projectName, tt.args.templateName); (err != nil) != tt.wantErr {
				t.Errorf("scriptsGenerator.InitProject() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				if tt.wantInterpolator.testInterpolator {
					tt.fields.interpolator.AssertExpectations(t)
				}

				if tt.wantDuplicator.testDuplicator {
					tt.fields.duplicator.AssertExpectations(t)
				}
			}
		})
	}
}
