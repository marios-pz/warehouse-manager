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

var categories = []string{
	"Ηλεκτρονικά",
	"Καλλυντικά",
	"Αλκοόλ",
	"Τρόφιμα",
	"Αναψυκτικά χωρίς αλκοόλ",
	"Βιβλία",
}

func parsePrice(s string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(strings.TrimSpace(s), ",", "."), 64)
}

func registerTab() fyne.CanvasObject {
	barcode := widget.NewEntry()
	stock := widget.NewEntry()
	category := widget.NewSelect(categories, nil)
	category.PlaceHolder = "Επιλέξτε κατηγορία"
	price := widget.NewEntry()
	discount := widget.NewEntry()
	result := widget.NewLabel("")

	form := widget.NewForm(
		widget.NewFormItem("Barcode", barcode),
		widget.NewFormItem("Απόθεμα", stock),
		widget.NewFormItem("Κατηγορία", category),
		widget.NewFormItem("Τιμή", price),
		widget.NewFormItem("Έκπτωση (%)", discount),
	)
	form.SubmitText = "Καταχώρηση"
	form.OnSubmit = func() {
		code := strings.TrimSpace(barcode.Text)
		if ok, errMsg := warehouse.CheckBarcode(code); !ok {
			result.SetText("Μη έγκυρο barcode: " + errMsg)
			return
		}
		stockValue, err := strconv.Atoi(strings.TrimSpace(stock.Text))
		if err != nil {
			result.SetText("Το απόθεμα πρέπει να είναι ακέραιος αριθμός")
			return
		}
		if category.Selected == "" {
			result.SetText("Επιλέξτε κατηγορία")
			return
		}
		priceValue, err := parsePrice(price.Text)
		if err != nil {
			result.SetText("Η τιμή πρέπει να είναι αριθμός")
			return
		}
		discountValue, err := strconv.Atoi(strings.TrimSpace(discount.Text))
		if err != nil {
			result.SetText("Η έκπτωση πρέπει να είναι ακέραιος αριθμός")
			return
		}

		country := warehouse.FindCountry(code)
		products = append(products, warehouse.Product{
			Barcode:  code,
			Country:  country,
			Stock:    stockValue,
			Category: category.Selected,
			Price:    priceValue,
			Discount: discountValue,
		})
		result.SetText(fmt.Sprintf("Καταχωρήθηκε το %s, χώρα: %s", code, country))
		barcode.SetText("")
		stock.SetText("")
		category.ClearSelected()
		price.SetText("")
		discount.SetText("")
	}

	return container.NewVBox(form, result)
}
