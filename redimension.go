package main


import "bufio"
import "io"
import "strings"


// redimension is where the actual work starts to be done in this program.
//
// Note that *what* is to be done is not decided in this func.
// *What* is to be done is communicated with this func via its parameters.
func redimension(w io.Writer, r io.Reader, number_of_columns uint64) {

	// If the user has specifed zero (0) for the number of columns then we should output
	// a single '\n' character an exit.
	//
	// This has the effect of discarding all the data!
		if 0 == number_of_columns {
			io.WriteString(w, "\n")
/////////////////////// RETURN
			return
		}

	// We want to be able to read the input line by line, so need to put the io.Reader
        // into a bufio.Reader.
		bioreader := bufio.NewReader(r)

	// Redimension.
		var line_separator byte = '\n'
		cell_separator := "\t"

		index := uint64(0)

		new_row_index    := uint64(0)
		new_column_index := uint64(0)

		for line, err := bioreader.ReadString(line_separator); nil == err; line, err = bioreader.ReadString(line_separator) {

			// A line has a '\n' character at the end.
			// The '\n' character is a control character, and not data.
			//
			// We define a row as being a line without the '\n' control character at the end.
				row := line[0:len(line)-1]			

			// A row is made up of one of more cells separated by '\t' chacters. The 
			// The '\t' character is a control character, and not data.
			//
			// We sprint the row string by the '\t' charactes to get the cells.
				cells := strings.Split(row, cell_separator)

			// For each cell ....
			//
				for _,cell := range cells {

					// Calculate the new column index and the new row index.
						new_column_index = index % number_of_columns
						new_row_index    = (index - new_column_index) / number_of_columns


					// Output cell separator.
						if 0 != new_row_index && 0 == new_column_index {
							io.WriteString(w, "\n")
						} else if 0 != new_column_index {
							io.WriteString(w, "\t")
						}

					// Output cell.
						_,_ = io.WriteString(w, cell)

					// Increment the overall counter.
						index++
				}
		}

	// Regardless of what the cell separator is, at the end we output a '\n'.
	//
	// We consider this necessary for a text file to be well-formed.
	// 
		_,_ = io.WriteString(w, "\n")

}
