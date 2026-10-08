# warehouse-manager-go

Εργασία στο μάθημα Ποιότητα και Αξιοπιστία Λογισμικού ΧΕΙΜ-2026

## Ομάδα

- Μάριος Παπάζογλου 21390179
- Χρήστος Κρικζώνης 21390108

## Κατανομή

| Μέρος                    | Υπεύθυνος                             | Τι περιλαμβάνει                                                              |
| ------------------------ | ------------------------------------- | ---------------------------------------------------------------------------- |
| A ([#1](../../issues/1)) | Μάριος (με συμμετοχή και του Χρήστου) | `CheckBarcode`, `FindCountry`, `CalculateCost`, εφαρμογή                     |
| B ([#2](../../issues/2)) | Χρήστος                               | `CalculateTotalCost`, `CountryProducts`, `ProductsToOrder`, `CalculateOrder` |
| Τεκμηρίωση (`document/`) | Και οι δύο, παράλληλα                 | Τεχνική έκθεση σε Typst                                                      |

## Development Setup

Μία φορά μετά το clone τρέξε το pre-commit hook (gofmt, go vet, go test):

```sh
git config core.hooksPath .githooks
```
