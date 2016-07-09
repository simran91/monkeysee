package monkey

import "image/color"

//
// This returns a sipmle "1" to multiply the ConvolutionMatrix weight by... so esentially, it does't modify the weight
// and leaves it to whatever you set in the matrix... it's a helper function for ApplyConvolution
//
func dontModifyConvolutionMatrixWeights(im ImageMatrix, imagePositionX int, imagePositionY int, kernelPixelX int, kernelPixel int, colour color.RGBA, distance float64) int {
	return 1
}

//
// This is a sample function that we can look up, just to see an example...
//
func convolutionMatrixSampleFunction(im ImageMatrix, imagePositionX int, imagePositionY int, kernelPixelX int, kernelPixel int, colour color.RGBA, distance float64) int {
	if (distance < 1) {
		return 1
	}

	if (colour.R > 150 && colour.G < 100 && colour.B < 100) {
		return int(5 * distance)
	}

	return 5
}