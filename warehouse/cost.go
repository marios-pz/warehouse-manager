package warehouse

import "math"

func round2(x float64) float64 {
	return math.Round(x*100) / 100
}

func CalculateCost(p Product) (priceWithSupplyAndTaxes, discount, finalPrice float64) {
	// Kane comment out autes tis grammes an theleis na kaneis reproduce to BUG-CALC-01
	if p.Price < 0 || p.Discount < 0 || p.Discount > 100 {
		return 0, 0, 0
	}

	var vatRate float64
	switch p.Category {
	case "Τρόφιμα", "Αναψυκτικά χωρίς αλκοόλ":
		vatRate = 0.13
	case "Βιβλία":
		vatRate = 0.06
	default:
		vatRate = 0.24
	}

	priceWithSupply := round2(p.Price + p.Price*0.05)

	priceWithSupplyAndTaxes = round2(priceWithSupply + priceWithSupply*vatRate)

	discount = round2(priceWithSupplyAndTaxes * float64(p.Discount) / 100)

	finalPrice = round2(priceWithSupplyAndTaxes - discount)

	return priceWithSupplyAndTaxes, discount, finalPrice
}
