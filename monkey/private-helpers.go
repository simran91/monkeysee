package monkey

import "image/color"
import "fmt"
import "math"

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
	if distance < 1 {
		return 1
	}

	if colour.R > 150 && colour.G < 100 && colour.B < 100 {
		return int(5 * distance)
	}

	return 5
}


//
// Easier to write debug("...") than have to import fmt.Println in every file i want to just print something from for debugging purposes...
//
func debug(v ...interface{}) {
    fmt.Println(v...)
}

//
// Handy to put in as a breakpoint in some code if i just want it to end (while debugging)...
//
func forcePanic() {
    panic("picknicking... :)")
}

//
// apply a convolution (based on a given convolution matrix) to a particular pixel
//
func applyConvolutionToPixel(im ImageMatrix, x int, y int, cmSize int, column []color.RGBA, cm ConvolutionMatrix, conFunc func(ImageMatrix, int, int, int, int, color.RGBA, float64) int) {
	currentColour := im[x][y]
	redTotal := 0
	greenTotal := 0
	blueTotal := 0
	weight := 0

	kernelMatrix := im.GetKernelMatrix(x, y, cmSize)

	for i, kernelColumn := range kernelMatrix {
		for j, kernelPixelColour := range kernelColumn {

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

	// If the convolution matrix normalised itself; aka, say it was something like:
	//                                                                             { 0 -1  0}
	//                                                                             {-1  4 -1}
	//                                                                             { 0 -1  0}
	// then adding the entries (4 + (-1) + (-1) + (-1) + (-1)) results in a zero, in which case we leave
	// the weights alone (by setting the weight to divide by to 1) (aka, the weights do not have an impact,
	// but the pixels with a weight more more or less than 0 still do have an impact on the pixel
	// we are changing of course)
	if weight == 0 {
		weight = 1
	}

	// Normalise the values (based on the weight (total's in the matrix))
	newRedValue := redTotal / weight
	newGreenValue := greenTotal / weight
	newBlueValue := blueTotal / weight

	// If the values are "out of range" (outside the colour range of 0-255) then set them to 0 (absence of that
	// colour) if they were negative or 255 (100% of that colour) if they were greater than the max allowed.
	if newRedValue < 0 {
		newRedValue = 0
	} else if newRedValue > 255 {
		newRedValue = 255
	}

	if newGreenValue < 0 {
		newGreenValue = 0
	} else if newGreenValue > 255 {
		newGreenValue = 255
	}

	if newBlueValue < 0 {
		newBlueValue = 0
	} else if newBlueValue > 255 {
		newBlueValue = 255
	}

	// Assign the new values to the pixel in the column 'column' at position y
	column[y] = color.RGBA{uint8(newRedValue), uint8(newGreenValue), uint8(newBlueValue), currentColour.A}
	// fmt.Printf("[%v,%v] %v => %v\n", x, y, currentColour, column[y])
}


//
// debugPrintData prints the RGBAMatrix to STDOUT...
//
func debugPrintMatrix(matrix ImageMatrix) {
    for x, rows := range matrix {
        for y, colour := range rows {
            fmt.Printf("x:%v y:%v colour:%v\n", x, y, colour)
        }
    }
}
