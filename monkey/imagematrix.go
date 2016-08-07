package monkey

import "image/color"
import "log"
import "math"
import "fmt"

//
// Point is a particular pixel position in an image
//
type Point struct {
	x int
	y int
}

//
// Path is a row of points that are connected
//
type Path []Point

//
// ImageRow is a row of color.RGBA values in an imagematrix...
//
type ImageRow []color.RGBA

//
// ImageMatrix defines how we store our matrix of colours...
//
type ImageMatrix []ImageRow

// ApplyFunctionToEveryPixel applys the given function to every pixel in the image
// (the function is passed the current pixel colour)
//
func (im ImageMatrix) ApplyFunctionToEveryPixel(modFunc func(ImageMatrix, int, int) color.RGBA) {
	for x, column := range im {
		for y := range column {
			c := modFunc(im, x, y)
			im[x][y] = c
		}
	}
}

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
// If you ask for a size 1 kernel at position 0,0 (for example); you will still get a 9x9 matrix with
// zero valued entries at the spots where the image x,y are not valid
// eg. That would return:
//            [c1, c2, c3]
//            [c4, c5, c6]
//            [c7, c8, c9]
//     Where as c5 is the position (0,0); c1, c2, c3, c4, c7 don't make sense as they are outside the bounds
//     of the image, so they will be set to the default of color.RGBA{}
//
func (im ImageMatrix) GetKernelMatrix(origX, origY, size int) ImageMatrix {

	width := im.GetWidth()
	height := im.GetHeight()

	kernelMatrix := ImageMatrix{}

	// fmt.Println("origX, origY:", origX, origY, width, height)

	for i := (origX - size); i <= (origX + size); i++ {
		kernelY := 0
		column := make([]color.RGBA, (2*size + 1))

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
		column := make([]color.RGBA, imHeight)
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

//
// SeamCarveHorizontal will carve the imagematrix by 1 pixel horizontally
//
func (im ImageMatrix) SeamCarveHorizontal() ImageMatrix {
	height := im.GetHeight()
	width := im.GetWidth()

	depth := 5
	var seam Path

	var startingPathOptions []Path
	var paths []Path

	for j := 0; j < height - 1; j++ {
		paths = im.getPathsHorizontal(Path{Point{0,j}}, depth)
		path := im.getLowestEnergyPath(paths)
		startingPathOptions = append(startingPathOptions, path)
	}

	// // paths := im.getPathsHorizontal(Path{Point{0, 190}}, depth)
	// paths := im.getPathsHorizontal(Path{Point{0, 22}}, depth)

	// for _, p := range paths {
	// 	debug("Starting Path Options: ", p)
	// }
	//
	// // fmt.Println("Starting path options are: ", paths)
	// debug("===================================================================================================")

	startingPathOptions = []Path{im.getLowestEnergyPath(paths)}

	seam = im.getLowestEnergyPath(startingPathOptions)

	for len(seam) < width {

		if len(seam)+depth > width {
			depth = width - len(seam)
		}

		// debug("========================================", seam)

		paths := im.getPathsHorizontal(seam, depth)
		lowestEnergyPath := im.getLowestEnergyPath(paths)
		seam = lowestEnergyPath

	}

	//
	newImage := im

	for _, point := range seam {
		newImage[point.x][point.y] = color.RGBA{255, 0, 255, 255}
	}

	//
	//
	//
	return newImage
}

//
// GetEnergyOfPixel returns the "energy" of a pixel - that is, the more different it is from it's surrounding
// pixels, the higher it's energy
//
func (im ImageMatrix) GetEnergyOfPixel(x, y int) float64 {

	// fmt.Println("START", x, y)
	cc := im[x][y] // centre colour (the pixel we are trying to get the enery for)
	// fmt.Println("END")

	kernelMatrix := im.GetKernelMatrix(x, y, 1)
	numBins := 0 // the number of bins (pixels) we will be comparing our central pixel to...
	energy := 0.0

	for _, row := range kernelMatrix {
		for _, c := range row {
			// ignore points where the RGBA value is 0,0,0,0 as they are probably the ones out of the image
			// eg. they are the entries above and to the left of 0,0 (this happens because the kernelMatrix
			// is always a square, so it returns zero'd entries for pixles that don't exist)
			if c.R == 0 && c.G == 0 && c.B == 0 && c.A == 0 {
				continue
			}

			// We should be doing the weight according to the Alpha channel... the close to 0 it is, the
			// less the energy should be (as we want transparent pixels to not add much energy at all)
			numBins++
			energy += math.Abs(float64(cc.R - c.R))
			energy += math.Abs(float64(cc.G - c.G))
			energy += math.Abs(float64(cc.B - c.B))
			energy += math.Abs(float64(cc.A-c.A)) * 3
		}
	}

	energy = energy / float64(numBins)

	// log.Printf("Energy=%v, Bins=%v x=%v y=%v", energy, numBins, x, y)
	// debugPrintMatrix(kernelMatrix)
	return energy
}

func (im ImageMatrix) getPathsHorizontal(path Path, depth int) []Path {
	var paths []Path
	height := im.GetHeight()
	width := im.GetWidth()
	x := path[len(path)-1].x
	y := path[len(path)-1].y

	// debug("x, y", x, y)

	if depth == 1 {
		if x+1 < width {
			pathWithRightPixel := append(path, Point{x + 1, y})
			paths = append(paths, pathWithRightPixel)
		}

		if x+1 < width && y-1 >= 0 {
			pathWithdiagonalUpPixel := append(path, Point{x + 1, y - 1})
			paths = append(paths, pathWithdiagonalUpPixel)
		}

		if x+1 < width && y+1 < height {
			pathWithDiagonalDownPixel := append(path, Point{x + 1, y + 1})
			paths = append(paths, pathWithDiagonalDownPixel)
		}

		// } else if (depth > 1 && x+1 <= width) {
	} else if depth > 1 {
		if x+1 < width {
			pathWithRightPixel := append(path, Point{x + 1, y})
			rightPaths := im.getPathsHorizontal(pathWithRightPixel, depth-1)

			for _, rightPath := range rightPaths {
				paths = append(paths, rightPath)
			}
		}

		if x+1 < width && y-1 >= 0 {
			pathWithDiagonalUpPixel := append(path, Point{x + 1, y - 1})
			diagonalUpPaths := im.getPathsHorizontal(pathWithDiagonalUpPixel, depth-1)

			for _, diagonalUpPath := range diagonalUpPaths {
				paths = append(paths, diagonalUpPath)
			}
		}

		if x+1 < width && y+1 < height {
			pathWithDiagonalDownPixel := append(path, Point{x + 1, y + 1})
			diagonalDownPaths := im.getPathsHorizontal(pathWithDiagonalDownPixel, depth-1)

			for _, diagonalDownPath := range diagonalDownPaths {
				paths = append(paths, diagonalDownPath)
			}
		}
	}

	// fmt.Println("Returning path", path)

	return paths
}

func (im ImageMatrix) getEnergyForPath(path Path) float64 {
	pathEnergy := 0.0

	for _, point := range path {
		pathEnergy += im.GetEnergyOfPixel(point.x, point.y)
	}

	return pathEnergy
}

func (im ImageMatrix) getLowestEnergyPath(paths []Path) Path {

	lowestEnergyPath := paths[0]

	for _, path := range paths {
		pathEnergy := im.getEnergyForPath(path)
		// debug("energy for path", pathEnergy)
		// debug("path is ", path)
		// debug("energy was", pathEnergy)
		if pathEnergy < im.getEnergyForPath(lowestEnergyPath) {
			lowestEnergyPath = path
		}
	}

	return lowestEnergyPath

}

func debug(v ...interface{}) {
	fmt.Println(v...)
}
