package mycalendari

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyCalendar
	}{
		// TODO: Add test cases.
		{
			name: "should return MyCalendar",
			want: MyCalendar{booked: [][]int{}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCalendar_Book(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name       string
		myCalendar *MyCalendar
		args       args
		want       bool
	}{
		// TODO: Add test cases.
		{
			name:       "test case 1",
			myCalendar: &MyCalendar{booked: [][]int{{1, 10}, {15, 20}}},
			args: args{
				start: 15,
				end:   20,
			},
			want: false,
		},
		{
			name:       "test case 2",
			myCalendar: &MyCalendar{booked: [][]int{{1, 10}, {15, 20}}},
			args: args{
				start: 20,
				end:   20,
			},
			want: true,
		},
		{
			name:       "test case 3",
			myCalendar: &MyCalendar{booked: [][]int{{1, 10}, {15, 20}}},
			args: args{
				start: 10,
				end:   20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.myCalendar.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MyCalendar.Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
