package buildarrayfrompermutation

func buildArray(nums []int) []int {
	var result []int
	for _, val := range nums {
		result = append(result, nums[val])
	}
	return result
}
