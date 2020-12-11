package cmd

import "testing"

func TestExecute(t *testing.T) {
	type args struct {
		args          []string
		versionNumber string
		buildTime     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid argument",
			args: args{
				args:          []string{"-h"},
				versionNumber: "0.0.0",
				buildTime:     "now",
			},
		},
		{
			name: "invalid argument",
			args: args{
				args: []string{"", "-Z"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Execute(tt.args.args, tt.args.versionNumber, tt.args.buildTime); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
