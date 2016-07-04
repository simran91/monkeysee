package mimage

import "fmt"

// debugPrintData prints the RGBAMatrix to STDOUT...
//
func debugPrintMatrix(matrix ImageMatrix) {
	for x, rows := range matrix {
		for y, colour := range rows {
			fmt.Printf("x:%v y:%v colour:%v\n", x, y, colour)
		}
	}
}
