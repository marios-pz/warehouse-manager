package warehouse_test

import (
	"reflect"
	"testing"

	"github.com/marios-pz/warehouse-manager/warehouse"
)

func TestProductsToOrder_EmptyOrInvalidInput_ReturnsEmptySlice(t *testing.T) {
	sampleProduct := warehouse.Product{
		Barcode: "5201234567894",
		Country: "Ελλάδα",
		Stock:   5,
	}

	tests := []struct {
		name       string
		products   []warehouse.Product
		stockLimit int
		country    string
	}{
		{
			name:       "TC_PTO_01_NilSlice",
			products:   nil,
			stockLimit: 10,
			country:    "Ελλάδα",
		},
		{
			name:       "TC_PTO_02_EmptySlice",
			products:   []warehouse.Product{},
			stockLimit: 10,
			country:    "Ελλάδα",
		},
		{
			name:       "TC_PTO_03_EmptyCountry",
			products:   []warehouse.Product{sampleProduct},
			stockLimit: 10,
			country:    "",
		},
		{
			name:       "TC_PTO_04_WhitespaceCountry",
			products:   []warehouse.Product{sampleProduct},
			stockLimit: 10,
			country:    "   \t  ",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.ProductsToOrder(tc.products, tc.stockLimit, tc.country)
			if got == nil {
				t.Errorf("ProductsToOrder returned nil, want empty slice []Product{}")
			}
			if len(got) != 0 {
				t.Errorf("ProductsToOrder returned %d items, want 0", len(got))
			}
		})
	}
}

func TestProductsToOrder_StockBoundariesAndFiltering(t *testing.T) {
	limit := 10

	pBelow := warehouse.Product{Barcode: "5200000000018", Country: "Ελλάδα", Stock: 9}  
	pEqual := warehouse.Product{Barcode: "5200000000025", Country: "Ελλάδα", Stock: 10}
	pAbove := warehouse.Product{Barcode: "5200000000032", Country: "Ελλάδα", Stock: 11} 
	pOtherCountry := warehouse.Product{Barcode: "3001234567892", Country: "Γαλλία", Stock: 3}

	allProducts := []warehouse.Product{pBelow, pEqual, pAbove, pOtherCountry}

	tests := []struct {
		name       string
		products   []warehouse.Product
		stockLimit int
		country    string
		want       []warehouse.Product
	}{
		{
			name:       "TC_PTO_05_BoundaryBelowLimit_Included",
			products:   []warehouse.Product{pBelow},
			stockLimit: limit,
			country:    "Ελλάδα",
			want:       []warehouse.Product{pBelow},
		},
		{
			name:       "TC_PTO_06_BoundaryExactLimit_Excluded",
			products:   []warehouse.Product{pEqual},
			stockLimit: limit,
			country:    "Ελλάδα",
			want:       []warehouse.Product{},
		},
		{
			name:       "TC_PTO_07_BoundaryAboveLimit_Excluded",
			products:   []warehouse.Product{pAbove},
			stockLimit: limit,
			country:    "Ελλάδα",
			want:       []warehouse.Product{},
		},
		{
			name:       "TC_PTO_08_FilterMixedProducts_ReturnsOnlyMatching",
			products:   allProducts,
			stockLimit: limit,
			country:    "Ελλάδα",
			want:       []warehouse.Product{pBelow},
		},
		{
			name:       "TC_PTO_09_DifferentCountry_Excluded",
			products:   allProducts,
			stockLimit: limit,
			country:    "Γερμανία",
			want:       []warehouse.Product{},
		},
		{
			name:       "TC_PTO_10_CaseInsensitiveAndTrimCountry",
			products:   allProducts,
			stockLimit: limit,
			country:    "   ελλάδα   ",
			want:       []warehouse.Product{pBelow},
		},
		{
			name:       "TC_PTO_11_NegativeStockLimit_ReturnsEmpty",
			products:   allProducts,
			stockLimit: -1,
			country:    "Ελλάδα",
			want:       []warehouse.Product{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.ProductsToOrder(tc.products, tc.stockLimit, tc.country)
			if got == nil {
				t.Fatalf("ProductsToOrder returned nil slice, expected non-nil slice")
			}
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ProductsToOrder =\n got  %v\n want %v", got, tc.want)
			}
		})
	}
}