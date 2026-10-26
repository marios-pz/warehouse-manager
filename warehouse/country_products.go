package warehouse

import "strings"

func matchCountry(productCountry, targetCountry string) bool {
	target := strings.TrimSpace(targetCountry)
	if target == "" {
		return false
	}
	return strings.EqualFold(strings.TrimSpace(productCountry), target)
}

func CountryProducts(products []Product, country string) []Product {
	matched := []Product{}
	if len(products) == 0 {
		return matched
	}

	for _, p := range products {
		if matchCountry(p.Country, country) {
			matched = append(matched, p)
		}
	}

	return matched
}
