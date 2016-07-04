package mimage

import "image"
import _ "image/png"  // The data we are given might be a png file... so need to import image/png to have it's initialisation effects...
import _ "image/jpeg" // The data we are given might be a jpg file... so need to import image/jpeg to have it's initialisation effects...
import _ "image/gif"  // The data we are given might be a gif file... so need to import image/gif to have it's initialisation effects...
import "strings"
import "../lib/util"
import "image/color"

//
// MImage is our main struct which will have methods we can call on once instantiated...
//
type MImage struct {
	rawdata string
}

//
// ImageMatrix defines how we store our matrix of Colours...
//
type ImageMatrix [][]color.Color

//
// ColourMatrix reads in the rawdata and returns a ColourMatrix
//
func (i *MImage) ColourMatrix() ImageMatrix {
	reader := strings.NewReader(i.rawdata)
	src, _, err := image.Decode(reader)
	util.CheckError(err)

	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	colourMatrix := ImageMatrix{}

	for x := 0; x < width; x++ {
		column := make([]color.Color, height)

		for y := 0; y < height; y++ {
			column[y] = src.At(x, y)
		}

		colourMatrix = append(colourMatrix, column)
	}


	// debugPrintMatrix(colourMatrix)

	return colourMatrix
}
