package warehouse_test

import (
	"reflect"
	"testing"

	"github.com/marios-pz/warehouse-manager/warehouse"
)

func TestCountryProducts_EmptyInput_ReturnsEmptySlice(t *testing.T) {
	prd := warehouse.Product{Barcode: "5201234567894", Country: "Ελλάδα"}

	tests := []struct {
		name     string
		products []warehouse.Product
		country  string
	}{
		{
			name:     "TC_CP_01_NilSlice",
			products: nil,
			country:  "Ελλάδα",
		},
		{
			name:     "TC_CP_02_EmptySlice",
			products: []warehouse.Product{},
			country:  "Ελλάδα",
		},
		{
			name:     "TC_CP_03_EmptyCountryQuery",
			products: []warehouse.Product{prd},
			country:  "",
		},
		{
			name:     "TC_CP_04_WhitespaceOnlyCountryQuery",
			products: []warehouse.Product{prd},
			country:  "   \t  ",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.CountryProducts(tc.products, tc.country)
			if len(got) != 0 {
				t.Errorf("CountryProducts returned %d items, want 0", len(got))
			}
		})
	}
}

func TestCountryProducts_Filtering_MatchesExpectedCriteria(t *testing.T) {
	grProduct1 := warehouse.Product{Barcode: "5201234567894", Country: "Ελλάδα"}
	grProduct2 := warehouse.Product{Barcode: "5209876543210", Country: "  Ελλάδα  "} // country με ακραία κενά
	deProduct := warehouse.Product{Barcode: "4001234567891", Country: "Γερμανία"}
	usProduct := warehouse.Product{Barcode: "0001234567895", Country: "Ηνωμένες Πολιτείες (USA)"}

	allProducts := []warehouse.Product{grProduct1, deProduct, grProduct2, usProduct}

	tests := []struct {
		name     string
		products []warehouse.Product
		country  string
		want     []warehouse.Product
	}{
		{
			name:     "TC_CP_05_ExactMatchSingle",
			products: allProducts,
			country:  "Γερμανία",
			want:     []warehouse.Product{deProduct},
		},
		{
			name:     "TC_CP_06_MultipleMatches",
			products: allProducts,
			country:  "Ελλάδα",
			want:     []warehouse.Product{grProduct1, grProduct2},
		},
		{
			name:     "TC_CP_07_CaseInsensitiveGreek",
			products: allProducts,
			country:  "ελλάδα",
			want:     []warehouse.Product{grProduct1, grProduct2},
		},
		{
			name:     "TC_CP_08_CaseInsensitiveWithParentheses",
			products: allProducts,
			country:  "ηνωμένες πολιτείες (usa)",
			want:     []warehouse.Product{usProduct},
		},
		{
			name:     "TC_CP_09_WhitespaceTrimmingInQuery",
			products: allProducts,
			country:  "   Γερμανία   \n",
			want:     []warehouse.Product{deProduct},
		},
		{
			name:     "TC_CP_10_NoMatchingCountry",
			products: allProducts,
			country:  "Ιταλία",
			want:     []warehouse.Product{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.CountryProducts(tc.products, tc.country)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CountryProducts =\n got  %v\n want %v", got, tc.want)
			}
		})
	}
}
