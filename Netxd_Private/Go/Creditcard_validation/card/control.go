package card

func Validation(cardno string) string {
	if len(cardno) < 13 && len(cardno) > 16 {
		return "Invalid card"
	}

	count := 0
	for i := 0; i < len(cardno); i++ {
		if cardno[i] < '0' || cardno[i] > '9' {
			return "Invalid card"
		}
	}

	for i := 0; i < len(cardno)-1; i++ {
		if cardno[i] == cardno[i+1] {
			count++
			if count > 3 {
				return "Invalid card"
			}
		} else {
			count = 0
		}
	}

	firstDigit := cardno[0]
	if firstDigit != '4' && firstDigit != '5' && firstDigit != '6' {
		return "Invalid card"
	}

	even := 0
	odd := 0
	for i := 0; i < len(cardno); i++ {
		if i%2 != 0 {
			digit := int(cardno[i] - '0')
			tmpe := 2 * digit
			if tmpe > 9 {
				dig := tmpe % 10
				tmpe = tmpe / 10
				even = even + (dig + tmpe)
			} else {
				even = even + tmpe
			}
		} else {
			digit := int(cardno[i] - '0')
			odd = odd + digit
		}
	}

	finalSum := even + odd

	if finalSum%10 != 0 {
		return "Invalid card"
	}

	return "Valid card"
}
