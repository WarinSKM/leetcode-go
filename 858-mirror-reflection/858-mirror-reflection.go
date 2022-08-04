package mirrorreflection

func mirrorReflection(p int, q int) int {
	leftRight := true // true = right; false = left
	upDown := true    // true = up; false = down
	currentPosition := q
	for currentPosition != 0 {
		if currentPosition == p {
			break
		}
		currentPosition += q
		leftRight = !leftRight
		if currentPosition >= p {
			currentPosition = currentPosition % p
			if currentPosition == 0 {
				break
			}
			upDown = !upDown
		}
	}

	if leftRight && upDown {
		return 1
	} else if !leftRight && upDown {
		return 2
	}
	return 0
}
