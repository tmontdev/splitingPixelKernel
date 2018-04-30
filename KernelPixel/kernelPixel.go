package KernelPixel

import (
        "image"
        "image/png"
        "image/color"
        "os"
        "fmt"
)

type Point struct {
        x int
        y int
}

func checkErro(erro error) {
        if erro != nil {
                panic(erro)
        }
}

func GetImageData(imagePath string) image.Image {
        imageFile, erro := os.Open(imagePath)
        checkErro(erro)
        defer imageFile.Close()

        imageData, erro := png.Decode(imageFile)
        checkErro(erro)

        return imageData
}

func GetSimilarityOfDataImages(imageFull image.Image, imagePartial image.Image){
        
        for x := 0; x <= imageFull.Bounds().Max.X; x++ {
                for y := 0; y <=imageFull.Bounds().Max.Y; y++ {
                        fmt.Println(searchPixel(imageFull.At(x, y), imagePartial))
                }
        }
}


func searchPixel(pixel color.Color, imageToCompare image.Image) int{
        var equalPixels int
        for x := 0; x <= imageToCompare.Bounds().Max.X; x++ {
                for y := 0; y <=imageToCompare.Bounds().Max.Y; y++ {
                        if imageToCompare.At(x, y) == pixel {
                                equalPixels++
                        }
                }
        }
        return equalPixels
}