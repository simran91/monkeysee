package mods

import "../monkey"

// ApplyConvolutionWithSampleFunction ...
var ApplyConvolutionWithSampleFunctionMatrix = monkey.ConvolutionMatrix{
	{1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
}

//
// ApplyConvolutionWithSampleFunction applies influence to the image based on the function ipmlementation...
// It's written more so that you can look at the implementation of it in the monkey/ directory and create your own more
// sensible filters... :)
//
func ApplyConvolutionWithSampleFunction(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolutionWithSampleFunction(ApplyConvolutionWithSampleFunctionMatrix)
	return newMatrix
}
