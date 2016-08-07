package monkey

import "image/color"
import "log"
import "math"

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
// *****************************************************************************
// *****************************************************************************
// TODO: We need to make this generic and do hozizontal and vertical of course, but for now we have just done
//       it this way to test it out...
//       Also, currently we colour the seam pink, but we should delete the seam from the image (carve it,
//       as "seam carving" implies)
//       Also, the depth is currently hardcoded, it should be possible to change this - but we are laving
//       it hardcoded for now because if we enter larger numbers, it can take a long long time to process...
// *****************************************************************************
// *****************************************************************************
//
func (im ImageMatrix) SeamCarveHorizontal() ImageMatrix {
	height := im.GetHeight()
	width := im.GetWidth()

	depth := 3
	var seam Path

	var startingPathOptions []Path
	var paths []Path

	// get the starting seam by going through all options that start in the first column (x=0) and getting
	// the best possible path to start with
	for j := 0; j < height; j++ {
		paths = im.GetPathsHorizontal(Path{Point{0, j}}, depth)
		path := im.GetLowestEnergyPath(paths)
		startingPathOptions = append(startingPathOptions, path)
	}

	seam = im.GetLowestEnergyPath(startingPathOptions)

	// while we have not got the end-to-end seam from left to right (till the width of the image), keep adding
	// to the seam...
	for len(seam) < width {
		if len(seam)+depth > width {
			depth = width - len(seam)
		}

		paths := im.GetPathsHorizontal(seam, depth)
		lowestEnergyPath := im.GetLowestEnergyPath(paths)
		seam = lowestEnergyPath
	}

	// generate a new image (we don't want to modify the original image)
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
// GetEnergyOfPoint returns the "energy" of a pixel - that is, the more different it is from it's surrounding
// pixels, the higher it's energy
//
func (im ImageMatrix) GetEnergyOfPoint(x, y int) float64 {
	cc := im[x][y] // centre colour (the pixel we are trying to get the enery for)

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
			energy += math.Abs(math.Abs(float64(cc.R) - float64(c.R)))
			energy += math.Abs(math.Abs(float64(cc.G) - float64(c.G)))
			energy += math.Abs(math.Abs(float64(cc.B) - float64(c.B)))
			energy += math.Abs(math.Abs(float64(cc.A)-float64(c.A))) * 3
		}

	}

	energy = energy / float64(numBins)
	// debug("Energy of point:", x, y, energy)

	// log.Printf("Energy=%v, Bins=%v x=%v y=%v", energy, numBins, x, y)
	// debugPrintMatrix(kernelMatrix)
	return energy
}

func (im ImageMatrix) GetPathsHorizontal(path Path, depth int) []Path {
	// debug("Need to find next path options for:", path)
	var paths []Path
	height := im.GetHeight()
	width := im.GetWidth()
	x := path[len(path)-1].x
	y := path[len(path)-1].y

	// debug("x, y", x, y)

	// debug("Path we were passed in is", path)

	if depth == 1 {
		if x+1 < width {
			pathWithRightPoint := make(Path, len(path))
			copy(pathWithRightPoint, path)
			pathWithRightPoint = append(pathWithRightPoint, Point{x + 1, y})
			paths = append(paths, pathWithRightPoint)
		}

		if x+1 < width && y-1 >= 0 {
			pathWithDiagonalUpPoint := make(Path, len(path))
			copy(pathWithDiagonalUpPoint, path)
			pathWithDiagonalUpPoint = append(pathWithDiagonalUpPoint, Point{x + 1, y - 1})
			paths = append(paths, pathWithDiagonalUpPoint)
		}

		if x+1 < width && y+1 < height {
			pathWithDiagonalDownPoint := make(Path, len(path))
			copy(pathWithDiagonalDownPoint, path)
			pathWithDiagonalDownPoint = append(pathWithDiagonalDownPoint, Point{x + 1, y + 1})
			paths = append(paths, pathWithDiagonalDownPoint)
		}
	} else if depth > 1 {
		if x+1 < width {
			pathWithRightPoint := make(Path, len(path))
			copy(pathWithRightPoint, path)
			pathWithRightPoint = append(pathWithRightPoint, Point{x + 1, y})
			rightPaths := im.GetPathsHorizontal(pathWithRightPoint, depth-1)

			for _, rightPath := range rightPaths {
				paths = append(paths, rightPath)
			}
		}

		if x+1 < width && y-1 >= 0 {
			pathWithDiagonalUpPoint := make(Path, len(path))
			copy(pathWithDiagonalUpPoint, path)
			pathWithDiagonalUpPoint = append(pathWithDiagonalUpPoint, Point{x + 1, y - 1})
			diagonalUpPaths := im.GetPathsHorizontal(pathWithDiagonalUpPoint, depth-1)

			for _, diagonalUpPath := range diagonalUpPaths {
				paths = append(paths, diagonalUpPath)
			}
		}

		if x+1 < width && y+1 < height {
			pathWithDiagonalDownPoint := make(Path, len(path))
			copy(pathWithDiagonalDownPoint, path)
			pathWithDiagonalDownPoint = append(pathWithDiagonalDownPoint, Point{x + 1, y + 1})
			diagonalDownPaths := im.GetPathsHorizontal(pathWithDiagonalDownPoint, depth-1)

			for _, diagonalDownPath := range diagonalDownPaths {
				paths = append(paths, diagonalDownPath)
			}
		}
	}

	return paths
}

func (im ImageMatrix) GetEnergyOfPath(path Path) float64 {
	pathEnergy := 0.0

	for _, point := range path {
		pathEnergy += im.GetEnergyOfPoint(point.x, point.y)
	}

	return pathEnergy
}

func (im ImageMatrix) GetLowestEnergyPath(paths []Path) Path {
	lowestEnergyPath := paths[0]
	lowestEnergy := im.GetEnergyOfPath(lowestEnergyPath)

	for _, path := range paths {
		pathEnergy := im.GetEnergyOfPath(path)

		if pathEnergy < lowestEnergy {
			lowestEnergyPath = path
			lowestEnergy = pathEnergy
		}

	}

	return lowestEnergyPath
}
