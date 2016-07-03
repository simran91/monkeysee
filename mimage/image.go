package mimage

import "image"

import _ "image/png"  // The data we are given might be a png file... so need to import image/png
import _ "image/jpeg" // The data we are given might be a jpg file... so need to import image/jpeg
import _ "image/gif"  // The data we are given might be a gif file... so need to import image/gif
import "strings"
import "../lib/util"

//
// MImage is our main struct which will have methods we can call on once instantiated...
//
type MImage struct {
	rawdata string
}

//
// ImageMatrix defines how we store our matrix of Red, Green, Blue and Alpha channel values...
//
type ImageMatrix map[string][][]uint32

//
// RGBAMatrix reads in the rawdata and returns an ImageMatrix
//
func (i *MImage) RGBAMatrix() ImageMatrix {
	reader := strings.NewReader(i.rawdata)
	src, _, err := image.Decode(reader)
	util.CheckError(err)

	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	rgba := ImageMatrix{}

	red := [][]uint32{}
	green := [][]uint32{}
	blue := [][]uint32{}
	alpha := [][]uint32{}

	for x := 0; x < width; x++ {
		redCol := make([]uint32, height)
		greenCol := make([]uint32, height)
		blueCol := make([]uint32, height)
		alphaCol := make([]uint32, height)

		for y := 0; y < height; y++ {
			colour := src.At(x, y)
			r, g, b, a := colour.RGBA()

			redCol[y] = r
			greenCol[y] = g
			blueCol[y] = b
			alphaCol[y] = a
		}

		red = append(red, redCol)
		green = append(green, greenCol)
		blue = append(blue, blueCol)
		alpha = append(alpha, alphaCol)
	}

	rgba["r"] = red
	rgba["g"] = green
	rgba["b"] = blue
	rgba["a"] = alpha

	// debugPrintMatrix(rgba)

	return rgba
}
