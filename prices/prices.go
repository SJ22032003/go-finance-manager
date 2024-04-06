package prices

import (
	"fmt"
	"time"

	conversion "github.com/SJ22032003/go-finance-manager-app/conversion"
	file_manager "github.com/SJ22032003/go-finance-manager-app/file_manager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64                  `json:"tax_rate"`
	InputPrices       []float64                `json:"input_prices"`
	TaxIncludedPrices map[string]float64       `json:"tax_included_prices"`
	IOManager         file_manager.FileManager `json:"-"`
}

func NewTaxIncludedPriceJob(fm file_manager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: fm,
	}
}

func (job *TaxIncludedPriceJob) Process(channel chan bool, errorChannel chan error) {
	job.TaxIncludedPrices = make(map[string]float64)
	for _, price := range job.InputPrices {
		key := fmt.Sprintf("%.2f", price)
		job.TaxIncludedPrices[key] = price * (1 + job.TaxRate)
	}

	time.Sleep(2 * time.Second) // Simulate a long running process

	path := fmt.Sprintf("tax_included_prices_%0.0f.json", job.TaxRate*100)
	job.IOManager.SetOutputPath(path)

	err := job.IOManager.WriteFileToJSON(job)
	if err != nil {
		fmt.Println(err)
		errorChannel <- err
		return
	}

	channel <- true

}

func (job *TaxIncludedPriceJob) LoadData() {
	const fileToRead = "prices.txt"
	job.IOManager.SetInputPath(fileToRead)

	lines, err := job.IOManager.ReadFileManager()
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = make([]float64, len(lines))

	for _, line := range lines {
		price, err := conversion.StringToFloat64(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		job.InputPrices = append(job.InputPrices, price)
	}

}
