package minimumnumberofdecibinarynumbers

func minPartitions(n string) int {
	var result int
	for _, value := range n {
		if runeToint := int(value) - '0'; runeToint > result {
			result = runeToint
		}
	}
	return result
}
