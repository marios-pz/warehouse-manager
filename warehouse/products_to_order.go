package warehouse

func ProductsToOrder(products []Product, stockLimit int, country string) []Product {
	result := []Product{}
	for _, p := range products {
		if p.Stock < stockLimit && matchCountry(p.Country, country) {
			result = append(result, p)
		}
	}
	return result
}