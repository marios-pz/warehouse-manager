package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/marios-pz/warehouse-manager/warehouse"
)

var columns = []struct {
	title string
	width float32
}{
	{"Barcode", 140},
	{"Χώρα", 190},
	{"Κατηγορία", 190},
	{"Απόθεμα", 80},
	{"Τιμή", 80},
	{"Έκπτωση", 80},
	{"Τελική τιμή", 100},
}

func cellText(p warehouse.Product, col int) string {
	switch col {
	case 0:
		return p.Barcode
	case 1:
		return p.Country
	case 2:
		return p.Category
	case 3:
		return strconv.Itoa(p.Stock)
	case 4:
		return fmt.Sprintf("%.2f €", p.Price)
	case 5:
		return fmt.Sprintf("%d%%", p.Discount)
	default:
		_, _, finalPrice := warehouse.CalculateCost(p)
		return fmt.Sprintf("%.2f €", finalPrice)
	}
}

func productTable(list []warehouse.Product) fyne.CanvasObject {
	table := widget.NewTable(
		func() (int, int) { return len(list), len(columns) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(cellText(list[id.Row], id.Col))
		},
	)
	table.ShowHeaderRow = true
	table.CreateHeader = func() fyne.CanvasObject { return widget.NewLabel("") }
	table.UpdateHeader = func(id widget.TableCellID, o fyne.CanvasObject) {
		label := o.(*widget.Label)
		label.TextStyle.Bold = true
		label.SetText(columns[id.Col].title)
	}
	for i, c := range columns {
		table.SetColumnWidth(i, c.width)
	}
	return table
}

func stockTab() fyne.CanvasObject {
	total := widget.NewLabel(fmt.Sprintf(
		"Προϊόντα: %d   Συνολικό κόστος: %.2f €",
		len(products), warehouse.CalculateTotalCost(products),
	))
	return container.NewBorder(nil, total, nil, nil, productTable(products))
}
