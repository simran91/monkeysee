package mimage

import (
	"image"
	"image/color"
)

//
// RGBAMatrixToImage converts our ImageMatrix to an image.Image so that we can then save it
// to a file (or call other functions/methods on it that the image package provides), etc...
//
func RGBAMatrixToImage(imageMatrix ImageMatrix) image.Image {
	// *****************************************************************************
	// *****************************************************************************
	// TODO: We are taking the width and height of the image from a sample taken
	// by the key "r" and the first row/col; this is fine of course, but we should
	// introduce error checking to ensure that the ImageMatrix passed in is a
	// true rectangle; if it is not, things will break without sensible errors!
	// *****************************************************************************
	// *****************************************************************************

	width := len(imageMatrix["r"])
	height := len(imageMatrix["r"][0])

	//
	// Create a new image.Image...
	//
	newImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	// *****************************************************************************
	// *****************************************************************************
	// TODO: I *think* we are getting some loss of colour because we are converting
	// to uint8 (from the original uint32) type... have to do this for now as
	// color.RGBA seems to only take uint8, but need to investigate this further...
	// *****************************************************************************
	// *****************************************************************************

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r := uint8(imageMatrix["r"][x][y])
			g := uint8(imageMatrix["g"][x][y])
			b := uint8(imageMatrix["b"][x][y])
			a := uint8(imageMatrix["a"][x][y])

			colour := color.RGBA{r, g, b, a}

			newImage.SetRGBA(x, y, colour)
		}
	}

	return newImage
}
