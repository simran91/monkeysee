package mods

import "../monkey"

// import "image/color"

// AverageBlurConvolution ...
// var AverageBlurConvolution = monkey.ConvolutionMatrix{
// 	{1, 1, 1, 1, 1, 1, 1},
// 	{1, 1, 1, 1, 1, 1, 1},
// 	{1, 1, 1, 1, 1, 1, 1},
// 	{1, 1, 1, 1, 1, 1, 1},
// 	{1, 1, 1, 1, 1, 1, 1},
// 	{1, 1, 1, 1, 1, 1, 1},
// 	{1, 1, 1, 1, 1, 1, 1},
// }
var AverageBlurConvolution = monkey.ConvolutionMatrix{
	{1, 1, 1},
	{1, 1, 1},
	{1, 1, 1},
}

//
// AverageBlur performs an average blur...
//
func AverageBlur(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(AverageBlurConvolution)
	return newMatrix
}
