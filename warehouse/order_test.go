package warehouse_test

import (
	"testing"

	"github.com/marios-pz/warehouse-manager/warehouse"
)

func TestCalculateOrder_EmptyOrder_ReturnsZeroAndExecutable(t *testing.T) {
	stock := []warehouse.Product{
		{Barcode: "5201234567894", Stock: 10, Price: 10, Category: "Βιβλία"},
	}

	cost, status := warehouse.CalculateOrder(stock, nil)
	if cost != 0.0 || status != "Executable" {
		t.Errorf("CalculateOrder(nil) = (%v, %q), want (0.0, \"Executable\")", cost, status)
	}

	cost, status = warehouse.CalculateOrder(stock, []warehouse.OrderItem{})
	if cost != 0.0 || status != "Executable" {
		t.Errorf("CalculateOrder([]) = (%v, %q), want (0.0, \"Executable\")", cost, status)
	}
}

func TestCalculateOrder_StockAvailabilityAndStatus(t *testing.T) {
	pAvailable := warehouse.Product{Barcode: "5200000000001", Stock: 10, Price: 10, Category: "Βιβλία", Discount: 0}
	pExact := warehouse.Product{Barcode: "5200000000018", Stock: 5, Price: 10, Category: "Βιβλία", Discount: 0}
	pShort := warehouse.Product{Barcode: "5200000000025", Stock: 4, Price: 10, Category: "Βιβλία", Discount: 0}
	stock := []warehouse.Product{pAvailable, pExact, pShort}

	tests := []struct {
		name       string
		order      []warehouse.OrderItem
		wantStatus string
	}{
		{
			name: "TC_ORD_01_AllAvailable_Executable",
			order: []warehouse.OrderItem{
				{Barcode: "5200000000001", Quantity: 5},
			},
			wantStatus: "Executable",
		},
		{
			name: "TC_ORD_02_ExactStock_Executable",
			order: []warehouse.OrderItem{
				{Barcode: "5200000000018", Quantity: 5},
			},
			wantStatus: "Executable",
		},
		{
			name: "TC_ORD_03_StockShortByOne_Pending",
			order: []warehouse.OrderItem{
				{Barcode: "5200000000025", Quantity: 5},
			},
			wantStatus: "Pending",
		},
		{
			name: "TC_ORD_04_NonExistentBarcode_Pending",
			order: []warehouse.OrderItem{
				{Barcode: "9999999999999", Quantity: 1},
			},
			wantStatus: "Pending",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, gotStatus := warehouse.CalculateOrder(stock, tc.order)
			if gotStatus != tc.wantStatus {
				t.Errorf("CalculateOrder() status = %q, want %q", gotStatus, tc.wantStatus)
			}
		})
	}
}

func TestCalculateOrder_ShippingFeeBoundaries(t *testing.T) {
	p4999 := warehouse.Product{Barcode: "5200000000117", Stock: 10, Price: 44.91, Category: "Βιβλία", Discount: 0}
	p5000 := warehouse.Product{Barcode: "5200000000124", Stock: 10, Price: 44.92, Category: "Βιβλία", Discount: 0}
	p5001 := warehouse.Product{Barcode: "5200000000131", Stock: 10, Price: 44.93, Category: "Βιβλία", Discount: 0}

	stock := []warehouse.Product{p4999, p5000, p5001}

	tests := []struct {
		name     string
		order    []warehouse.OrderItem
		wantCost float64
	}{
		{
			name: "TC_ORD_05_BelowFifty_AddsTenShipping",
			order: []warehouse.OrderItem{
				{Barcode: "5200000000117", Quantity: 1},
			},
			wantCost: 59.99,
		},
		{
			name: "TC_ORD_06_ExactFifty_FreeShipping",
			order: []warehouse.OrderItem{
				{Barcode: "5200000000124", Quantity: 1},
			},
			wantCost: 50.00,
		},
		{
			name: "TC_ORD_07_AboveFifty_FreeShipping",
			order: []warehouse.OrderItem{
				{Barcode: "5200000000131", Quantity: 1},
			},
			wantCost: 50.01,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotCost, _ := warehouse.CalculateOrder(stock, tc.order)
			if gotCost != tc.wantCost {
				t.Errorf("CalculateOrder cost = %v, want %v", gotCost, tc.wantCost)
			}
		})
	}
}
