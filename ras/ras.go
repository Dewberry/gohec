package ras

// OutputTable contains table data from HEC-RAS outputs stored in HDF files.
type OutputTable struct {
	TableRows int
	TableCols int
	TableData []float32
}

// FilterTable reads output from a ras HDF output table
// Need to add error handling
func FilterTable(colID int, t *OutputTable) []float32 {

	result := make([]float32, 0)
	var rowIDX int

	for i := 0; i < t.TableRows; i++ {

		switch i {

		case 0:
			rowIDX = colID

		default:
			rowIDX += t.TableCols
		}

		result = append(result, t.TableData[int(rowIDX)])
	}

	return result
}
