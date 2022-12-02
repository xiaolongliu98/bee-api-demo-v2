package response

import (
	"bee-api-demo/models/response"
	"reflect"
	"testing"
)

func TestAddResponse(t *testing.T) {
	type args struct {
		r *response.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := response.AddResponse(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("AddResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetResponseById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *response.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := response.GetResponseById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResponseById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResponseById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
