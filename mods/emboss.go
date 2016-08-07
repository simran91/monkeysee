package mods

import "../monkey"

// EmbossConvolution ...
var EmbossConvolution = monkey.ConvolutionMatrix{
	{-2, -1, 0},
	{-1, 1, 1},
	{0, 1, 2},
}

//
// Emboss performs a embossing of the image...
//
func Emboss(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(EmbossConvolution)
	return newMatrix
}
