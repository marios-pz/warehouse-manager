// Package warehouse is the product management library (ManageProductsLib).
package warehouse

type Product struct {
	Barcode  string
	Country  string
	Stock    int
	Category string
	Price    float64
	Discount int // percent, 0-100
}

type OrderItem struct {
	Barcode  string
	Quantity int
}
