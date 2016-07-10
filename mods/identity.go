package mods

import "../monkey"

// EmbossConvolution ...
var IdentityConvolution = monkey.ConvolutionMatrix{
	{0, 0, 0,},
	{0, 1, 0,},
	{0, 0, 0,},
}

//
// Identity saves the image as is... (an identity matrix doesn't alter the matrix at all) :) 
//
func Identity(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(IdentityConvolution)
	return newMatrix
}
