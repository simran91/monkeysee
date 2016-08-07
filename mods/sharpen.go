package mods

import "../monkey"

// SharpenConvolution ...
var SharpenConvolution = monkey.ConvolutionMatrix{
	{0, -1, 0},
	{-1, 5, -1},
	{0, -1, 0},
}

//
// Sharpen performs a sharpening of the image...
//
func Sharpen(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(SharpenConvolution)
	return newMatrix
}
