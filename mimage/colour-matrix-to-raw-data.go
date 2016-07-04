package mimage

import (
	"image"
	"image/color"
)

//
// ColourMatrixToImage converts our ImageMatrix to an image.Image so that we can then save it
// to a file (or call other functions/methods on it that the image package provides), etc...
//
func ColourMatrixToImage(imageMatrix ImageMatrix) image.Image {
	// *****************************************************************************
	// *****************************************************************************
	// TODO: We are taking the width and height of the image from a sample taken.
	// The first row/col is fine of course, but we should
	// introduce error checking to ensure that the ImageMatrix passed in is a
	// true rectangle; if it is not, things will break without sensible errors!
	// *****************************************************************************
	// *****************************************************************************

	// *****************************************************************************
	// *****************************************************************************
	// TODO: At the moment we cannot deal with JPEG images as they are in a YCbCr
	// colour model. We only know how to deal with RGBA's for now... should introduce
	// working with other models in the future...
	// *****************************************************************************
	// *****************************************************************************

	width := len(imageMatrix)
	height := len(imageMatrix[0])

	//
	// Create a new image.Image...
	//
	newImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	// newImage := image.NewYCbCr(image.Rectangle{image.Point{0, 0}, image.Point{width, height}}, image.YCbCrSubsampleRatio440)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			colour := imageMatrix[x][y].(color.RGBA)
			newImage.SetRGBA(x, y, colour)
			// colour := imageMatrix[x][y].(color.YCbCr)
			// newImage.Set(x, y, colour)

		}
	}

	return newImage
}
