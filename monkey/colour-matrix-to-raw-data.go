package monkey

import (
	"image"
	"image/color"
)

//
// ImageMatrixToImage converts our ImageMatrix to an image.Image so that we can then save it
// to a file (or call other functions/methods on it that the image package provides), etc...
//
func ImageMatrixToImage(imageMatrix ImageMatrix) image.Image {
	// *****************************************************************************
	// *****************************************************************************
	// TODO: We are taking the width and height of the image from a sample taken.
	// The first row/col is fine of course, but we should
	// introduce error checking to ensure that the ImageMatrix passed in is a
	// true rectangle; if it is not, things will break without sensible errors!
	// *****************************************************************************
	// *****************************************************************************

	width := len(imageMatrix)
	height := len(imageMatrix[0])

	//
	// Create a new image.Image...
	//
	newImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			value := imageMatrix[x][y]
			colour := color.RGBA{}

			if value == nil {
				colour = color.RGBA{0, 0, 0, 0}
			} else {
				colour = value.(color.RGBA)
			}

			newImage.SetRGBA(x, y, colour)
		}
	}

	return newImage
}
