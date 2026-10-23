package warehouse_test

import (
	"testing"

	"github.com/marios-pz/warehouse-manager/warehouse"
)

func TestCalculateTotalCost_EmptyInput_ReturnsZero(t *testing.T) {
	tests := []struct {
		name     string
		products []warehouse.Product
		want     float64
	}{
		{
			name:     "TC_TOT_01_NilSlice",
			products: nil,
			want:     0.0,
		},
		{
			name:     "TC_TOT_02_EmptySlice",
			products: []warehouse.Product{},
			want:     0.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.CalculateTotalCost(tc.products)
			if got != tc.want {
				t.Errorf("CalculateTotalCost = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCalculateTotalCost_Products_ReturnsSumOfFinalPrices(t *testing.T) {
	book := warehouse.Product{Price: 10, Category: "Βιβλία", Discount: 10}       
	food := warehouse.Product{Price: 10, Category: "Τρόφιμα", Discount: 10}   
	elec := warehouse.Product{Price: 10, Category: "Ηλεκτρονικά", Discount: 10} 

	tests := []struct {
		name     string
		products []warehouse.Product
		want     float64
	}{
		{
			name:     "TC_TOT_03_SingleProduct",
			products: []warehouse.Product{book},
			want:     10.02,
		},
		{
			name:     "TC_TOT_04_MultipleProducts",
			products: []warehouse.Product{book, food},
			want:     20.70, 
		},
		{
			name:     "TC_TOT_05_SumRoundingCheck",
			products: []warehouse.Product{elec, elec, elec},
			want: 35.16,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.CalculateTotalCost(tc.products)
			if got != tc.want {
				t.Errorf("CalculateTotalCost = %v, want %v", got, tc.want)
			}
		})
	}
}
