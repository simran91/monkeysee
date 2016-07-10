package main

import "fmt"
import "./monkey"
import "./mods"
import "regexp"
import "./util"
import "path/filepath"
import "reflect"
import "runtime"
import "strings"
import "os"

//
// Demo program to show the usage of the functions and some of the mods provided...
//
func main() {

	sourceDir := "samples"
	destDir := "samples/autogenerated"
	sourceFiles := []string{}

	if (len(os.Args) > 1) {
		sourceFiles = os.Args[1:]
	} else {
		sourceFiles = []string{
                                "rgb.gif",
                                "rgb.jpg",
                                "rgb.png",
                                "flower.jpg",
                                "forest.png",
                                "gradient-rainbow.jpg",
                                "rgb-venn-diagram.png",
                                "sharp-leaf.jpg",
                                "water-on-leaf.jpg",
		}
	}

	// Compile the regex on where the '.' is so that we can parse filenames and save to new filenames later
	stringRegex := regexp.MustCompile(`\.`)

	for _, sourceFile := range sourceFiles {
		fmt.Println("************************************************************")
		fmt.Println("In:", filepath.Join(sourceDir, sourceFile))
		fmt.Println("************************************************************")

		image := monkey.LoadImageFromFile(filepath.Join(sourceDir, sourceFile))
		destFileInit := string(stringRegex.ReplaceAll([]byte(sourceFile), []byte{'-'}))

		runMod(modSwapRGBtoGBR, destDir, destFileInit, image.ImageMatrix())
		runMod(modGreyscaleAverageWithTranslusence, destDir, destFileInit, image.ImageMatrix())
		runMod(modBlur, destDir, destFileInit, image.ImageMatrix(), 8)
		runMod(modBlurWithKernelMethod, destDir, destFileInit, image.ImageMatrix(), 8)
		runMod(modGaussianBlur, destDir, destFileInit, image.ImageMatrix())
		runMod(modAverageBlur, destDir, destFileInit, image.ImageMatrix())
		runMod(modApplyConvolutionWithSampleFunction, destDir, destFileInit, image.ImageMatrix())
	}
}

// get the name of the funciton... will start with "main." as we are in the main package here!
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// call the correct mod function as required...
func runMod(modfunc func(string, string, monkey.ImageMatrix, ...interface{}) monkey.ImageMatrix,
	destDir, destFileInit string, imageMatrix monkey.ImageMatrix, vars ...interface{}) {
	// Get the function name so that we can use it in the directory/filenames we create...
	modName := getFunctionName(modfunc)
	modName = strings.Replace(modName, "main.mod", "", 1)

	// Inform the user on which function we are about to run...
	fmt.Printf("Running mod %v\n", modName)

	// Ensure the destination directory exists...
	destDir = filepath.Join(destDir, modName)
	err := os.MkdirAll(destDir, os.ModePerm)
	util.CheckError(err)

	// Call the actual mod func and create the new image...
	newImageMatrix := modfunc(destDir, destFileInit, imageMatrix, vars...)
	newImage := monkey.ImageMatrixToImage(newImageMatrix)

	// Save as PNG
	destImage := filepath.Join(destDir, destFileInit+"-to-png.png")
	fmt.Println("Out:", destImage)
	util.SaveImageToFileAsPNG(destImage, newImage)

	// Save as JPG
	destImage = filepath.Join(destDir, destFileInit+"-to-jpg.jpg")
	fmt.Println("Out:", destImage)
	util.SaveImageToFileAsJPG(destImage, newImage)

	// Save as GIF
	destImage = filepath.Join(destDir, destFileInit+"-to-gif.gif")
	fmt.Println("Out:", destImage)
	util.SaveImageToFileAsGIF(destImage, newImage)

	fmt.Println()
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: SwapRGBtoGBR
//
func modSwapRGBtoGBR(destDir, destFileInit string, imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.SwapRGBtoGBR(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: GreyscaleAverageWithTranslusence
//
func modGreyscaleAverageWithTranslusence(destDir, destFileInit string, imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.GreyscaleAverageWithTranslusence(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: Blur
//
func modBlur(destDir, destFileInit string, imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	blurAmount := vars[0].(int)
	newImageMatrix := mods.Blur(imageMatrix, blurAmount)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modBlurWithKernelMethod
//
func modBlurWithKernelMethod(destDir, destFileInit string, imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	blurAmount := vars[0].(int)
	newImageMatrix := mods.BlurWithKernelMethod(imageMatrix, blurAmount)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modGaussianBlur
//
func modGaussianBlur(destDir, destFileInit string, imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.GaussianBlur(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modAverageBlur
//
func modAverageBlur(destDir, destFileInit string, imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.AverageBlur(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyConvolutionWithSampleFunction
//
func modApplyConvolutionWithSampleFunction(destDir, destFileInit string, imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.ApplyConvolutionWithSampleFunction(imageMatrix)
	return newImageMatrix
}
