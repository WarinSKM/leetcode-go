package finalvalueofvariableafterperformingoperations

func finalValueAfterOperations(operations []string) int {
	var result int
	operationsMap := map[string]int{
		"++X": 1,
		"X++": 1,
		"--X": -1,
		"X--": -1,
	}
	for _, operation := range operations {
		result += operationsMap[operation]
	}
	return result
}
