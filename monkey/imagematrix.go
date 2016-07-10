package monkey

import "image/color"
import "log"
import "math"

//
// ImageMatrix defines how we store our matrix of Colours...
//
type ImageMatrix [][]color.Color

// GetKernelMatrix returns an ImageMatrix around the pixel (x,y) based on the size of the kernel we requested
// eg. A GetKernelMatrix(5, 5, 1) will return an ImageMatrix that is built from the
//     pixels: 4,4 5,4 6,4
//  		   4,5 5,5 6,5
// 			   4,6 5,6 6,6
// The above will be returned as it's own new ImageMatrix
//
// No matter what the pixel or the size of the kenel being requested, it will always return an
// a square ImageMatrix of each side = (size*2 + 1). The square matrix should make it easier
// for us to do matrix maths on.
//
// See https://en.wikipedia.org/wiki/Kernel_(image_processing)#Origin for how to do Gaussian Blur's,
// Image Sharpen's, etc, etc, etc...
//
func (im ImageMatrix) GetKernelMatrix(origX, origY, size int) ImageMatrix {

	width := im.GetWidth()
	height := im.GetHeight()

	kernelMatrix := ImageMatrix{}

	// fmt.Println("origX, origY:", origX, origY, width, height)

	for i := (origX - size); i <= (origX + size); i++ {
		kernelY := 0
		column := make([]color.Color, (2*size + 1))

		if i < 0 || i >= width {
			kernelMatrix = append(kernelMatrix, column)
			continue
		}

		for j := (origY - size); j <= (origY + size); j++ {
			if j < 0 || j >= height {
				continue
			}

			column[kernelY] = im[i][j]
			kernelY++
		}

		kernelMatrix = append(kernelMatrix, column)
	}

	return kernelMatrix
}

//
// ApplyConvolution apply's a convolution matrix to the current image.
//
func (im ImageMatrix) ApplyConvolution(cm ConvolutionMatrix) ImageMatrix {
	return im.ApplyConvolutionFunction(cm, dontModifyConvolutionMatrixWeights)
}

//
// ApplyConvolutionWithSampleFunction apply's a weights to the convolution matrix (in addition to the weights
// in the matrix, based on the return values of the function)
//
func (im ImageMatrix) ApplyConvolutionWithSampleFunction(cm ConvolutionMatrix) ImageMatrix {
	return im.ApplyConvolutionFunction(cm, convolutionMatrixSampleFunction)
}

//
// ApplyConvolutionFunction apply's a convolution matrix to the current image, however the weight of the each
// entry in the convolution matrix is dependant on what your function (that you pass in) returns.
// This is so that you can return different weights for different conditions of the image.
// Please see the example, ApplyConvolutionWithRedInfluenceSampleFunction for more detail
//
// Please note that as the convolution matrix has weights itself, the result of the function will be multiplied by the
// weight in the convolution matrix to end up with the final weight that the pixel should have
//
func (im ImageMatrix) ApplyConvolutionFunction(cm ConvolutionMatrix, conFunc func(ImageMatrix, int, int, int, int, color.RGBA, float64) int) ImageMatrix {
	cmWidth := cm.GetWidth()
	cmHeight := cm.GetHeight()

	// Check to ensure that the convolutio matrix is a square and an odd number of rows/cols
	// This must be the case as we look up/down and left/right and equal amount from our current pixel
	// and we would not be able to have our pixel of interest in the absolute middle if the rows/cols
	// were not odd!
	if cmWidth != cmHeight {
		log.Fatalln("The convolution matrix passed in is not a square matrix!")
	} else if cmWidth%2 == 0 {
		log.Fatalln("The convolution matrix must be an odd number of rows/cols in size")
	}

	//
	//
	newMatrix := ImageMatrix{}
	imWidth := im.GetWidth()
	imHeight := im.GetHeight()
	cmSize := int(cmWidth / 2)

	//
	// for each row of the image...
	for x := 0; x < imWidth; x++ {
		column := make([]color.Color, imHeight)
		// for each column of the image...

		for y := 0; y < imHeight; y++ {
			// look at the current pixel so that we can use it's values as the initial values of the
			// new pixel in it's place
			applyConvolutionToPixel(im, x, y, cmSize, column, cm, conFunc)
		}

		// fmt.Println("Column:", column)
		newMatrix = append(newMatrix, column)

	}

	return newMatrix
}

//
//
//
func applyConvolutionToPixel(im ImageMatrix, x int, y int, cmSize int, column []color.Color, cm ConvolutionMatrix, conFunc func(ImageMatrix, int, int, int, int, color.RGBA, float64) int) {
	currentColour := im[x][y].(color.RGBA)
	redTotal := 0
	greenTotal := 0
	blueTotal := 0
	weight := 0

	kernelMatrix := im.GetKernelMatrix(x, y, cmSize)

	for i, column := range kernelMatrix {
		for j, colour := range column {
			if colour == nil {
				continue
			}

			//
			kernelPixelColour := colour.(color.RGBA)

			// get the distance of the current pixel compared to the centre of the kernel
			// the centre one is the one we are modifying and saving to a new image/matrix of course...
			distance := math.Sqrt(math.Pow(float64(cmSize-i), 2) + math.Pow(float64(cmSize-j), 2))

			// Call the function the user passed and get the return weight of how much influence
			// it should have over the centre pixel we want to change
			// We are multipling it by the weight in the convolution matrix as that way you can
			// control an aspect of the weight through the matrix as well (as well as the function that
			// we pass in of course :)
			cmValue := conFunc(im, x, y, i, j, kernelPixelColour, distance) * int(cm[i][j])

			// apply the influence / weight ... (eg. if cmValue was 0, then the current pixel would have
			// no influence over the pixel we are changing, if it was large in comparision to what we return
			// for the other kernel pixels, then it will have a large influence)
			redTotal += int(kernelPixelColour.R) * cmValue
			greenTotal += int(kernelPixelColour.G) * cmValue
			blueTotal += int(kernelPixelColour.B) * cmValue
			weight += cmValue
		}
	}

	//
	if weight == 0 {
		weight = 1
	}

	//
	newRedValue := uint8(redTotal / weight)
	newGreenValue := uint8(greenTotal / weight)
	newBlueValue := uint8(blueTotal / weight)

	column[y] = color.RGBA{newRedValue, newGreenValue, newBlueValue, currentColour.A}
	// fmt.Printf("[%v,%v] %v => %v\n", x, y, currentColour, column[y])
}

//
// GetWidth returns the height of the image
//
func (im ImageMatrix) GetWidth() int {
	return len(im)
}

//
// GetHeight returns the height of the image
//
func (im ImageMatrix) GetHeight() int {
	return len(im[0])
}
