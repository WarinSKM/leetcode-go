package maximumwordsfoundinsentences

import "testing"

func Test_mostWordsFound(t *testing.T) {
	type args struct {
		sentences []string
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
				sentences: []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
			},
			want: 6,
		},
		{
			name: "test case 2",
			args: args{
				sentences: []string{"please wait", "continue to fight", "continue to win"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := mostWordsFound(tt.args.sentences); got != tt.want {
				t.Errorf("mostWordsFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
