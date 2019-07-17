package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	// Read image from file that already exists
	existingImageFile, err := os.Open("../assets/player/1.png")
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}
	fmt.Println(imageData)
	fmt.Println(imageType)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}

	fmt.Println(loadedImage)
	fmt.Println(loadedImage.At(6, 6))

	// c := colorful.Color{0.313725, 0.478431, 0.721569}
	c, err := colorful.Hex("#517AB8")
	if err != nil {
		log.Fatal(err)
	}
	c = colorful.Hsv(216.0, 0.56, 0.722)
	c = colorful.Xyz(0.189165, 0.190837, 0.480248)
	c = colorful.Xyy(0.219895, 0.221839, 0.190837)
	c = colorful.Lab(0.507850, 0.040585, -0.370945)
	c = colorful.Luv(0.507849, -0.194172, -0.567924)
	c = colorful.Hcl(276.2440, 0.373160, 0.507849)
	fmt.Printf("RGB values: %v, %v, %v", c.R, c.G, c.B)

	// rgb := NewRGBA(loadedImage)

	// c.

	bounds := loadedImage.Bounds()

	fmt.Println(bounds)

	// b := img.Bounds()
	imgWidth := bounds.Max.X + 1
	imgHeight := bounds.Max.Y + 1

	// bounds := loadedImage.Bounds()

	fmt.Println(bounds)

	image := make([]int, imgWidth*imgHeight)

	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			// fmt.Print(loadedImage.At(x, y))
			r, g, b, a := loadedImage.At(x, y).RGBA()
			// fin_col := 255 << 24
			if a > 0 {
				fin_col := 255<<24 | r<<16 | g<<8 | b
				// fmt.Printf("%h", fin_col)
				fmt.Printf("Failing at x:%d, y:%d, imgW:%d, %h", x, y, imgWidth, fin_col)

				image[x*(imgHeight-1)+y] = fin_col

			} else {
				fmt.Print("0000000000000000000000")

			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
