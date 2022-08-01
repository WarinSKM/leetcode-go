package minimumnumberofdecibinarynumbers

import "testing"

func Test_minPartitions(t *testing.T) {
	type args struct {
		n string
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
				n: "32",
			},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{
				n: "82734",
			},
			want: 8,
		},
		{
			name: "test case 3",
			args: args{
				n: "27346209830709182346",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := minPartitions(tt.args.n); got != tt.want {
				t.Errorf("minPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
