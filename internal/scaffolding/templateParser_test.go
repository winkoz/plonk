package scaffolding

import "testing"

func Test_templateParser_Parse(t *testing.T) {
	type args struct {
		variables map[string]string
		content   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successfully replaces all keys from the map with their respective values on the contents parameter",
			args: args{
				variables: map[string]string{
					"key":  "value",
					"key2": "value2",
				},
				content: "{{.key}} == value and value2 == {{.key2}}",
			},
			wantErr: false,
			want:    "value == value and value2 == value2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := templateParser{}
			got, err := tp.Parse(tt.args.variables, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("templateParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("templateParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
