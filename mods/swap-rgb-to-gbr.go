package mods

import "../monkey"

//
// SwapRGBtoGBR is a mod that swaps the colours around... it's a very simple mod designed
// to show how we loop over the ImageMatrix and read the colour values...
//
// IMPORTANT NOTE (or you will send yourself crazy wondering why this function doesn't actually
// swap the colours around)....
// eg. Doing R,G,B = G,B,R as we do below, copies the values (brightness) of the colours into the new
//     fields... but the image when rendered is still interpreted as RGB...
//     Consider the following for example:
//       If a pixel in the original image was: R=255 G=10 B=20
//       Then in the new image it will be    : R=10  G=20 B=255 (**remember** the image when rendered
//       is always intepreted as RGB)
//
func SwapRGBtoGBR(matrix monkey.ImageMatrix) monkey.ImageMatrix {

	width := matrix.GetWidth()
	height := matrix.GetHeight()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			colour := matrix[x][y]
			colour.R, colour.G, colour.B = colour.G, colour.B, colour.R
			matrix[x][y] = colour
		}
	}

	return matrix
}
