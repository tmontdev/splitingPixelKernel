package KernelPixel

import (
        "image"
        "image/png"
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

func GetSimilarityOfDataImages(imageFull image.Image){
        for x := 0; x <= imageFull.Bounds().Max.X; x++ {
                for y := 0; y <=imageFull.Bounds().Max.Y; y++ {
                        fmt.Println(imageFull.At(x, y))
                }
        }
}