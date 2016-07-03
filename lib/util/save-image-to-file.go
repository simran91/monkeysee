package util

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

//
// SaveImageToFileAsPNG will save an image to the filesystem as a png...
//
func SaveImageToFileAsPNG(filename string, image image.Image) {
	outfile, err := os.Create(filename)
	CheckError(err)
	defer outfile.Close()

	png.Encode(outfile, image)
}

//
// SaveImageToFileAsJPG will save an image to the filesystem as a jpg...
//
func SaveImageToFileAsJPG(filename string, image image.Image) {
	outfile, err := os.Create(filename)
	CheckError(err)
	defer outfile.Close()

	// jpeg.Encode(outfile, image, nil)
	jpeg.Encode(outfile, image, &jpeg.Options{jpeg.DefaultQuality})
}

//
// SaveImageToFileAsGIF will save an image to the filesystem as a gif...
//
func SaveImageToFileAsGIF(filename string, image image.Image) {
	outfile, err := os.Create(filename)
	CheckError(err)
	defer outfile.Close()

	gif.Encode(outfile, image, nil)
}
