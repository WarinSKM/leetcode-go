package defanginganipaddress

func defangIPaddr(address string) string {
	var result string
	for _, char := range address {
		if char == '.' {
			result = result + "[.]"
			continue
		}
		result = result + string(char)
	}
	return result
}
