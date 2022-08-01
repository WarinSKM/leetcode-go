package strongpasswordchecker

import "testing"

func Test_strongPasswordChecker(t *testing.T) {
	type args struct {
		password string
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
				password: "a",
			},
			want: 5,
		},
		{
			name: "test case 2",
			args: args{
				password: "aA1",
			},
			want: 3,
		},
		{
			name: "test case 3",
			args: args{
				password: "1337C0d3",
			},
			want: 0,
		},
		{
			name: "test case 4",
			args: args{
				password: "aaaB1",
			},
			want: 1,
		},
		{
			name: "test case 5",
			args: args{
				password: "1111111111",
			},
			want: 3,
		},
		{
			name: "test case 6",
			args: args{
				password: "aaa111",
			},
			want: 2,
		},
		{
			name: "test case 7",
			args: args{
				password: "aaa123",
			},
			want: 1,
		},
		{
			name: "test case 8",
			args: args{
				password: "ABABABABABABABABABAB1",
			},
			want: 2,
		},
		{
			name: "test case 9",
			args: args{
				password: "bbaaaaaaaaaaaaaaacccccc", // repeat 7 needto - 3
			},
			want: 8,
		},
		{
			name: "test case 10",
			args: args{
				password: "aaaaAAAAAA000000123456",
			},
			want: 5,
		},
		{
			name: "test case 11",
			args: args{
				password: "aabbx",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := strongPasswordChecker(tt.args.password); got != tt.want {
				t.Errorf("\nstrongPasswordChecker() = %v, want %v\n\n", got, tt.want)
			}
		})
	}
	// for _, tt := range tests {
	// 	tt := tt
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		t.Parallel()
	// 		got1, got2, got3 := strongPasswordChecker(tt.args.password)
	// 		// if got != tt.want {
	// 		t.Errorf("\nstrongPasswordChecker() =\n lengthCase: %v,\n charCase: %v,\n repeat: %v, want %v\n\n", got1, got2, got3, tt.want)
	// 		// }
	// 	})
	// }
}
