package warehouse_test

import (
	"testing"

	"github.com/marios-pz/warehouse-manager/warehouse"
)

func TestCheckBarcode_ValidBarcode_ReturnsTrue(t *testing.T) {
	tests := []struct {
		name    string
		barcode string
	}{
		{"TC_BC_01_Valid", "4006381333931"},
		{"TC_BC_02_ValidCheckDigitZero", "5201234567870"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok, errMsg := warehouse.CheckBarcode(tc.barcode)
			if !ok {
				t.Errorf("CheckBarcode(%q) = %v, want true", tc.barcode, ok)
			}
			if errMsg != "" {
				t.Errorf("CheckBarcode(%q) errMsg = %q, want empty", tc.barcode, errMsg)
			}
		})
	}
}

func TestCheckBarcode_InvalidBarcode_ReturnsFalse(t *testing.T) {
	tests := []struct {
		name    string
		barcode string
	}{
		{"TC_BC_03_WrongCheckDigit", "4006381333935"},
		{"TC_BC_04_Empty", ""},
		{"TC_BC_05_TwelveDigits", "400638133393"},
		{"TC_BC_06_FourteenDigits", "40063813339310"},
		{"TC_BC_07_NonNumeric", "40063813339A1"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok, errMsg := warehouse.CheckBarcode(tc.barcode)
			if ok {
				t.Errorf("CheckBarcode(%q) = %v, want false", tc.barcode, ok)
			}
			if errMsg == "" {
				t.Errorf("CheckBarcode(%q) errMsg is empty, want a message", tc.barcode)
			}
		})
	}
}
