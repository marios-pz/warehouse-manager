package warehouse

import "regexp"

var digitsOnly = regexp.MustCompile("^[0-9]+$")

func CheckBarcode(barcode string) (ok bool, errMsg string) {
	if barcode == "" {
		return false, "Το barcode είναι κενό"
	}

	if len(barcode) != 13 {
		return false, "Το barcode πρέπει να έχει 13 ψηφία"
	}

	if !digitsOnly.MatchString(barcode) {
		return false, "Το barcode πρέπει να περιέχει μόνο ψηφία"
	}

	sum := 0
	for i := 0; i < 12; i++ {
		digit := int(barcode[i] - '0')
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}

	checkDigit := (10 - sum%10) % 10

	if checkDigit != int(barcode[12]-'0') {
		return false, "Λάθος check digit"
	}

	return true, ""
}
