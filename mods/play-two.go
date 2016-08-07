package mods

import "../monkey"

// EmbossConvolution ...
var PlayTwoConvolution = monkey.ConvolutionMatrix{
	{2, 1, 0},
	{1, 1, -1},
	{0, -1, -2},
}

//
// PlayTwo ...
//
func PlayTwo(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(PlayTwoConvolution)
	return newMatrix
}
