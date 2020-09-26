package scaffolding

import (
	"testing"
)

func Test_scaffolder_Install(t *testing.T) {
	type fields struct {
		targetPath    string
		sourcePath    string
		templatePaths []string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scaffolder{
				targetPath:    tt.fields.targetPath,
				sourcePath:    tt.fields.sourcePath,
				templatePaths: tt.fields.templatePaths,
			}
			if err := s.Install(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("scaffolder.Install() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
