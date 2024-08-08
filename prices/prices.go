package prices

import (
	"bufio"
	"fmt"
	"os"
	"program/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		file.Close()
		return
	}
	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	job.InputPrices = prices
	file.Close()

}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxrate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxrate,
	}

}
