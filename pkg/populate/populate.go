package main

import (
	"delivery-much-challenge/conf"
	"delivery-much-challenge/datasource"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func csvProductsToModel(records [][]string) []datasource.Product {
	items := []datasource.Product{}
	for i := 0; i < len(records); i++ {
		record := records[i]
		name := record[0]
		price := record[1]
		quantity := record[2]

		newPrice, _ := strconv.ParseFloat(price, 64)
		newQuantity, _ := strconv.ParseFloat(quantity, 64)

		product := datasource.Product{
			Name:     name,
			Price:    newPrice,
			Quantity: newQuantity,
		}

		items = append(items, product)
	}

	return items
}

func main() {
	environmentConfiguration := conf.Env(godotenv.Load)
	records := readCsvFile("./resource/products.csv")
	models := csvProductsToModel(records)

	database := datasource.DatabaseConnect(environmentConfiguration)
	database.AutoMigrate(&datasource.Product{})

	repository := datasource.NewProductRepository(database)
	repository.Populate(models)
}
