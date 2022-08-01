package concatarray

import (
	"reflect"
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2, 3, 1, 2, 3},
		},
		{
			name: "test case 2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4, 1, 2, 3, 4},
		},
		{
			name: "test case 1",
			args: args{
				nums: []int{1, 3, 2, 1},
			},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := getConcatenation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}
