package utils

import (
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "json_to_byte",
			args: args{
				value: `{"value": "value",
				"numbers":1234,
				"float:124234.89798,
				"array":[{
					name: "value",
					numbers:1234,
				}]}`,
			},
			want: []byte{255, 128, 12, 0, 125, 123, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 118, 97, 108, 117, 101, 34, 44, 10, 9, 9, 9, 9, 34, 110, 117, 109, 98, 101, 114, 115, 34, 58, 49, 50, 51, 52, 44, 10, 9, 9, 9, 9, 34, 102, 108, 111, 97, 116, 58, 49, 50, 52, 50, 51, 52, 46, 56, 57, 55, 57, 56, 44, 10, 9, 9, 9, 9, 34, 97, 114, 114, 97, 121, 34, 58, 91, 123, 10, 9, 9, 9, 9, 9, 110, 97, 109, 101, 58, 32, 34, 118, 97, 108, 117, 101, 34, 44, 10, 9, 9, 9, 9, 9, 110, 117, 109, 98, 101, 114, 115, 58, 49, 50, 51, 52, 44, 10, 9, 9, 9, 9, 125, 93, 125},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Serialize(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeserialize(t *testing.T) {
	type args struct {
		byt []byte
		ptr interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "byte_to_json",
			args: args{
				byt: []byte{66, 12, 0, 63, 123, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 118, 97, 108, 117, 101, 34, 44, 10, 9, 9, 9, 9, 34, 110, 117, 109, 98, 101, 114, 115, 34, 58, 49, 50, 51, 52, 44, 10, 9, 9, 9, 9, 34, 102, 108, 111, 97, 116, 58, 49, 50, 52, 50, 51, 52, 46, 56, 57, 55, 57, 56, 125},
				ptr: `{"value": "value"}`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Deserialize(tt.args.byt, tt.args.ptr); (err != nil) != tt.wantErr {
				t.Errorf("Deserialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
