package mycalendari

type MyCalendar struct {
	booked [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{booked: [][]int{}}
}

// ["MyCalendar","book","book","book","book","book","book","book","book","book","book"]
// [[],[47,50],[33,41],[39,45],[33,42],[25,32],[26,35],[19,25],[3,8],[8,13],[18,27]]
func (myCalendar *MyCalendar) Book(start int, end int) bool {
	for _, booked := range myCalendar.booked {
		var bookedStart = booked[0]
		var bookedEnd = booked[1]

		if !(end <= bookedStart || bookedEnd <= start) {
			return false
		}
	}
	myCalendar.booked = append(myCalendar.booked, []int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
