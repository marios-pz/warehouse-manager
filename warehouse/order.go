package warehouse

import "math"


func CalculateOrder(stock []Product, order []OrderItem) (cost float64, status string) {
	if len(order) == 0 {
		return 0.0, "Executable"
	}


	stockMap := make(map[string]Product, len(stock))
	for _, p := range stock {
		stockMap[p.Barcode] = p
	}

	status = "Executable"
	itemsCost := 0.0

	for _, item := range order {
		product, exists := stockMap[item.Barcode]

		if !exists || product.Stock < item.Quantity {
			status = "Pending"
		}

		if exists && item.Quantity > 0 {
			_, _, finalPrice := CalculateCost(product)
			itemsCost += finalPrice * float64(item.Quantity)
		}
	}

	itemsCost = math.Round(itemsCost*100) / 100

	if itemsCost == 0.0 {
		return 0.0, status
	}

	totalCost := itemsCost
	if itemsCost < 50.00 {
		totalCost += 10.00
	}

	totalCost = math.Round(totalCost*100) / 100
	return totalCost, status
}