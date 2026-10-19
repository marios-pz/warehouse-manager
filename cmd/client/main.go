package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/marios-pz/warehouse-manager/warehouse"
)

var products = []warehouse.Product{
	{
		Barcode:  "5201234567894",
		Country:  warehouse.FindCountry("5201234567894"),
		Stock:    40,
		Category: "Τρόφιμα",
		Price:    2.50,
		Discount: 0,
	},
	{
		Barcode:  "4001234567891",
		Country:  warehouse.FindCountry("4001234567891"),
		Stock:    5,
		Category: "Ηλεκτρονικά",
		Price:    199.90,
		Discount: 10,
	},
	{
		Barcode:  "3001234567892",
		Country:  warehouse.FindCountry("3001234567892"),
		Stock:    12,
		Category: "Καλλυντικά",
		Price:    24.00,
		Discount: 20,
	},
	{
		Barcode:  "8001234567897",
		Country:  warehouse.FindCountry("8001234567897"),
		Stock:    2,
		Category: "Αλκοόλ",
		Price:    15.50,
		Discount: 0,
	},
	{
		Barcode:  "5291234567895",
		Country:  warehouse.FindCountry("5291234567895"),
		Stock:    8,
		Category: "Βιβλία",
		Price:    12.00,
		Discount: 5,
	},
}

func main() {
	a := app.New()
	w := a.NewWindow("Warehouse Manager")

	stock := container.NewTabItem("Απόθεμα", stockTab())
	price := container.NewTabItem("Ανάλυση τιμής", priceTab())
	tabs := container.NewAppTabs(
		container.NewTabItem("Καταχώρηση", registerTab()),
		stock,
		price,
		container.NewTabItem("Ανά χώρα", countryTab()),
		container.NewTabItem("Προς παραγγελία", toOrderTab()),
		container.NewTabItem("Παραγγελία", orderTab()),
	)
	tabs.OnSelected = func(t *container.TabItem) {
		switch t {
		case stock:
			t.Content = stockTab()
		case price:
			t.Content = priceTab()
		}
		tabs.Refresh()
	}

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}
