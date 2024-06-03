package importer

import (
	"github.com/tealeg/xlsx"
)

func ReadFileXLSX(pathFile string) ([][]string, error) {
	file, err := xlsx.OpenFile(pathFile)
	if err != nil {
		return nil, err
	}

	var data [][]string

	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			var rowData []string
			for _, cell := range row.Cells {
				rowData = append(rowData, cell.String())
			}
			data = append(data, rowData)
		}
	}

	return data, nil
}
