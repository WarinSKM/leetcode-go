package addtwointegers

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				num1: 4,
				num2: 8,
			},
			want: 12,
		},
		{
			name: "test case 2",
			args: args{
				num1: -4,
				num2: 8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sum(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
