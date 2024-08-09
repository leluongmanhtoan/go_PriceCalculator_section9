package main

import (
	"fmt"
	"program/filemanager"
	"program/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	//result := make(map[float64][]float64)
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errChans[index])

		//if err != nil {
		//	fmt.Println("Could not process job")
		//	fmt.Println(err)
		//}

	}
	for index, _ := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")

		}

	}

	/*for _, errorChan := range errChans {
		<-errorChan
	}
	for _, doneChan := range doneChans {
		<-doneChan
	}*/
}
