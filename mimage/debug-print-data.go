package mimage

import "fmt"

// debugPrintData prints the RGBAMatrix to STDOUT...
//
func debugPrintMatrix(rgba ImageMatrix) {
	for label, matrix := range rgba {
		for _, row := range matrix {
			fmt.Printf("%v: %v\n", label, row)
		}
	}
}
