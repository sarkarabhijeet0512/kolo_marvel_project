package server

import (
	"testing"
)

func Test_inLambda(t *testing.T) {
	type args struct {
		IsinLambda string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "this_should_be_true",
			args: args{
				IsinLambda: "LAMBDA_TASK_ROOT",
			},
			want: true,
		},
		{
			name: "this_should_be_true",
			args: args{
				IsinLambda: "IS_NOT_LAMBDA_TASK_ROOT",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inLambda(tt.args.IsinLambda); got != tt.want {
				t.Errorf("inLambda() = %v, want %v", got, tt.want)
			}
		})
	}
}
