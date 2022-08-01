package strongpasswordchecker

func strongPasswordChecker(password string) int {
	lengthCase := 0
	countPassword := len(password)

	// character case
	uppercase := 0
	lowercase := 0
	digi := 0
	charCase := 0

	// repeat case
	repeat := 0

	// length case
	if countPassword < 6 {
		lengthCase = 6 - countPassword
	} else if countPassword > 20 {
		lengthCase = countPassword - 20
	}

	for i := 0; i < countPassword; i++ {
		if i+2 > 20 {
			break
		}

		// repeat case
		if i+2 < countPassword && i+2 <= 20 {
			if password[i] == password[i+1] && password[i] == password[i+2] {
				repeat += 1
				i += 2
			}
		}

		if password[i] >= 'a' && password[i] <= 'z' {
			lowercase += 1
		}
		if password[i] >= '0' && password[i] <= '9' {
			digi += 1
		}
		if password[i] >= 'A' && password[i] <= 'Z' {
			uppercase += 1
		}
	}

	// final check charcase
	if lowercase == 0 {
		charCase += 1
	}
	if digi == 0 {
		charCase += 1
	}
	if uppercase == 0 {
		charCase += 1
	}

	if charCase > 0 && lengthCase > 0 {
		if charCase > lengthCase {
			lengthCase = 0
		} else {
			charCase = 0
		}
	}

	if repeat > 0 && charCase > 0 {
		if charCase > repeat {
			repeat = 0
		} else {
			charCase = 0
		}
	}
	if repeat > 0 && lengthCase > 0 {
		if lengthCase > repeat {
			repeat = 0
		} else {
			lengthCase = 0
		}
	}

	return lengthCase + charCase + repeat
}
