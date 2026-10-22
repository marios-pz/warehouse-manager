package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/marios-pz/warehouse-manager/warehouse"
)

func parseOrder(text string) ([]warehouse.OrderItem, error) {
	var items []warehouse.OrderItem
	for i, line := range strings.Split(text, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if len(fields) != 2 {
			return nil, fmt.Errorf("γραμμή %d: χρειάζεται barcode και ποσότητα", i+1)
		}
		quantity, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("γραμμή %d: η ποσότητα πρέπει να είναι ακέραιος", i+1)
		}
		items = append(items, warehouse.OrderItem{Barcode: fields[0], Quantity: quantity})
	}
	return items, nil
}

func missingBarcodes(items []warehouse.OrderItem) []string {
	known := make(map[string]bool, len(products))
	for _, p := range products {
		known[p.Barcode] = true
	}
	var missing []string
	for _, item := range items {
		if !known[item.Barcode] {
			missing = append(missing, item.Barcode)
		}
	}
	return missing
}

func orderTab() fyne.CanvasObject {
	lines := widget.NewMultiLineEntry()
	lines.SetPlaceHolder("Μία γραμμή ανά προϊόν: barcode ποσότητα\nπ.χ. 5201234567894 3")
	lines.SetMinRowsVisible(6)
	result := widget.NewLabel("")

	submit := widget.NewButton("Υπολογισμός παραγγελίας", func() {
		items, err := parseOrder(lines.Text)
		if err != nil {
			result.SetText(err.Error())
			return
		}
		cost, status := warehouse.CalculateOrder(products, items)
		text := fmt.Sprintf("Κατάσταση: %s\nΚόστος: %.2f €", status, cost)
		if missing := missingBarcodes(items); len(missing) > 0 {
			text += "\nΔεν υπάρχουν στην αποθήκη: " + strings.Join(missing, ", ")
		}
		result.SetText(text)
	})

	return container.NewVBox(lines, submit, result)
}
