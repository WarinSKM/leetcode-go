package subrectanglequeries

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		rectangle [][]int
	}
	tests := []struct {
		name string
		args args
		want SubrectangleQueries
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				[][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}},
			},
			want: SubrectangleQueries{
				rectangle: [][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}},
			},
		},
		{
			name: "test case 1",
			args: args{
				[][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
			},
			want: SubrectangleQueries{
				rectangle: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Constructor(tt.args.rectangle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubrectangleQueries_UpdateSubrectangle(t *testing.T) {
	type args struct {
		row1     int
		col1     int
		row2     int
		col2     int
		newValue int
	}
	tests := []struct {
		name string
		this *SubrectangleQueries
		args args
		want SubrectangleQueries
	}{
		// TODO: Add test cases.
		// [0,0,3,2,5]
		{
			name: "test case 1",
			this: &SubrectangleQueries{
				rectangle: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			},
			args: args{
				row1:     0,
				col1:     0,
				col2:     3,
				row2:     2,
				newValue: 5,
			},
			want: SubrectangleQueries{
				rectangle: [][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}, {5, 5, 5}},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.this.UpdateSubrectangle(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2, tt.args.newValue)
			t.Logf("got: %v", tt.this.rectangle)
			t.Logf("expect: %v", tt.want.rectangle)
		})
	}
}

func TestSubrectangleQueries_GetValue(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		this *SubrectangleQueries
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			this: &SubrectangleQueries{
				rectangle: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			args: args{
				row: 1,
				col: 1,
			},
			want: 5,
		},
		{
			name: "test case 2",
			this: &SubrectangleQueries{
				rectangle: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			args: args{
				row: 0,
				col: 1,
			},
			want: 2,
		},
		{
			name: "test case 3",
			this: &SubrectangleQueries{
				rectangle: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			args: args{
				row: 2,
				col: 2,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.this.GetValue(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("SubrectangleQueries.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
