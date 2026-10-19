package warehouse_test

import (
	"testing"

	"github.com/marios-pz/warehouse-manager/warehouse"
)

func TestFindCountry_KnownPrefix_ReturnsCountry(t *testing.T) {
	tests := []struct {
		name    string
		barcode string
		country string
	}{
		{"TC_FC_01_Greece", "5201234567894", "Ελλάδα"},
		{"TC_FC_02_France", "3001234567892", "Γαλλία"},
		{"TC_FC_03_Italy", "8391234567899", "Ιταλία"},
		{"TC_FC_04_Spain", "8491234567896", "Ισπανία"},
		{"TC_FC_07_UsaLower", "0001234567895", "Ηνωμένες Πολιτείες (USA)"},
		{"TC_FC_08_UsaUpper", "0191234567893", "Ηνωμένες Πολιτείες (USA)"},
		{"TC_FC_09_InternalLower", "0201234567899", "Περιορισμένη διανομή (Internal/Store use)"},
		{"TC_FC_10_InternalUpper", "0291234567890", "Περιορισμένη διανομή (Internal/Store use)"},
		{"TC_FC_11_FranceUpper", "3791234567892", "Γαλλία"},
		{"TC_FC_12_Bulgaria", "3801234567898", "Βουλγαρία"},
		{"TC_FC_13_GermanyLower", "4001234567891", "Γερμανία"},
		{"TC_FC_14_GermanyUpper", "4401234567899", "Γερμανία"},
		{"TC_FC_15_Taiwan", "4711234567899", "Ταϊβάν"},
		{"TC_FC_16_Cyprus", "5291234567895", "Κύπρος"},
		{"TC_FC_17_Poland", "5901234567893", "Πολωνία"},
		{"TC_FC_18_ChinaLower", "6901234567892", "Κίνα"},
		{"TC_FC_19_ChinaUpper", "6951234567897", "Κίνα"},
		{"TC_FC_20_SwedenLower", "7301234567899", "Σουηδία"},
		{"TC_FC_21_SwedenUpper", "7391234567890", "Σουηδία"},
		{"TC_FC_22_ItalyLower", "8001234567897", "Ιταλία"},
		{"TC_FC_23_SpainLower", "8401234567895", "Ισπανία"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.FindCountry(tc.barcode)
			if got != tc.country {
				t.Errorf("FindCountry(%q) = %q, want %q", tc.barcode, got, tc.country)
			}
		})
	}
}

func TestFindCountry_InvalidBarcode_ReturnsUnknown(t *testing.T) {
	tests := []struct {
		name    string
		barcode string
		country string
	}{
		{"TC_FC_05_WrongCheckDigit", "4006381333935", "Unknown"},
		{"TC_FC_06_Empty", "", "Unknown"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.FindCountry(tc.barcode)
			if got != tc.country {
				t.Errorf("FindCountry(%q) = %q, want %q", tc.barcode, got, tc.country)
			}
		})
	}
}

func TestFindCountry_UnknownPrefix_ReturnsUnknown(t *testing.T) {
	tests := []struct {
		name    string
		barcode string
		country string
	}{
		{"TC_FC_24_AfterInternal", "0301234567896", "Unknown"},
		{"TC_FC_25_AfterBulgaria", "3811234567897", "Unknown"},
		{"TC_FC_26_BeforeGermany", "3991234567896", "Unknown"},
		{"TC_FC_27_AfterGermany", "4411234567898", "Unknown"},
		{"TC_FC_28_BeforeChina", "6891234567896", "Unknown"},
		{"TC_FC_29_AfterChina", "6961234567896", "Unknown"},
		{"TC_FC_30_AfterSpain", "8501234567892", "Unknown"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := warehouse.FindCountry(tc.barcode)
			if got != tc.country {
				t.Errorf("FindCountry(%q) = %q, want %q", tc.barcode, got, tc.country)
			}
		})
	}
}
