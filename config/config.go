package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ConfigFile = "config.txt"
	ProductCSV string
	VatRate    float64 = 5.0
)

func LoadConfig(inputScanner *bufio.Scanner) bool {
	file, err := os.Open(ConfigFile)
	if err != nil {
		fmt.Println("--- Initial Setup ---")
		fmt.Println("Please enter the path to the product data: ")
		inputScanner.Scan()

		ProductCSV = strings.TrimSpace(inputScanner.Text())
		if ProductCSV == "" {
			ProductCSV = "products.csv"
		}

		fmt.Println("Please enter the VAT rate (default is 5% for UAE): ")
		inputScanner.Scan()

		vatRate, err := strconv.ParseFloat(inputScanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid VAT rate. Using default 5%.")
			vatRate = 5.0
		}
		VatRate = vatRate
		SaveConfig()
		return true
	} else {
		fileScanner := bufio.NewScanner(file)

		if fileScanner.Scan() {
			ProductCSV = fileScanner.Text()
		}

		if fileScanner.Scan() {
			VatRate, _ = strconv.ParseFloat(fileScanner.Text(), 64)
		}
		file.Close()
		fmt.Println("Configuration loaded successfully.")
		return false
	}
}

func SaveConfig() {
	file, err := os.Create(ConfigFile)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		return
	}

	defer file.Close()
	fmt.Fprintln(file, ProductCSV)
	fmt.Fprintf(file, "%.2f\n", VatRate)
}
