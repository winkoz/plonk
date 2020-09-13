package io

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

func Test_stitcher_Stitch(t *testing.T) {
	fixturesPath := "../fixtures/stitcher"
	targetPath := "/tmp/plonk/tests/deploy"
	sharedtesting.DeletePath(targetPath)
	sharedtesting.CreatePath(targetPath)
	type args struct {
		sourcePath     string
		targetPath     string
		targetFilename string
		filePaths      []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "succeeds in merging all sources into a single output.yml file",
			args: args{
				sourcePath:     fixturesPath,
				targetPath:     targetPath,
				targetFilename: "output.yml",
				filePaths: []string{
					"1.yaml",
					"2.yaml",
					"3.yaml",
				},
			},
			wantErr: false,
			want: `- file1
- file1_line2
- name: service
- line 1
- line 2
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stitcher{}
			if err := s.Stitch(tt.args.sourcePath, tt.args.targetPath, tt.args.targetFilename, tt.args.filePaths); (err != nil) != tt.wantErr {
				t.Errorf("stitcher.Stitch() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				outputPath := fmt.Sprintf("%s/%s", tt.args.targetPath, tt.args.targetFilename)
				outputBytes, err := ioutil.ReadFile(outputPath)
				if err != nil {
					t.Errorf("stitcher.Stitch() output_read_error = %v", err)
				}
				outputContents := string(outputBytes)
				assert.Equal(t, tt.want, outputContents)
			}
		})
	}
}
