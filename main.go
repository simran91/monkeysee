package main

import "fmt"
import "./monkey"
import "./mods"
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

	sourceDir := "."
	autogeneratedDir := filepath.Join(sourceDir, "autogenerated")
	sourceFiles := []string{}

	if len(os.Args) > 1 {
		sourceFiles = os.Args[1:]
	} else {
		fmt.Printf("Usage: monkeysee [files...]")
		os.Exit(1)
	}

	for _, sourceFile := range sourceFiles {
		fmt.Println("************************************************************")
		fmt.Println("In:", filepath.Join(sourceDir, sourceFile))
		fmt.Println("************************************************************")

		monkey := monkey.LoadImageFromFile(filepath.Join(sourceDir, sourceFile))
		destDir := filepath.Join(autogeneratedDir, sourceFile)

		runMod(modSwapRGBtoGBR, destDir, monkey.ImageMatrix())
		runMod(modGreyscaleAverageWithTranslusence, destDir, monkey.ImageMatrix())
		runMod(modBlur, destDir, monkey.ImageMatrix(), 8)
		runMod(modBlurWithKernelMethod, destDir, monkey.ImageMatrix(), 8)
		runMod(modGaussianBlur, destDir, monkey.ImageMatrix())
		runMod(modAverageBlur, destDir, monkey.ImageMatrix())
		runMod(modApplyConvolutionWithSampleFunction, destDir, monkey.ImageMatrix())
		runMod(modApplyFunctionToEveryPixelExample, destDir, monkey.ImageMatrix())
		runMod(modSharpen, destDir, monkey.ImageMatrix())
		runMod(modEdgeDetect, destDir, monkey.ImageMatrix())
		runMod(modEmboss, destDir, monkey.ImageMatrix())
		runMod(modIdentity, destDir, monkey.ImageMatrix())
		runMod(modPlayOne, destDir, monkey.ImageMatrix())
		runMod(modPlayTwo, destDir, monkey.ImageMatrix())
		runMod(modSeamCarveHorizontal, destDir, monkey.ImageMatrix())
	}
}

// get the name of the funciton... will start with "main." as we are in the main package here!
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// call the correct mod function as required...
func runMod(modfunc func(monkey.ImageMatrix, ...interface{}) monkey.ImageMatrix,
	destDir string, imageMatrix monkey.ImageMatrix, vars ...interface{}) {
	// Get the function name so that we can use it in the directory/filenames we create...
	modName := getFunctionName(modfunc)
	modName = strings.Replace(modName, "main.mod", "", 1)

	// Inform the user on which function we are about to run...
	fmt.Printf("Running mod %v\n", modName)

	// Ensure the destination directory exists...
	err := os.MkdirAll(destDir, os.ModePerm)
	util.CheckError(err)

	// Call the actual mod func and create the new image...
	newImageMatrix := modfunc(imageMatrix, vars...)
	newImage := monkey.ImageMatrixToImage(newImageMatrix)

	// Save as PNG
	destImage := filepath.Join(destDir, modName+".png")
	fmt.Println("Out:", destImage)
	util.SaveImageToFileAsPNG(destImage, newImage)

	// // Save as JPG
	// destImage = filepath.Join(destDir, modName+".jpg")
	// fmt.Println("Out:", destImage)
	// util.SaveImageToFileAsJPG(destImage, newImage)
	//
	// // Save as GIF
	// destImage = filepath.Join(destDir, modName+".gif")
	// fmt.Println("Out:", destImage)
	// util.SaveImageToFileAsGIF(destImage, newImage)

	fmt.Println()
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: SwapRGBtoGBR
//
func modSwapRGBtoGBR(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.SwapRGBtoGBR(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: GreyscaleAverageWithTranslusence
//
func modGreyscaleAverageWithTranslusence(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.GreyscaleAverageWithTranslusence(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: Blur
//
func modBlur(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	blurAmount := vars[0].(int)
	newImageMatrix := mods.Blur(imageMatrix, blurAmount)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modBlurWithKernelMethod
//
func modBlurWithKernelMethod(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	blurAmount := vars[0].(int)
	newImageMatrix := mods.BlurWithKernelMethod(imageMatrix, blurAmount)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modGaussianBlur
//
func modGaussianBlur(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.GaussianBlur(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modAverageBlur
//
func modAverageBlur(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.AverageBlur(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyConvolutionWithSampleFunction
//
func modApplyConvolutionWithSampleFunction(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.ApplyConvolutionWithSampleFunction(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyFunctionToEveryPixelExample
//
func modApplyFunctionToEveryPixelExample(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.ApplyFunctionToEveryPixelExample(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyFunctionToEveryPixelExample
//
func modSharpen(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.Sharpen(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyFunctionToEveryPixelExample
//
func modEdgeDetect(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.EdgeDetect(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyFunctionToEveryPixelExample
//
func modEmboss(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.Emboss(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyFunctionToEveryPixelExample
//
func modIdentity(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.Identity(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyFunctionToEveryPixelExample
//
func modPlayOne(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.PlayOne(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modApplyFunctionToEveryPixelExample
//
func modPlayTwo(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.PlayTwo(imageMatrix)
	return newImageMatrix
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// mod: modSeamCarveHorizontal
//
func modSeamCarveHorizontal(imageMatrix monkey.ImageMatrix, vars ...interface{}) monkey.ImageMatrix {
	newImageMatrix := mods.SeamCarveHorizontal(imageMatrix)
	return newImageMatrix
}
