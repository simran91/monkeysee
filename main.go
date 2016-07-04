package main

import "fmt"
import "./mimage"
import "./mods"
import "regexp"
import "./lib/util"

//
// Demo program to show the usage of the functions and some of the mods provided...
//
func main() {

	stringRegex := regexp.MustCompile(`\.`)

	// sourceFiles := []string{"samples/rgb.jpg", "samples/rgb.jpg", "samples/rgb.jpg"}
	sourceFiles := []string{"samples/rgb.png", "samples/rgb.gif"}

	for _, sourceImage := range sourceFiles {
		fmt.Println("In:", sourceImage)
		image := mimage.LoadImageFromFile(sourceImage)
		colourMatrix := image.ColourMatrix()

		/////////////////////////////////////////////////////////////////////////////////////////////////////
		// mod: SwapRGBtoGBR
		//
		{
			fmt.Println("Running mods.SwapRGBtoGBR")

			newColourMatrix := mods.SwapRGBtoGBR(colourMatrix)
			newImage := mimage.ColourMatrixToImage(newColourMatrix)

			// Save as PNG
			{
				destImage := string(stringRegex.ReplaceAll([]byte(sourceImage), []byte{'-'})) + "-to-png-mod-SwapRGBtoGBR-autogenerated.png"
				fmt.Println("Out:", destImage)
				util.SaveImageToFileAsPNG(destImage, newImage)
			}

			// Save as JPG
			{
				destImage := string(stringRegex.ReplaceAll([]byte(sourceImage), []byte{'-'})) + "-to-jpg-mod-SwapRGBtoGBR-autogenerated.jpg"
				fmt.Println("Out:", destImage)
				util.SaveImageToFileAsJPG(destImage, newImage)
			}

			// Save as GIF
			{
				destImage := string(stringRegex.ReplaceAll([]byte(sourceImage), []byte{'-'})) + "-to-gif-mod-SwapRGBtoGBR-autogenerated.gif"
				fmt.Println("Out:", destImage)
				util.SaveImageToFileAsGIF(destImage, newImage)
			}

			fmt.Println()
		}

		/////////////////////////////////////////////////////////////////////////////////////////////////////
		// mod: GreyscaleAverageWithTranslusence
		//
		{
			fmt.Println("Running mods.GreyscaleAverageWithTranslusence")

			newColourMatrix := mods.GreyscaleAverageWithTranslusence(colourMatrix)
			newImage := mimage.ColourMatrixToImage(newColourMatrix)

			// Save as PNG
			{
				destImage := string(stringRegex.ReplaceAll([]byte(sourceImage), []byte{'-'})) + "-to-png-mod-GreyscaleAverageWithTranslusence-autogenerated.png"
				fmt.Println("Out:", destImage)
				util.SaveImageToFileAsPNG(destImage, newImage)
			}

			// Save as JPG
			{
				destImage := string(stringRegex.ReplaceAll([]byte(sourceImage), []byte{'-'})) + "-to-jpg-mod-GreyscaleAverageWithTranslusence-autogenerated.jpg"
				fmt.Println("Out:", destImage)
				util.SaveImageToFileAsJPG(destImage, newImage)
			}

			// Save as GIF
			{
				destImage := string(stringRegex.ReplaceAll([]byte(sourceImage), []byte{'-'})) + "-to-gif-mod-GreyscaleAverageWithTranslusence-autogenerated.gif"
				fmt.Println("Out:", destImage)
				util.SaveImageToFileAsGIF(destImage, newImage)
			}

			fmt.Println()
		}

	}
}
