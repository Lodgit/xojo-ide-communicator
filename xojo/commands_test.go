package xojo

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_checkForErrorResponse(t *testing.T) {
	type args struct {
		jsonb []byte
		err   error
	}
	responseOk := []byte("{\"tag\":\"build\",\"script\":\"OpenFile(\\\"%s\\\")\nprint \\\"Project is opened.\\\"\"}" + XojoNullChar)
	responseErr := []byte("{\"tag\":\"build\",\"response\":{\"openErrors\":[{\"projectError\":{}}]}\"}" + XojoNullChar)
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "no contain error response",
			args: args{
				jsonb: responseOk,
			},
			want: responseOk,
		},
		{
			name: "contain error response",
			args: args{
				jsonb: responseErr,
			},
			want:    responseErr,
			wantErr: true,
		},
		{
			name: "error param provided",
			args: args{
				jsonb: responseErr,
				err:   fmt.Errorf("ERR"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckForErrorResponse(tt.args.jsonb, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkForErrorResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkForErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
