package main

import (
	"strconv"
)

// return input table, removed index list, error
func parseInputCsv(records [][]string) [][]*int {
	inputTable := [][]*int{}
	for _, record := range records {
		row := []*int{}
		for _, cell := range record {
			v, err := strconv.Atoi(cell)
			if err != nil {
				row = append(row, nil)
			} else {
				if v != 0 && v != 1 {
					v = 1
				}
				row = append(row, &v)
			}
		}
		inputTable = append(inputTable, row)
	}
	return inputTable
}

func parseOutputCsv(records [][]string) [][]*int {
	outputTable := [][]*int{}
	for _, record := range records {
		row := []*int{}
		for _, cell := range record {
			v, err := strconv.Atoi(cell)
			if err != nil {
				row = append(row, nil)
				continue
			}
			if v != 0 && v != 1 {
				v = 1
			}
			row = append(row, &v)
		}
		outputTable = append(outputTable, row)
	}
	return outputTable
}
