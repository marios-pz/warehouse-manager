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
		result.SetText(fmt.Sprintf("Κατάσταση: %s\nΚόστος: %.2f €", status, cost))
	})

	return container.NewVBox(lines, submit, result)
}
