package mods

import "../monkey"

// EdgeDetectConvolution ...
var EdgeDetectConvolution = monkey.ConvolutionMatrix{
	{0, 1, 0},
	{1, -4, 1},
	{0, 1, 0},
}

//
// EdgeDetect performs a edge-detection on the image...
//
func EdgeDetect(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(EdgeDetectConvolution)
	return newMatrix
}
