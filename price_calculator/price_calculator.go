package price

import (
	"fmt"
	"training/price_calculator/cmdmanager"
	"training/price_calculator/filemanager"
	"training/price_calculator/prices"
)

type Processor interface {
	Process(chan bool, chan error)
}

func App(debug bool) {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		var jobProcessor Processor
		if debug {
			cmd := cmdmanager.New() //debug mode
			jobProcessor = prices.NewTaxIncludedPriceJob(cmd, taxRate)
		} else {
			fm := filemanager.New("price_calculator/prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
			jobProcessor = prices.NewTaxIncludedPriceJob(fm, taxRate)
		}
		go jobProcessor.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
