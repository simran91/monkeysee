package mods

import "../monkey"
import "image/color"

//
// Blur is a mod that blur's the image... it uses "kernel convolution". THe blurAmount determines the size
// of the kernel as we look at the pixels 'blurAmount' either side of the current pixel.
//
// *****************************************************************************
// *****************************************************************************
// TODO: This is a quick-and-dirty simple average blur implementation (inspired after watching the video
//       at https://www.youtube.com/watch?v=C_zFhWdM4ic (How Blurs & Filters Work - Computerphile))
//		 We still need to ensure we are handling the edges of the image properly in cases where the
//		 blurAmount is greater than 1
// *****************************************************************************
// *****************************************************************************
//
func Blur(matrix monkey.ImageMatrix, blurAmount int) monkey.ImageMatrix {

	width := len(matrix)
	height := len(matrix[0])
	newMatrix := monkey.ImageMatrix{}

	// for each row of the image...
	for x := 0; x < width; x++ {
		column := make([]color.RGBA, height)
		// for each column of the image...
		for y := 0; y < height; y++ {

			// look at the current pixel so that we can use it's values as the initial values of the
			// new pixel in it's place
			currentColour := matrix[x][y]
			redTotal := int(currentColour.R)
			greenTotal := int(currentColour.G)
			blueTotal := int(currentColour.B)
			samples := 1

			// Look on each side (left, right, above, below) of the current pixel based on the blurAmount
			// and set the pixel value to the average of all the pixels we looked at
			for i := (x - blurAmount); i <= (x + blurAmount); i++ {
				if i < 0 || i >= width {
					column[y] = color.RGBA{}
					continue
				}

				for j := (y - blurAmount); j <= (y + blurAmount); j++ {
					if j < 0 || j >= height {
						continue
					}

					// fmt.Println("x, y, i, j:", x, y, i, j)
					colour := matrix[i][j]
					redTotal += int(colour.R)
					greenTotal += int(colour.G)
					blueTotal += int(colour.B)
					samples++

					// fmt.Println("redTotal [%v,%v] %v\n", i, j)
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
