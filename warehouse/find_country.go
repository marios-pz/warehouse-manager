package warehouse

import "strconv"

func FindCountry(barcode string) string {
	ok, _ := CheckBarcode(barcode)
	if !ok {
		return "Unknown"
	}

	prefix := barcode[:3]
	prefixNumber, _ := strconv.Atoi(prefix)

	switch {
	case prefixNumber <= 19:
		return "Ηνωμένες Πολιτείες (USA)"
	case prefixNumber >= 20 && prefixNumber <= 29:
		return "Περιορισμένη διανομή (Internal/Store use)"
	case prefixNumber >= 300 && prefixNumber <= 379:
		return "Γαλλία"
	case prefixNumber == 380:
		return "Βουλγαρία"
	case prefixNumber >= 400 && prefixNumber <= 440:
		return "Γερμανία"
	case prefixNumber == 471:
		return "Ταϊβάν"
	case prefixNumber == 520:
		return "Ελλάδα"
	case prefixNumber == 529:
		return "Κύπρος"
	case prefixNumber == 590:
		return "Πολωνία"
	case prefixNumber >= 690 && prefixNumber <= 695:
		return "Κίνα"
	case prefixNumber >= 730 && prefixNumber <= 739:
		return "Σουηδία"
	case prefixNumber >= 800 && prefixNumber <= 839:
		return "Ιταλία"
	case prefixNumber >= 840 && prefixNumber <= 849:
		return "Ισπανία"
	default:
		return "Unknown"
	}
}
