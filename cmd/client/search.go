package main

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/marios-pz/warehouse-manager/warehouse"
)

func countryTab() fyne.CanvasObject {
	country := widget.NewEntry()
	country.SetPlaceHolder("π.χ. Ελλάδα")
	results := container.NewStack()

	search := widget.NewButton("Αναζήτηση", func() {
		found := warehouse.CountryProducts(products, country.Text)
		results.Objects = []fyne.CanvasObject{productTable(found)}
		results.Refresh()
	})

	top := widget.NewForm(widget.NewFormItem("Χώρα", country))
	return container.NewBorder(container.NewVBox(top, search), nil, nil, nil, results)
}

func toOrderTab() fyne.CanvasObject {
	limit := widget.NewEntry()
	limit.SetPlaceHolder("π.χ. 10")
	country := widget.NewEntry()
	country.SetPlaceHolder("π.χ. Ελλάδα")
	message := widget.NewLabel("")
	results := container.NewStack()

	search := widget.NewButton("Εύρεση", func() {
		stockLimit, err := strconv.Atoi(strings.TrimSpace(limit.Text))
		if err != nil {
			message.SetText("Το όριο πρέπει να είναι ακέραιος αριθμός")
			return
		}
		message.SetText("")
		found := warehouse.ProductsToOrder(products, stockLimit, country.Text)
		results.Objects = []fyne.CanvasObject{productTable(found)}
		results.Refresh()
	})

	top := widget.NewForm(
		widget.NewFormItem("Ελάχιστο απόθεμα", limit),
		widget.NewFormItem("Χώρα", country),
	)
	return container.NewBorder(
		container.NewVBox(top, search, message), nil, nil, nil, results,
	)
}
