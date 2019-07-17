package main

import (
	"fmt"

	"github.com/gdamore/tcell"
)

type scoreboard struct {
	texture *image
}

func createScoreboard() *scoreboard {
	return &scoreboard{
		texture: loadImageStruct(scoreboardImgPath),
	}
}

func (b *scoreboard) renderBackground(s tcell.Screen) {
	for x := 0; x < screenWidth; x++ {
		for y := screenHeight - scoreboardHeight; y < screenHeight; y++ {
			colorLocation := (x%b.texture.width)*b.texture.height + (y % b.texture.height)
			renderPixel(s, x, y, b.texture.image[colorLocation])
		}
	}
	s.Show()
}
func (b *scoreboard) renderScore(s tcell.Screen, score int) {
	scoreText := fmt.Sprintf("Score: %d", score)

	bgColorLocation := (screenHeight - 2) % b.texture.height

	// puts(s, style, 0, screenHeight-1, scoreText)
	renderText(s, screenWidth/2, screenHeight-2, scoreText, b.texture.image[bgColorLocation], uint32(tcell.ColorWhiteSmoke))
	s.Show()
}
