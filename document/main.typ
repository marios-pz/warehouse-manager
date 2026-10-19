#import "@preview/fletcher:0.5.8": diagram, node, edge

#set document(title: "Warehouse Manager: Τεχνική Έκθεση")
#set text(lang: "el", size: 11pt)
#set par(justify: true)
#set heading(numbering: "1.1")
#set page(
  paper: "a4",
  margin: (top: 3cm, x: 2.5cm, bottom: 2.5cm),
  header: image("assets/uniwa-logo.png", height: 1.5cm),
  numbering: "1",
)
#show table.cell.where(y: 0): strong
#set table(inset: 6pt, align: left + horizon)
#show table: set par(justify: false)
#show figure.where(kind: table): set figure.caption(position: top)
#show figure.where(kind: table): set block(breakable: true)

#let todo(body) = text(fill: red)[*TODO:* #body]

#let ep-table(caption, ..rows) = figure(
  table(
    columns: (auto, 1fr, 1fr, 1fr),
    table.header(
      [Παράμετρος Εισόδου],
      [Έγκυρες Κλάσεις Ισοδυναμίας (Valid Classes)],
      [Άκυρες Κλάσεις Ισοδυναμίας (Invalid Classes)],
      [Οριακές Τιμές (Boundaries)],
    ),
    ..rows,
  ),
  caption: caption,
)

#let tc-table(caption, ..rows) = figure(
  table(
    columns: (auto, 1fr, 1fr, 1fr),
    table.header(
      [Test Case ID],
      [Στόχος / Περιγραφή Ελέγχου],
      [Δεδομένα Εισόδου],
      [Αναμενόμενο Αποτέλεσμα],
    ),
    ..rows,
  ),
  caption: caption,
)

#let shot(path, caption) = figure(
  if path == none {
    rect(width: 100%, height: 4cm, stroke: (paint: red, dash: "dashed"))[
      #align(center + horizon)[#todo[screenshot]]
    ]
  } else { image(path, width: 100%) },
  caption: caption,
)

#let bug-report(
  id: none,
  test: none,
  purpose: none,
  failure: none,
  expected: none,
  actual: none,
  red: none,
  app: none,
  green: none,
) = {
  let v(x) = if x == none { todo[] } else { x }
  figure(
    table(
      columns: (auto, 1fr),
      table.header[Πεδίο][Τιμή],
      [Αναγνωριστικό Σφάλματος (Bug ID)], v(id),
      [Όνομα Μεθόδου Ελέγχου], v(test),
      [Σκοπός του Ελέγχου], v(purpose),
      [Μήνυμα Αποτυχίας], v(failure),
      [Αναμενόμενη Τιμή], v(expected),
      [Πραγματική Τιμή], v(actual),
    ),
    caption: [Αναφορά σφάλματος #v(id)],
  )
  shot(red, [Red Phase: η δοκιμή αποτυγχάνει (#v(id))])
  shot(app, [App Behavior: λανθασμένη συμπεριφορά στην εφαρμογή (#v(id))])
  shot(green, [Green Phase: όλες οι δοκιμές περνούν (#v(id))])
}

#align(center)[
  Πανεπιστήμιο Δυτικής Αττικής \
  Πολυτεχνική Σχολή \
  Τμήμα Μηχανικών Πληροφορικής και Υπολογιστών

  #v(1cm)

  #text(size: 18pt, weight: "bold")[Ποιότητα και Αξιοπιστία Λογισμικού]

  Ακαδημαϊκό Έτος 2026-2027

  #v(1cm)

  #text(size: 14pt, weight: "bold")[Απαλλακτική Εργασία]

  #text(size: 13pt)[
    Ανάπτυξη Εφαρμογής Διαχείρισης Αποθήκης, Σχεδιασμός Περιπτώσεων
    Ελέγχου, Unit Testing & Αναφορές Σφαλμάτων
  ]

  #v(1cm)

  Team PK

  #table(
    columns: 2,
    align: left,
    table.header[Ονοματεπώνυμο][Αριθμός Μητρώου],
    [Μάριος Παπάζογλου], [21390179],
    [Χρήστος Κρικζώνης], [21390108],
  )
]

#pagebreak()
#outline(depth: 2)
#pagebreak()

= Εισαγωγή

Στην εργασία αναπτύξαμε μια βιβλιοθήκη διαχείρισης προϊόντων αποθήκης
(ManageProductsLib), μια εφαρμογή πελάτη που τη χρησιμοποιεί
(ProductClientApp) και ένα σύνολο αυτοματοποιημένων ελέγχων μονάδας
(ManageProductsLib.Tests). Για κάθε μία από τις 7 μεθόδους της βιβλιοθήκης
παρουσιάζουμε τον πίνακα κλάσεων ισοδυναμίας και οριακών τιμών, τον πίνακα
περιπτώσεων ελέγχου και τουλάχιστον μία αναφορά σφάλματος με τα αντίστοιχα
στιγμιότυπα οθόνης.

= Αρχιτεκτονική

== Τεχνολογίες

Υλοποιήσαμε την εργασία σε Go αντί για C\#/.NET, με τη σύμφωνη γνώμη του
διδάσκοντα. Η βιβλιοθήκη `ManageProductsLib.dll` αντιστοιχεί στο package
`warehouse`, το project ελέγχων `ManageProductsLib.Tests` στο package
`warehouse_test` (black-box tests που βλέπουν μόνο το δημόσιο API) και το
Test Explorer στην έξοδο της εντολής `go test -v`, όπου η κόκκινη και η
πράσινη φάση εμφανίζονται ως `FAIL` και `PASS` αντίστοιχα. Την κάλυψη κώδικα
τη μετράμε με `go test -cover`.

#figure(
  table(
    columns: (1fr, 1fr),
    table.header[Εκφώνηση (.NET)][Υλοποίηση (Go)],
    [Visual Studio Solution], [Go module (`go.mod`)],
    [Class Library `ManageProductsLib.dll`], [package `warehouse`],
    [Test Project `ManageProductsLib.Tests`],
    [package `warehouse_test`, πακέτο `testing`],

    [MSTest / NUnit / xUnit], [`go test`],
    [Test Explorer (κόκκινο / πράσινο)], [`go test -v` (`FAIL` / `PASS`)],
    [Stack Trace αποτυχίας], [αρχείο και γραμμή της αποτυχημένης assertion],
    [Code Coverage], [`go test -cover`, `go tool cover -html`],
    [Παράμετροι `ref`], [πολλαπλές τιμές επιστροφής],
    [`null`], [`nil` slice για πίνακες, `""` για strings],
    [`double`], [`float64`],
    [`class`], [`struct` και συναρτήσεις σε επίπεδο package],
    [Desktop UI (WinForms/WPF/MAUI)], [Desktop UI με Fyne],
  ),
  caption: [Αντιστοίχιση εκφώνησης και υλοποίησης],
)

== Δομή Λύσης (Solution Structure)

Ακολουθήσαμε τις συμβάσεις της Go: ένας φάκελος ανά package με όνομα ίδιο με
το package, και οι έλεγχοι σε αρχεία `*_test.go` δίπλα στον κώδικα που
ελέγχουν. Οι έλεγχοι δηλώνουν `package warehouse_test`, οπότε
μεταγλωττίζονται ως ξεχωριστό package και βλέπουν μόνο το δημόσιο API της
βιβλιοθήκης, όπως ένα ξεχωριστό Test Project.

#figure(
  table(
    columns: (1fr, 1fr),
    table.header[Project εκφώνησης][Φάκελος / package],
    [ManageProductsLib], [`warehouse/` (package `warehouse`)],
    [ManageProductsLib.Tests],
    [`warehouse/*_test.go` (package `warehouse_test`)],

    [ProductClientApp], [`cmd/client/` (package `main`)],
  ),
  caption: [Αντιστοίχιση projects και φακέλων],
)

```
warehouse-manager/
├── go.mod
├── warehouse/                  ManageProductsLib
│   ├── product.go              Product, OrderItem
│   ├── barcode.go              CheckBarcode
│   ├── find_country.go         FindCountry
│   ├── cost.go                 CalculateCost
│   ├── total_cost.go           CalculateTotalCost
│   ├── country_products.go     CountryProducts
│   ├── products_to_order.go    ProductsToOrder
│   ├── order.go                CalculateOrder
│   └── *_test.go               ManageProductsLib.Tests
└── cmd/client/                 ProductClientApp (Fyne GUI)
```

== Διάγραμμα Κλάσεων (Class Diagram)

Η Go δεν έχει κλάσεις. Τα δεδομένα ορίζονται ως structs και η λειτουργικότητα
ως συναρτήσεις σε επίπεδο package. Στο διάγραμμα σχεδιάσαμε κάθε struct ως
κλάση με πεδία και χωρίς μεθόδους, και το package `warehouse` ως κλάση με
στερεότυπο «package» που περιέχει τις 7 συναρτήσεις, το αντίστοιχο μιας
static κλάσης στη C\#. Τα διακεκομμένα βέλη δείχνουν εξαρτήσεις: η εφαρμογή
και οι έλεγχοι χρησιμοποιούν τη βιβλιοθήκη, και οι συναρτήσεις της
βιβλιοθήκης δέχονται `Product` και `OrderItem`.

#let mono(body) = text(size: 7.5pt, font: "DejaVu Sans Mono", body)
#let uml(name, fields, methods: none) = table(
  columns: 1,
  inset: 5pt,
  align: left,
  fill: (_, y) => if y == 0 { luma(235) },
  table.cell(align: center)[#name],
  mono(fields),
  ..if methods != none { (mono(methods),) },
)

#figure(
  diagram(
    spacing: (0cm, 1.2cm),
    edge-stroke: 0.6pt,
    node((0, 0), uml([«package main» \ client], [Fyne GUI])),
    node((2, 0), uml([«package» \ warehouse_test], [unit tests])),
    node(
      (1, 1),
      uml(
        [«package» \ warehouse],
        [ManageProductsLib],
        methods: [
          \+ CheckBarcode(barcode string) (bool, string) \
          \+ FindCountry(barcode string) string \
          \+ CalculateCost(p Product) (float64, float64, float64) \
          \+ CalculateTotalCost(products []Product) float64 \
          \+ CountryProducts(products []Product, country string) \
          #h(1em) []Product \
          \+ ProductsToOrder(products []Product, stockLimit int, \
          #h(1em) country string) []Product \
          \+ CalculateOrder(stock []Product, order []OrderItem) \
          #h(1em) (float64, string)
        ],
      ),
    ),
    node(
      (0, 2),
      uml([Product], [
        \+ Barcode: string \
        \+ Country: string \
        \+ Stock: int \
        \+ Category: string \
        \+ Price: float64 \
        \+ Discount: int
      ]),
    ),
    node(
      (2, 2),
      uml([OrderItem], [
        \+ Barcode: string \
        \+ Quantity: int
      ]),
    ),
    edge((0, 0), (1, 1), "-->", [«use»], label-side: right),
    edge((2, 0), (1, 1), "-->", [«use»], label-side: left),
    edge((1, 1), (0, 2), "-->", [«use»], label-side: right),
    edge((1, 1), (2, 2), "-->", [«use»], label-side: left),
  ),
  caption: [Διάγραμμα κλάσεων],
)

== Μέθοδοι της Βιβλιοθήκης

#figure(
  table(
    columns: (1fr, 1fr),
    table.header[Υπογραφή εκφώνησης][Υπογραφή Go],
    [`bool CheckBarcode(string Barcode, ref string ErrorMsg)`],
    [`CheckBarcode(barcode string) (ok bool, errMsg string)`],

    [`string FindCountry(string Barcode)`],
    [`FindCountry(barcode string) string`],

    raw(
      "void CalculateCost(Product Prd, ref double PriceWithSupplyandTaxes, "
        + "ref double Discount, ref double FinalPrice)",
    ),
    raw(
      "CalculateCost(p Product) "
        + "(priceWithSupplyAndTaxes, discount, finalPrice float64)",
    ),

    [`double CalculateTotalCost(Product[] Prds)`],
    [`CalculateTotalCost(products []Product) float64`],

    [`Product[] CountryProducts(Product[] Prds, string Country)`],
    [`CountryProducts(products []Product, country string) []Product`],

    raw(
      "ProductsToOrder(Product[] Prds, int StockLimit, string Country, "
        + "ref Product[] PrdsToOrder)",
    ),
    raw(
      "ProductsToOrder(products []Product, stockLimit int, "
        + "country string) []Product",
    ),

    raw(
      "double CalculateOrder(Product[] StockProducts, "
        + "OrderItem[] OrderProducts, ref string OrderStatus)",
    ),
    raw(
      "CalculateOrder(stock []Product, order []OrderItem) "
        + "(cost float64, status string)",
    ),
  ),
  caption: [Υπογραφές μεθόδων],
)

= Παραδοχές

Τα παρακάτω σημεία δεν ορίζονται ρητά από την εκφώνηση. Καταγράφουμε εδώ τις
αποφάσεις μας, στις οποίες βασίζονται και οι έλεγχοι.

#figure(
  table(
    columns: (auto, 1fr, 1fr),
    table.header[\#][Ερώτημα][Απόφαση],
    [Π1],
    [
      `CalculateCost`: η έκπτωση υπολογίζεται πάνω στην τιμή με προμήθεια
      και ΦΠΑ;
    ],
    [
      Ναι. Ακολουθούμε τη σειρά των βημάτων της εκφώνησης: προμήθεια, ΦΠΑ,
      ποσό έκπτωσης πάνω στην τιμή με ΦΠΑ, τελική τιμή.
    ],

    [Π2],
    [
      `CalculateCost`: τι επιστρέφεται για Discount εκτός [0, 100] ή
      αρνητική τιμή;
    ],
    [
      Επιστρέφουμε (0, 0, 0). Η υπογραφή δεν έχει τρόπο να επιστρέψει
      σφάλμα, και η εφαρμογή απορρίπτει τέτοιες τιμές κατά την καταχώρηση.
    ],

    [Π3],
    [
      `CalculateTotalCost`: άθροισμα τελικών τιμών ανά μονάδα ή επί
      `Stock`;
    ],
    todo[],

    [Π4],
    [`CalculateOrder`: το «κόστος με ΦΠΑ» είναι `Quantity × FinalPrice`;],
    todo[],

    [Π5],
    [
      `CalculateOrder`: σε `"Pending"`, τα προϊόντα που λείπουν μετράνε
      στο κόστος;
    ],
    todo[],

    [Π6],
    [
      `ProductsToOrder`: case-insensitive σύγκριση και trim στη χώρα, όπως
      στην `CountryProducts`;
    ],
    todo[],

    [Π7],
    [
      `FindCountry`: σε ποια μορφή επιστρέφεται το όνομα της χώρας
      (π.χ. «Ελλάδα»);
    ],
    todo[],
  ),
  caption: [Παραδοχές],
)

= Σχεδιασμός Δοκιμών

Για κάθε μέθοδο εφαρμόσαμε Ισοδύναμη Διαμέριση (Equivalence Partitioning)
και Ανάλυση Οριακών Τιμών (Boundary Value Analysis). Κάθε περίπτωση ελέγχου
αντιστοιχεί σε μία μέθοδο ελέγχου στα αρχεία `warehouse/*_test.go`.

== CheckBarcode

Το `CheckBarcode` κάνει τέσσερις ελέγχους με τη σειρά: αν το barcode είναι
κενό, αν έχει ακριβώς 13 χαρακτήρες, αν όλοι είναι ψηφία και αν το check
digit του Παραρτήματος Α ταιριάζει με το 13ο ψηφίο. Κάθε έλεγχος αντιστοιχεί
σε μία άκυρη κλάση. Στο μήκος ελέγξαμε τις οριακές τιμές 12, 13 και 14.
Στο check digit ελέγξαμε και την περίπτωση όπου το άθροισμα είναι
πολλαπλάσιο του 10, γιατί μόνο τότε το check digit γίνεται 0.

#ep-table(
  [Κλάσεις ισοδυναμίας και οριακές τιμές: CheckBarcode],
  [Barcode (string)],
  [13 ψηφία 0-9 με σωστό check digit],
  [
    Κενό string \
    Μήκος ≠ 13 \
    Μη αριθμητικοί χαρακτήρες \
    Λάθος check digit
  ],
  [
    Μήκος 12, 13, 14 \
    Check digit 0 (άθροισμα πολλαπλάσιο του 10)
  ],
)

#tc-table(
  [Περιπτώσεις ελέγχου: CheckBarcode],
  [`TC_BC_01`],
  [Έγκυρο barcode EAN-13],
  [Barcode = `"4006381333931"`],
  [true, κενό μήνυμα],

  [`TC_BC_02`],
  [Έγκυρο barcode με check digit 0],
  [Barcode = `"5201234567870"`],
  [true, κενό μήνυμα],

  [`TC_BC_03`],
  [Λάθος check digit],
  [Barcode = `"4006381333935"`],
  [false, «Λάθος check digit»],

  [`TC_BC_04`],
  [Κενό barcode],
  [Barcode = `""`],
  [false, «Το barcode είναι κενό»],

  [`TC_BC_05`],
  [Μήκος 12, ένα ψηφίο λιγότερο],
  [Barcode = `"400638133393"`],
  [false, «Το barcode πρέπει να έχει 13 ψηφία»],

  [`TC_BC_06`],
  [Μήκος 14, ένα ψηφίο περισσότερο],
  [Barcode = `"40063813339310"`],
  [false, «Το barcode πρέπει να έχει 13 ψηφία»],

  [`TC_BC_07`],
  [Μη αριθμητικός χαρακτήρας],
  [Barcode = `"40063813339A1"`],
  [false, «Το barcode πρέπει να περιέχει μόνο ψηφία»],
)

== FindCountry

Το `FindCountry` ελέγχει πρώτα το barcode με το `CheckBarcode` και επιστρέφει
"Unknown" αν δεν είναι έγκυρο. Αλλιώς μετατρέπει τα 3 πρώτα ψηφία σε αριθμό
και τον αντιστοιχίζει σε χώρα. Τα διψήφια προθέματα του Παραρτήματος Β τα
γράψαμε ως τριψήφια διαστήματα, για παράδειγμα το 30-37 της Γαλλίας ως
300-379, ώστε όλες οι χώρες να ελέγχονται με τον ίδιο τρόπο. Κάθε χώρα είναι
μία κλάση ισοδυναμίας. Για κάθε διάστημα ελέγξαμε το κάτω και το πάνω όριο,
και για τα κενά ανάμεσα στα διαστήματα την τιμή ακριβώς έξω από το όριο.

#ep-table(
  [Κλάσεις ισοδυναμίας και οριακές τιμές: FindCountry],
  [Barcode (string)],
  [
    Έγκυρο barcode με πρόθεμα σε ένα από τα 13 διαστήματα του
    Παραρτήματος Β (μία κλάση ανά χώρα) \
    Έγκυρο barcode με πρόθεμα εκτός όλων των διαστημάτων ("Unknown")
  ],
  [Άκυρο barcode σύμφωνα με το `CheckBarcode` ("Unknown")],
  [
    000, 019 / 020, 029 / 030, 379 / 380 / 381, 399 / 400, 440 / 441,
    689 / 690, 695 / 696, 730, 739, 800, 839 / 840, 849 / 850
  ],
)

#tc-table(
  [Περιπτώσεις ελέγχου: FindCountry],
  [`TC_FC_01`],
  [Ελλάδα, πρόθεμα 520],
  [`"5201234567894"`],
  [Ελλάδα],

  [`TC_FC_02`],
  [Γαλλία, κάτω όριο 300],
  [`"3001234567892"`],
  [Γαλλία],

  [`TC_FC_03`],
  [Ιταλία, πάνω όριο 839],
  [`"8391234567899"`],
  [Ιταλία],

  [`TC_FC_04`],
  [Ισπανία, πάνω όριο 849],
  [`"8491234567896"`],
  [Ισπανία],

  [`TC_FC_05`],
  [Λάθος check digit],
  [`"4006381333935"`],
  [Unknown],

  [`TC_FC_06`],
  [Κενό barcode],
  [`""`],
  [Unknown],

  [`TC_FC_07`],
  [USA, κάτω όριο 000],
  [`"0001234567895"`],
  [Ηνωμένες Πολιτείες (USA)],

  [`TC_FC_08`],
  [USA, πάνω όριο 019],
  [`"0191234567893"`],
  [Ηνωμένες Πολιτείες (USA)],

  [`TC_FC_09`],
  [Internal, κάτω όριο 020],
  [`"0201234567899"`],
  [Περιορισμένη διανομή (Internal/Store use)],

  [`TC_FC_10`],
  [Internal, πάνω όριο 029],
  [`"0291234567890"`],
  [Περιορισμένη διανομή (Internal/Store use)],

  [`TC_FC_11`],
  [Γαλλία, πάνω όριο 379],
  [`"3791234567892"`],
  [Γαλλία],

  [`TC_FC_12`],
  [Βουλγαρία, πρόθεμα 380],
  [`"3801234567898"`],
  [Βουλγαρία],

  [`TC_FC_13`],
  [Γερμανία, κάτω όριο 400],
  [`"4001234567891"`],
  [Γερμανία],

  [`TC_FC_14`],
  [Γερμανία, πάνω όριο 440],
  [`"4401234567899"`],
  [Γερμανία],

  [`TC_FC_15`],
  [Ταϊβάν, πρόθεμα 471],
  [`"4711234567899"`],
  [Ταϊβάν],

  [`TC_FC_16`],
  [Κύπρος, πρόθεμα 529],
  [`"5291234567895"`],
  [Κύπρος],

  [`TC_FC_17`],
  [Πολωνία, πρόθεμα 590],
  [`"5901234567893"`],
  [Πολωνία],

  [`TC_FC_18`],
  [Κίνα, κάτω όριο 690],
  [`"6901234567892"`],
  [Κίνα],

  [`TC_FC_19`],
  [Κίνα, πάνω όριο 695],
  [`"6951234567897"`],
  [Κίνα],

  [`TC_FC_20`],
  [Σουηδία, κάτω όριο 730],
  [`"7301234567899"`],
  [Σουηδία],

  [`TC_FC_21`],
  [Σουηδία, πάνω όριο 739],
  [`"7391234567890"`],
  [Σουηδία],

  [`TC_FC_22`],
  [Ιταλία, κάτω όριο 800],
  [`"8001234567897"`],
  [Ιταλία],

  [`TC_FC_23`],
  [Ισπανία, κάτω όριο 840],
  [`"8401234567895"`],
  [Ισπανία],

  [`TC_FC_24`],
  [Πρόθεμα 030, μετά το Internal],
  [`"0301234567896"`],
  [Unknown],

  [`TC_FC_25`],
  [Πρόθεμα 381, μετά τη Βουλγαρία],
  [`"3811234567897"`],
  [Unknown],

  [`TC_FC_26`],
  [Πρόθεμα 399, πριν τη Γερμανία],
  [`"3991234567896"`],
  [Unknown],

  [`TC_FC_27`],
  [Πρόθεμα 441, μετά τη Γερμανία],
  [`"4411234567898"`],
  [Unknown],

  [`TC_FC_28`],
  [Πρόθεμα 689, πριν την Κίνα],
  [`"6891234567896"`],
  [Unknown],

  [`TC_FC_29`],
  [Πρόθεμα 696, μετά την Κίνα],
  [`"6961234567896"`],
  [Unknown],

  [`TC_FC_30`],
  [Πρόθεμα 850, μετά την Ισπανία],
  [`"8501234567892"`],
  [Unknown],
)

== CalculateCost

Το `CalculateCost` έχει τρεις εισόδους που επηρεάζουν το αποτέλεσμα: την
τιμή, την κατηγορία και το ποσοστό έκπτωσης. Την κατηγορία τη χωρίσαμε σε
κλάσεις με βάση τον συντελεστή ΦΠΑ του Παραρτήματος Γ. Για το ποσοστό
έκπτωσης ελέγξαμε και τις έξι οριακές τιμές γύρω από τα όρια 0 και 100.

#ep-table(
  [Κλάσεις ισοδυναμίας και οριακές τιμές: CalculateCost],
  [Price (float64)],
  [Price ≥ 0],
  [Price < 0],
  [-1, 0],

  [Category (string)],
  [
    24%: Ηλεκτρονικά, Καλλυντικά, Αλκοόλ \
    13%: Τρόφιμα, Αναψυκτικά χωρίς αλκοόλ \
    6%: Βιβλία \
    Άγνωστη κατηγορία (24%)
  ],
  [Καμία (η άγνωστη κατηγορία παίρνει 24%)],
  [Δεν ορίζονται (μη διατεταγμένες τιμές)],

  [Discount (int)],
  [Ακέραιος στο διάστημα [0, 100]],
  [Discount < 0 ή Discount > 100],
  [-1, 0, 1, 99, 100, 101],
)

Σε όλες τις περιπτώσεις η έξοδος είναι η τριάδα (PriceWithSupplyAndTaxes,
Discount, FinalPrice). Όπου δεν αναφέρεται διαφορετικά, η είσοδος είναι
Price = 10 και Discount = 10.

#tc-table(
  [Περιπτώσεις ελέγχου: CalculateCost],
  [`TC_CALC_01`], [ΦΠΑ 6%], [Category = «Βιβλία»], [(11.13, 1.11, 10.02)],
  [`TC_CALC_02`],
  [ΦΠΑ 24%],
  [Category = «Ηλεκτρονικά»],
  [(13.02, 1.30, 11.72)],

  [`TC_CALC_03`],
  [ΦΠΑ 13% και στρογγυλοποίηση (11.865 → 11.87)],
  [Category = «Τρόφιμα»],
  [(11.87, 1.19, 10.68)],

  [`TC_CALC_04`],
  [Άγνωστη κατηγορία, ΦΠΑ 24%],
  [Category = «Παιχνίδια»],
  [(13.02, 1.30, 11.72)],

  [`TC_CALC_05`],
  [Discount στο κάτω όριο],
  [Category = «Βιβλία», Discount = 0],
  [(11.13, 0, 11.13)],

  [`TC_CALC_06`],
  [Discount μετά το κάτω όριο],
  [Category = «Βιβλία», Discount = 1],
  [(11.13, 0.11, 11.02)],

  [`TC_CALC_07`],
  [Discount πριν από το πάνω όριο],
  [Category = «Βιβλία», Discount = 99],
  [(11.13, 11.02, 0.11)],

  [`TC_CALC_08`],
  [Discount στο πάνω όριο],
  [Category = «Βιβλία», Discount = 100],
  [(11.13, 11.13, 0)],

  [`TC_CALC_09`],
  [Discount κάτω από το κάτω όριο],
  [Category = «Βιβλία», Discount = -1],
  [(0, 0, 0)],

  [`TC_CALC_10`],
  [Discount πάνω από το πάνω όριο],
  [Category = «Βιβλία», Discount = 101],
  [(0, 0, 0)],

  [`TC_CALC_11`],
  [Μηδενική τιμή],
  [Category = «Βιβλία», Price = 0],
  [(0, 0, 0)],

  [`TC_CALC_12`],
  [Αρνητική τιμή],
  [Category = «Βιβλία», Price = -1],
  [(0, 0, 0)],
)

== CalculateTotalCost

#ep-table(
  [Κλάσεις ισοδυναμίας και οριακές τιμές: CalculateTotalCost],
  todo[], [], [], [],
)
#tc-table(
  [Περιπτώσεις ελέγχου: CalculateTotalCost],
  [`TC_TOT_01`], todo[], [], [],
)

== CountryProducts

#ep-table(
  [Κλάσεις ισοδυναμίας και οριακές τιμές: CountryProducts],
  todo[], [], [], [],
)
#tc-table(
  [Περιπτώσεις ελέγχου: CountryProducts],
  [`TC_CP_01`], todo[], [], [],
)

== ProductsToOrder

#ep-table(
  [Κλάσεις ισοδυναμίας και οριακές τιμές: ProductsToOrder],
  todo[], [], [], [],
)
#tc-table(
  [Περιπτώσεις ελέγχου: ProductsToOrder],
  [`TC_PTO_01`], todo[], [], [],
)

== CalculateOrder

#ep-table(
  [Κλάσεις ισοδυναμίας και οριακές τιμές: CalculateOrder],
  todo[], [], [], [],
)
#tc-table(
  [Περιπτώσεις ελέγχου: CalculateOrder],
  [`TC_ORD_01`], todo[], [], [],
)

= Υλοποίηση Ελέγχων Μονάδας

Γράψαμε τους ελέγχους στα αρχεία `warehouse/*_test.go` και τους εκτελούμε με
`go test -v ./warehouse`. Για κάθε μέθοδο της βιβλιοθήκης γράψαμε
τουλάχιστον 2 μεθόδους ελέγχου, που καλύπτουν κανονικές ροές, οριακές τιμές
και εσφαλμένες εισόδους. Στα ονόματα ακολουθούμε το μοτίβο
`Test<Μέθοδος>_<Σενάριο>_<Αναμενόμενο>`.

#figure(
  table(
    columns: (1fr, auto, auto, auto, auto),
    table.header(
      [Μέθοδος],
      [Κανονική ροή],
      [Οριακές τιμές],
      [Εσφαλμένη είσοδος],
      [Σύνολο],
    ),
    [CheckBarcode], todo[], [], [], [],
    [FindCountry], todo[], [], [], [],
    [CalculateCost], todo[], [], [], [],
    [CalculateTotalCost], todo[], [], [], [],
    [CountryProducts], todo[], [], [], [],
    [ProductsToOrder], todo[], [], [], [],
    [CalculateOrder], todo[], [], [], [],
  ),
  caption: [Πλήθος μεθόδων ελέγχου ανά κατηγορία],
)

= Αναφορές Σφαλμάτων

Για κάθε μέθοδο τεκμηριώνουμε τουλάχιστον ένα σφάλμα που εντοπίσαμε κατά την
ανάπτυξη (TDD / Regression Testing), με τρία στιγμιότυπα: την αποτυχία της
δοκιμής (Red Phase), τη λανθασμένη συμπεριφορά στην εφαρμογή (App Behavior)
και την επιτυχία όλων των δοκιμών μετά τη διόρθωση (Green Phase).

== CheckBarcode
Το σφάλμα το εντοπίσαμε με το `TC_BC_02`. Ο τύπος του Παραρτήματος Α για το
check digit είναι $(10 - (S mod 10)) mod 10$, όπου $S$ το άθροισμα των 12
πρώτων ψηφίων με βάρη 1 και 3. Στην υλοποίηση είχαμε παραλείψει το
εξωτερικό $mod 10$ και υπολογίζαμε $10 - (S mod 10)$.

Για τα περισσότερα barcodes τα δύο δίνουν το ίδιο αποτέλεσμα, γιατί το
$10 - (S mod 10)$ είναι ήδη μεταξύ 1 και 9. Η διαφορά εμφανίζεται μόνο όταν
το $S$ είναι πολλαπλάσιο του 10. Για το barcode "5201234567870" το άθροισμα
είναι $S = 100$, άρα $S mod 10 = 0$ και $10 - 0 = 10$. Το check digit
βγαίνει 10 αντί για 0, δεν ταιριάζει με το 13ο ψηφίο, και ένα έγκυρο
barcode απορρίπτεται ως «Λάθος check digit». Το εξωτερικό $mod 10$
μετατρέπει το 10 σε 0, γι' αυτό και χρειάζεται.

Το `TC_BC_01` δεν μπορούσε να εντοπίσει το σφάλμα, γιατί εκεί $S = 89$ και
το check digit βγαίνει σωστά 1 και με τους δύο τύπους. Το `TC_BC_02` το
προσθέσαμε ως οριακή τιμή ακριβώς για την περίπτωση όπου το check digit
είναι 0. Η διόρθωση ήταν να προστεθεί το εξωτερικό $mod 10$.

#bug-report(
  id: [`BUG-BC-01`],
  test: [
    `TestCheckBarcode_ValidBarcode_ReturnsTrue` \
    `/TC_BC_02_ValidCheckDigitZero`
  ],
  purpose: [
    Ελέγχει ότι το `CheckBarcode` δέχεται έγκυρο barcode του οποίου το
    check digit είναι 0.
  ],
  failure: [
    `barcode_test.go:21: CheckBarcode("5201234567870")` \
    `= false, want true` \
    `barcode_test.go:24: CheckBarcode("5201234567870")` \
    `errMsg = "Λάθος check digit", want empty`
  ],
  expected: [true, κενό μήνυμα],
  actual: [false, «Λάθος check digit»],
  red: "assets/bug-bc-01-red.png",
  green: "assets/bug-bc-01-green.png",
)

== FindCountry
Το σφάλμα το εντοπίσαμε με τα `TC_FC_24` έως `TC_FC_26`. Στην πρώτη
υλοποίηση κάθε `case` του `switch` έλεγχε μόνο το πάνω όριο του διαστήματος
(για παράδειγμα πρόθεμα ≤ 379 για τη Γαλλία) και βασιζόταν στη σειρά των
`case` για το κάτω όριο. Αυτό δουλεύει μόνο όταν τα διαστήματα είναι
συνεχόμενα. Ανάμεσα στο 029 και στο 300 όμως, και ανάμεσα στο 380 και στο
400, υπάρχουν προθέματα που δεν ανήκουν σε καμία χώρα. Το πρόθεμα 030 έπεφτε
στο `case` της Γαλλίας και τα 381 και 399 στο `case` της Γερμανίας, αντί να
επιστρέφουν "Unknown".

Οι έλεγχοι με τις γνωστές χώρες περνούσαν όλοι, γιατί κανένας δεν
χρησιμοποιούσε πρόθεμα από τα κενά. Το σφάλμα φάνηκε μόνο όταν προσθέσαμε
τις τιμές ακριβώς έξω από τα όρια κάθε διαστήματος. Το διορθώσαμε
προσθέτοντας και κάτω όριο στα `case` της Internal διανομής, της Γαλλίας και
της Γερμανίας.

#bug-report(
  id: [`BUG-FC-01`],
  test: [
    `TestFindCountry_UnknownPrefix_ReturnsUnknown` \
    `/TC_FC_24_AfterInternal`, `TC_FC_25_AfterBulgaria`,
    `TC_FC_26_BeforeGermany`
  ],
  purpose: [
    Ελέγχει ότι το `FindCountry` επιστρέφει "Unknown" για έγκυρα barcodes
    με πρόθεμα ακριβώς έξω από τα διαστήματα του Παραρτήματος Β.
  ],
  failure: [
    `FindCountry("0301234567896") = "Γαλλία", want "Unknown"` \
    `FindCountry("3811234567897") = "Γερμανία", want "Unknown"` \
    `FindCountry("3991234567896") = "Γερμανία", want "Unknown"`
  ],
  expected: ["Unknown" και για τα τρία barcodes],
  actual: ["Γαλλία" για το 030, "Γερμανία" για τα 381 και 399],
  red: "assets/bug-fc-01-red.png",
  green: "assets/bug-fc-01-green.png",
)

== CalculateCost
Το σφάλμα το εντοπίσαμε με το `TC_CALC_09`. Χωρίς έλεγχο της εισόδου, ένα
αρνητικό ποσοστό έκπτωσης γινόταν αρνητικό ποσό έκπτωσης, και η τελική τιμή
έβγαινε μεγαλύτερη από την τιμή χωρίς έκπτωση. Το διορθώσαμε προσθέτοντας
στην αρχή της συνάρτησης έλεγχο που επιστρέφει (0, 0, 0) για Discount εκτός
[0, 100] ή αρνητική τιμή, σύμφωνα με την παραδοχή Π2.

#bug-report(
  id: [`BUG-CALC-01`],
  test: [
    `TestCalculateCost_InvalidInput_ReturnsZero` \
    `/TC_CALC_09_NegativeDiscount`
  ],
  purpose: [
    Ελέγχει ότι το `CalculateCost` απορρίπτει αρνητικό ποσοστό έκπτωσης
    και επιστρέφει (0, 0, 0).
  ],
  failure: [
    `cost_test.go:81: CalculateCost = (11.13, -0.11, 11.24),` \
    `want (0, 0, 0)`
  ],
  expected: [(0, 0, 0)],
  actual: [
    (11.13, -0.11, 11.24): αρνητική έκπτωση και τελική τιμή μεγαλύτερη από
    την τιμή χωρίς έκπτωση
  ],
  red: "assets/bug-calc-01-red.png",
  green: "assets/bug-calc-01-green.png",
)

== CalculateTotalCost
#bug-report(id: [`BUG-TOT-01`])

== CountryProducts
#bug-report(id: [`BUG-CP-01`])

== ProductsToOrder
#bug-report(id: [`BUG-PTO-01`])

== CalculateOrder
#bug-report(id: [`BUG-ORD-01`])

= Code Coverage

#shot(none, [Κάλυψη κώδικα της βιβλιοθήκης (`go test -cover`)])
#shot(none, [Αναλυτική κάλυψη ανά γραμμή (`go tool cover -html`)])

= Εφαρμογή Πελάτη

#figure(
  table(
    columns: (auto, 1fr, auto),
    table.header[\#][Λειτουργία][Μέθοδοι βιβλιοθήκης],
    [1],
    [Καταχώρηση προϊόντος, με έλεγχο Barcode και αυτόματη χώρα],
    [CheckBarcode, FindCountry],

    [2], [Προβολή αποθέματος και συνολικού κόστους], [CalculateTotalCost],
    [3], [Υπολογισμός τελικής τιμής προϊόντος], [CalculateCost],
    [4], [Αναζήτηση προϊόντων ανά χώρα], [CountryProducts],
    [5], [Εντοπισμός προϊόντων προς παραγγελία], [ProductsToOrder],
    [6], [Παραγγελία], [CalculateOrder],
  ),
  caption: [Λειτουργίες της εφαρμογής και μέθοδοι που χρησιμοποιούν],
)

#shot(none, [Καταχώρηση προϊόντος])
#shot(none, [Προβολή αποθέματος και συνολικού κόστους])
#shot(none, [Υπολογισμός τελικής τιμής προϊόντος])
#shot(none, [Αναζήτηση προϊόντων ανά χώρα])
#shot(none, [Προϊόντα προς παραγγελία])
#shot(none, [Παραγγελία])

= Οδηγίες Εκτέλεσης

```sh
go test -v ./warehouse                     # όλοι οι έλεγχοι
go test -cover ./warehouse                 # κάλυψη κώδικα
go test -coverprofile=coverage.out ./warehouse
go tool cover -html=coverage.out           # κάλυψη ανά γραμμή
go run ./cmd/client                        # εφαρμογή πελάτη
```
