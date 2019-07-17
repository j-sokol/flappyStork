package main

import (
	"image/png"
	"os"
)

type image struct {
	width  int
	height int
	image  []uint32
	next   *image
}

func loadImage(path string) ([]uint32, int, int, error) {

	existingImageFile, err := os.Open(path)

	if err != nil {
		// Handle error
		return nil, 0, 0, nil
	}
	defer existingImageFile.Close()

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}

	// Get image boundaries
	bounds := loadedImage.Bounds()
	imgWidth := bounds.Max.X
	imgHeight := bounds.Max.Y

	image := make([]uint32, imgWidth*imgHeight)

	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			r, g, b, a := loadedImage.At(x, y).RGBA()
			// enter only if pixel has alpha value
			if a > 0 {
				/// RGB format in one int
				image[x*imgHeight+y] = 255<<24 | r<<16 | g<<8 | b
			}
		}
	}
	return image, imgWidth, imgHeight, nil
}

func loadImageStruct(path string) *image {
	img, width, height, err := loadImage(path)
	if err != nil {
		return nil
	}

	// imgStruct := &image{
	// 	width:  width,
	// 	height: height,
	// 	image:  img,
	// }
	// return imgStruct

	return &image{
		width:  width,
		height: height,
		image:  img,
	}
}
