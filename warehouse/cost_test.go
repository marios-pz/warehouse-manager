package warehouse_test

import (
	"testing"

	"github.com/marios-pz/warehouse-manager/warehouse"
)

func TestCalculateCost_VatByCategory_ReturnsCorrectPrices(t *testing.T) {
	tests := []struct {
		name                                        string
		category                                    string
		wantPriceWithTaxes, wantDiscount, wantFinal float64
	}{
		{"TC_CALC_01_Books", "Βιβλία", 11.13, 1.11, 10.02},
		{"TC_CALC_02_Electronics", "Ηλεκτρονικά", 13.02, 1.30, 11.72},
		{"TC_CALC_03_Food", "Τρόφιμα", 11.87, 1.19, 10.68},
		{"TC_CALC_04_UnknownCategory", "Παιχνίδια", 13.02, 1.30, 11.72},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			product := warehouse.Product{Price: 10, Category: tc.category, Discount: 10}
			priceWithTaxes, discount, finalPrice := warehouse.CalculateCost(product)
			if priceWithTaxes != tc.wantPriceWithTaxes {
				t.Errorf("priceWithTaxes = %v, want %v", priceWithTaxes, tc.wantPriceWithTaxes)
			}
			if discount != tc.wantDiscount {
				t.Errorf("discount = %v, want %v", discount, tc.wantDiscount)
			}
			if finalPrice != tc.wantFinal {
				t.Errorf("finalPrice = %v, want %v", finalPrice, tc.wantFinal)
			}
		})
	}
}

func TestCalculateCost_DiscountBoundaries_ReturnsCorrectPrices(t *testing.T) {
	tests := []struct {
		name                                        string
		discount                                    int
		wantPriceWithTaxes, wantDiscount, wantFinal float64
	}{
		{"TC_CALC_05_ZeroDiscount", 0, 11.13, 0, 11.13},
		{"TC_CALC_06_OnePercentDiscount", 1, 11.13, 0.11, 11.02},
		{"TC_CALC_07_NinetyNinePercentDiscount", 99, 11.13, 11.02, 0.11},
		{"TC_CALC_08_FullDiscount", 100, 11.13, 11.13, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			product := warehouse.Product{Price: 10, Category: "Βιβλία", Discount: tc.discount}
			priceWithTaxes, discount, finalPrice := warehouse.CalculateCost(product)
			if priceWithTaxes != tc.wantPriceWithTaxes {
				t.Errorf("priceWithTaxes = %v, want %v", priceWithTaxes, tc.wantPriceWithTaxes)
			}
			if discount != tc.wantDiscount {
				t.Errorf("discount = %v, want %v", discount, tc.wantDiscount)
			}
			if finalPrice != tc.wantFinal {
				t.Errorf("finalPrice = %v, want %v", finalPrice, tc.wantFinal)
			}
		})
	}
}

func TestCalculateCost_InvalidInput_ReturnsZero(t *testing.T) {
	tests := []struct {
		name     string
		price    float64
		discount int
	}{
		{"TC_CALC_09_NegativeDiscount", 10, -1},
		{"TC_CALC_10_DiscountAboveHundred", 10, 101},
		{"TC_CALC_11_ZeroPrice", 0, 10},
		{"TC_CALC_12_NegativePrice", -1, 10},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			product := warehouse.Product{Price: tc.price, Category: "Βιβλία", Discount: tc.discount}
			priceWithTaxes, discount, finalPrice := warehouse.CalculateCost(product)
			if priceWithTaxes != 0 || discount != 0 || finalPrice != 0 {
				t.Errorf("CalculateCost = (%v, %v, %v), want (0, 0, 0)", priceWithTaxes, discount, finalPrice)
			}
		})
	}
}
