package mimage

import "image/color"
import "log"

//
// ConvolutionMatrix defines how we store our convolution matrices... 
//
type ConvolutionMatrix [][]uint8

//
// ApplyConvolution apply's a convolution matrix to the current image.
//
func (im ImageMatrix) ApplyConvolution(cm ConvolutionMatrix) ImageMatrix {
	cmWidth := len(cm)
	cmHeight := len(cm[0])

	// Check to ensure that the convolutio matrix is a square and an odd number of rows/cols
	// This must be the case as we look up/down and left/right and equal amount from our current pixel
	// and we would not be able to have our pixel of interest in the absolute middle if the rows/cols
	// were not odd!
	if (cmWidth != cmHeight) {
		log.Fatalln("The convolution matrix passed in is not a square matrix!")
	} else if (cmWidth % 2 == 0) {
		log.Fatalln("The convolution matrix must be an odd number of rows/cols in size")
	}

	//
	//
	newMatrix := ImageMatrix{}
	imWidth := len(im)
	imHeight := len(im[0])
	cmSize := int(cmWidth / 2)

	//
	// for each row of the image...
	for x := 0; x < imWidth; x++ {
		column := make([]color.Color, imHeight)
		// for each column of the image...
		for y := 0; y < imHeight; y++ {
			// look at the current pixel so that we can use it's values as the initial values of the
			// new pixel in it's place
			currentColour := im[x][y].(color.RGBA)
			redTotal := 0 // int(currentColour.R)
			greenTotal := 0 // int(currentColour.G)
			blueTotal := 0 // int(currentColour.B)
			weight := 0

			kernelMatrix := im.GetKernelMatrix(x, y, cmSize) // size hardcoded!!! need to apply right size...

			for i, column := range kernelMatrix {
				for j, colour := range column {
					if (colour == nil) {
						continue
					}

					cmValue := int(cm[i][j])
					c := colour.(color.RGBA)

					redTotal += int(c.R) * cmValue
					greenTotal += int(c.G) * cmValue
					blueTotal += int(c.B) * cmValue
					weight += cmValue
				}
			}

			//
			if (weight == 0) {
				weight = 1
			}

			//
			newRedValue := uint8(redTotal / weight)
			newGreenValue := uint8(greenTotal / weight)
			newBlueValue := uint8(blueTotal / weight)

			column[y] = color.RGBA{newRedValue, newGreenValue, newBlueValue, currentColour.A}
			// fmt.Printf("[%v,%v] %v => %v : %v\n", x, y, currentColour, column[y], redTotal)
		}

		// fmt.Println("Column:", column)
		newMatrix = append(newMatrix, column)
	}

	return newMatrix
}
