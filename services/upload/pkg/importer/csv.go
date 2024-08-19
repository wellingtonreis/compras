package importer

import (
	"encoding/csv"
	"os"
)

func ReadCSVFile(pathFile string) ([][]string, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var data [][]string

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		data = append(data, record)
	}
	return data, nil
}
