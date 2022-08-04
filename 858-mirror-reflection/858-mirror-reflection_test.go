package mirrorreflection

import "testing"

func Test_mirrorReflection(t *testing.T) {
	type args struct {
		p int
		q int
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
				p: 2,
				q: 1,
			},
			want: 2,
		},
		{
			name: "test case 2",
			args: args{
				p: 3,
				q: 1,
			},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{
				p: 3,
				q: 2,
			},
			want: 0,
		},
		{
			name: "test case 4",
			args: args{
				p: 13,
				q: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := mirrorReflection(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("mirrorReflection(%v, %v) = %v, want %v", tt.args.p, tt.args.q, got, tt.want)
			}
		})
	}
}
