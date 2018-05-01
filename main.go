package main

import "splitingPixelKernel/KernelPixel"

//import "os"

func main() {
	//imagePath := os.Args[1]
	imagePathFull := "/home/jzes/lab/go/src/splitingPixelKernel/testFiles/full.png"
	imagePathPartial := "/home/jzes/lab/go/src/splitingPixelKernel/testFiles/partial.png"
	imageDataFull := KernelPixel.GetImageData(imagePathFull)
	imageDataPartial := KernelPixel.GetImageData(imagePathPartial)
	KernelPixel.GetSimilarityOfDataImages(imageDataFull, imageDataPartial)

}
