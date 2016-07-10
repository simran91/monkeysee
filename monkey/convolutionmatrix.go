package monkey

//
// ConvolutionMatrix defines how we store our convolution matrices...
//
type ConvolutionMatrix [][]int8

//
// GetWidth returns the height of the image
//
func (cm ConvolutionMatrix) GetWidth() int {
	return len(cm)
}

//
// GetHeight returns the height of the image
//
func (cm ConvolutionMatrix) GetHeight() int {
	return len(cm[0])
}
