package parsing

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
)

type Service interface {
	ParseCSV(csvFile io.Reader) error
}

type parsingService struct {
	log *log.Logger
}

func (ps *parsingService) ParseCSV(csvFile io.Reader) error {

	csvReader := csv.NewReader(csvFile)

	headers, err := csvReader.Read()
	if err != nil {

		return err

	}

	parsedCSV := make([]map[string]string, 0)
	for {

		csvEntry, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {

			return err

		}

		mappedEntry := make(map[string]string)
		for idx, keys := range headers {

			mappedEntry[keys] = csvEntry[idx]

		}

		parsedCSV = append(parsedCSV, mappedEntry)

	}

	parsedCSVJson, err := json.Marshal( parsedCSV )
	if err != nil {
		return err
	}
	ps.log.Print(string(parsedCSVJson))

	return nil
}

func NewService(log *log.Logger) Service {

	return &parsingService{log: log}

}
