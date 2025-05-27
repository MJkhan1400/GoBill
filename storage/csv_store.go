package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/MJkhan1400/GoBill/config"
	"github.com/MJkhan1400/GoBill/model"
)

var Products = make(map[string]model.Product)

func LoadProducts() {
	// Try to open the CSV file specified in config.ProductCSV
	file, err := os.Open(config.ProductCSV)
	if err != nil {
		// If the file doesn't exist or can't be opened, print a message and return
		fmt.Println("No existing product data found. Starting fresh.")
		return
	}
	defer file.Close() // Ensure the file is closed when the function ends

	// Create a new CSV reader for the opened file
	r := csv.NewReader(file)

	// Read all records (rows) from the CSV file
	records, err := r.ReadAll()
	if err != nil {
		// If reading fails, print an error and return
		fmt.Println("Failed to read product data: ", err)
		return
	}

	// Loop through each record (row) in the CSV
	for _, record := range records {
		if len(record) < 3 {
			// Skip rows that don't have at least 3 columns (name, price, quantity)
			continue
		}

		// Parse the price from string to float64
		price, _ := strconv.ParseFloat(record[1], 64)
		// Parse the quantity from string to int
		quantity, _ := strconv.Atoi(record[2])

		// Add the product to the Products map
		Products[record[0]] = model.Product{
			Name:     record[0],
			Price:    price,
			Quantity: quantity,
		}
	}
	fmt.Println("Loaded products from CSV.")
}

func SaveProducts() {
	file, err := os.Create(config.ProductCSV)
	if err != nil {
		fmt.Println("Error saving product data: ", err)
		return
	}
	defer file.Close()

	write := csv.NewWriter(file)
	for _, product := range Products {
		write.Write([]string{
			product.Name,
			fmt.Sprintf("%.2f", product.Price),
			strconv.Itoa(product.Quantity),
		})
	}
}
