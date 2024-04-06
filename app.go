package main

import (
	"fmt"

	file_manager "github.com/SJ22032003/go-finance-manager-app/file_manager"
	prices "github.com/SJ22032003/go-finance-manager-app/prices"
)

func main() {

	taxRate := []float64{0.2, 0.3, 0.4, 0.5, 0.6}
	doneRoutine := make([]chan bool, len(taxRate))
	errorRoutine := make([]chan error, len(taxRate))

	for index, rate := range taxRate {
		priceJob := prices.NewTaxIncludedPriceJob(file_manager.FileManager{}, rate)

		doneRoutine[index] = make(chan bool)
		errorRoutine[index] = make(chan error)

		priceJob.LoadData()
		go priceJob.Process(doneRoutine[index], errorRoutine[index])
	}

	for index := range taxRate {
		select {
		case err := <-errorRoutine[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneRoutine[index]:
			fmt.Println("Done")
		}
	}

}
