package main

import (
	"program/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	//result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.LoadData()
		priceJob.Process()

	}

}
