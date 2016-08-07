package mods

import "../monkey"

// EmbossConvolution ...
var PlayOneConvolution = monkey.ConvolutionMatrix{
	{4, 4, 4, 0, 0, 0, 0},
	{4, 2, 2, 0, 0, 0, 0},
	{4, 2, 1, 0, 0, 0, 0},
	{4, 2, 1, 0, -1, -2, -4},
	{0, 0, 0, 0, -1, -2, -4},
	{0, 0, 0, 0, -2, -2, -4},
	{0, 0, 0, 0, -4, -4, -4},
}

//
// PlayOne ...
//
func PlayOne(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(PlayOneConvolution)
	return newMatrix
}
