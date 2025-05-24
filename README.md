BillGo/
├── main.go                  # Entry point with CLI logic
├── config/
│   └── config.go            # Load/save config (CSV path, VAT)
├── billing/
│   └── billing.go           # Handles billing logic and file generation
├── product/
│   └── product.go           # Product struct, load/save/add product
├── data/
│   ├── products.csv         # Product list (editable)
│   └── config.txt           # Stores CSV path and VAT rate
├── bills/
│   └── bill_table_xx_yyyy-mm-dd_hh-mm-ss.txt  # Generated bills
└── README.md                # Instructions and usage
