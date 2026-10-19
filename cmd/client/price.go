package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/marios-pz/warehouse-manager/warehouse"
)

func priceTab() fyne.CanvasObject {
	details := widget.NewLabel("Επιλέξτε προϊόν")

	barcodes := make([]string, len(products))
	for i, p := range products {
		barcodes[i] = p.Barcode
	}

	choose := widget.NewSelect(barcodes, func(code string) {
		for _, p := range products {
			if p.Barcode != code {
				continue
			}
			withTaxes, discount, finalPrice := warehouse.CalculateCost(p)
			details.SetText(fmt.Sprintf(
				"Κατηγορία: %s\n"+
					"Αρχική τιμή: %.2f €\n"+
					"Τιμή με προμήθεια και ΦΠΑ: %.2f €\n"+
					"Έκπτωση (%d%%): %.2f €\n"+
					"Τελική τιμή: %.2f €",
				p.Category, p.Price, withTaxes, p.Discount, discount, finalPrice,
			))
			return
		}
	})
	choose.PlaceHolder = "Barcode προϊόντος"

	return container.NewVBox(choose, details)
}
