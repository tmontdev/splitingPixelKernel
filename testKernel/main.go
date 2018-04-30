package main

import "splitingpixel/KernelPixel"

//import "os"

func main() {
        //imagePath := os.Args[1]
        imagePath := "/home/cabox/workspace/src/splitingpixel/testFiles/full.png"
        imageData := KernelPixel.GetImageData(imagePath)
        KernelPixel.GetSimilarityOfDataImages(imageData)
        
}