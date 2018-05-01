package KernelPixel

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type pixel struct {
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

func isPixielInPixels(pix pixel, pixels []pixel) bool {
	for _, pixelTime := range pixels {
		if pixelTime == pix {
			return true
		}
	}
	return false
}

func GetSimilarityOfDataImages(imageFull image.Image, imagePartial image.Image) {
	pixels := make([]pixel, 1)
	for x := 0; x <= imageFull.Bounds().Max.X; x++ {
		for y := 0; y <= imageFull.Bounds().Max.Y; y++ {
			if searchPixel(imageFull.At(x, y), imagePartial) {
				if isPixielInPixels(pixel{x, y - 1}, pixels) {
					pixels = append(pixels, pixel{x, y})
				}

			}
		}
	}
	fmt.Println(pixels)
	fmt.Println(len(pixels))
}

func searchPixel(pixel color.Color, imageToCompare image.Image) bool {
	for x := 0; x <= imageToCompare.Bounds().Max.X; x++ {
		for y := 0; y <= imageToCompare.Bounds().Max.Y; y++ {
			if imageToCompare.At(x, y) == pixel {
				return true
			}
		}
	}
	return false
}
