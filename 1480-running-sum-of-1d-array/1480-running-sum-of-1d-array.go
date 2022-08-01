package runningsumof1darray

func runningSum(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		var tempSum int
		for j := 0; j <= i; j++ {
			tempSum += nums[j]
		}
		result = append(result, tempSum)
	}
	return result
}
