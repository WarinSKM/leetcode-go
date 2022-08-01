package finalvalueofvariableafterperformingoperations

import "testing"

func Test_finalValueAfterOperations(t *testing.T) {
	type args struct {
		operations []string
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
				operations: []string{"--X", "X++", "X++"},
			},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{
				operations: []string{"++X", "++X", "X++"},
			},
			want: 3,
		},
		{
			name: "test case 3",
			args: args{
				operations: []string{"X++", "++X", "--X", "X--"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := finalValueAfterOperations(tt.args.operations); got != tt.want {
				t.Errorf("finalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
