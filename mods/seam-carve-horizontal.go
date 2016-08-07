package mods

import "../monkey"

//
// SeamCarveHorizontal ... 
//
func SeamCarveHorizontal(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.SeamCarveHorizontal()
	return newMatrix
}
