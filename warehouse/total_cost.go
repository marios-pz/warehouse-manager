package warehouse

func CalculateTotalCost(products []Product) float64 {
	totalCost := 0.0
	for i := range products {
		_, _, finalPrice := CalculateCost(products[i])
		totalCost += finalPrice
	}
	return round2(totalCost)
}
