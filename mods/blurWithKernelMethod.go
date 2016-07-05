package mods

import "../mimage"
import "image/color"

//
// BlurWithKernelMethod is a simpler version of "Blur" as we are using the helper method GetKernelMatrix
// to get the matrix...
//
func BlurWithKernelMethod(matrix mimage.ImageMatrix, blurAmount int) mimage.ImageMatrix {

	width := len(matrix)
	height := len(matrix[0])
	newMatrix := mimage.ImageMatrix{}

	// for each row of the image...
	for x := 0; x < width; x++ {
		column := make([]color.Color, height)
		// for each column of the image...
		for y := 0; y < height; y++ {

			// look at the current pixel so that we can use it's values as the initial values of the
			// new pixel in it's place
			currentColour := matrix[x][y].(color.RGBA)
			redTotal := int(currentColour.R)
			greenTotal := int(currentColour.G)
			blueTotal := int(currentColour.B)
			samples := 1

			kernelMatrix := matrix.GetKernelMatrix(x, y, blurAmount)

			for _, column := range kernelMatrix {
				for _, colour := range column {
					if (colour == nil) {
						continue
					}

					c := colour.(color.RGBA)
					redTotal += int(c.R)
					greenTotal += int(c.G)
					blueTotal += int(c.B)
					samples++
				}
			}

			newRedValue := uint8(redTotal / samples)
			newGreenValue := uint8(greenTotal / samples)
			newBlueValue := uint8(blueTotal / samples)

			column[y] = color.RGBA{newRedValue, newGreenValue, newBlueValue, currentColour.A}
			// fmt.Printf("[%v,%v] %v => %v : %v\n", x, y, currentColour, column[y], redTotal)
		}

		// fmt.Println("Column:", column)
		newMatrix = append(newMatrix, column)
	}

	return newMatrix
}
