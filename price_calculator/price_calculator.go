package price

import (
	"fmt"
	"training/price_calculator/cmdmanager"
	"training/price_calculator/filemanager"
	"training/price_calculator/prices"
)

type Processor interface {
	Process() error
}

func App(debug bool) {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		var jobProcessor Processor
		if debug {
			cmd := cmdmanager.New() //debug mode
			jobProcessor = prices.NewTaxIncludedPriceJob(cmd, taxRate)
		} else {
			fm := filemanager.New("price_calculator/prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
			jobProcessor = prices.NewTaxIncludedPriceJob(fm, taxRate)
		}
		err := jobProcessor.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
