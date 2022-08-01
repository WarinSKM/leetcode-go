package rootequalssumofchildren

import "testing"

func Test_checkTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: true,
		},
		{
			name: "test case 1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := checkTree(tt.args.root); got != tt.want {
				t.Errorf("checkTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
