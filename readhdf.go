package main

import (
	"fmt"

	"github.com/dewberry/gohec/ras"
	"gonum.org/v1/hdf5"
)

// Test data hard coded for dev.
var testFile string = "DC_F01_B05_E0100.p01.hdf"
var group string = "/Results/Unsteady/Output/Output Blocks/Base Output/Unsteady Time Series/2D Flow Areas/D01/"
var table string = "Depth"

func main() {

	// More test data hard coded for dev.
	colIDs := []int{68039, 68040, 68041}

	// Open RAS HDF output file in read only mode
	f, err := hdf5.OpenFile(testFile, hdf5.F_ACC_RDONLY)
	if err != nil {
		panic(err)
	}

	// Open the HDF Group
	dataGroup, err := f.OpenGroup(group)
	if err != nil {
		panic(err)
	}

	defer dataGroup.Close()

	// Placeholder for searching for relevant data later
	n, _ := dataGroup.NumObjects()

	for i := 0; i < int(n); i++ {
		name, _ := dataGroup.ObjectNameByIndex(uint(i))

		switch name {

		case table:
			fmt.Println(i, name)

		default:
			continue
		}
	}

	// Open specified table
	dataSet, err := dataGroup.OpenDataset(table)
	if err != nil {
		panic(err)
	}
	defer dataSet.Close()

	// Read in table dimensions
	dataSpace := dataSet.Space()
	defer dataSpace.Close()

	tableDims, _, _ := dataSpace.SimpleExtentDims()
	tableRows := tableDims[0]
	tableCols := tableDims[1]

	// Instantiate OutputTable with info from the dataSpace & dataSet objects
	rasResults := ras.OutputTable{TableRows: int(tableRows), TableCols: int(tableCols), TableData: make([]float32, tableRows*tableCols)}

	// Read data to buffer
	err = dataSet.Read(&rasResults.TableData)
	if err != nil {
		panic(err)
	}

	// Filter data by index
	// Need to add index look-up functions
	for _, colID := range colIDs {
		qData := ras.FilterTable(colID, &rasResults)
		fmt.Printf("\n%d:: data: %v\n", colID, qData)
	}

	fmt.Println(colIDs)

	f.Close()
}
