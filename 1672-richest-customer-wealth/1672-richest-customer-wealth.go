package richestcustomerwealth

func maximumWealth(accounts [][]int) int {
	var max int
	for i := 0; i < len(accounts); i++ {
		var sum int
		for _, num := range accounts[i] {
			sum += num
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
