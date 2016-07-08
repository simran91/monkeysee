package mimage

import "image/color"

//
// ImageMatrix defines how we store our matrix of Colours...
//
type ImageMatrix [][]color.Color

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
func (im ImageMatrix) GetKernelMatrix(origX, origY, size int) ImageMatrix {

	width := len(im)
	height := len(im[0])

    kernelMatrix := ImageMatrix{}


	// fmt.Println("origX, origY:", origX, origY, width, height)

	for i := (origX - size); i <= (origX + size); i++ {
	    kernelY := 0
        column := make([]color.Color, (2*size+1))

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

