package main

import "github.com/gdamore/tcell"

type player struct {
	x     float64
	y     float64
	score int
	image *image
}

func createPlayer() *player {
	player := &player{
		x:     float64(playerOffset),
		y:     float64(screenHeight) / 2.0,
		image: nil,
	}
	img, width, height, err := loadImage(playerImgPath)
	if err != nil {
		return player
	}

	imgStruct := &image{
		width:  width,
		height: height,
		image:  img,
	}
	player.image = imgStruct
	return player
}

func (p *player) gravity() error {
	p.y += constGravity
	return nil
}

func (p *player) move() error {
	if jumpImpulse {
		jumpTime = impulseTime
		jumpImpulse = false
	}

	if jumpTime > 0 {
		p.y -= constJump
		jumpTime--
	}
	return nil
}

func (p *player) hit(b *board) bool {
	if p.y > 0 && int(p.y) < screenHeight {
		if b.hasPlayer.column[int(p.y)] == 1 {
			return true
		}

	}
	return false
}

// Renders player texture
func (p *player) renderPlayer(s tcell.Screen) {
	for x := 0; x < p.image.width; x++ {
		for y := 0; y < p.image.height; y++ {
			// Render only if color is not 0
			if p.image.image[x*p.image.height+y] != 0 {
				renderPixel(s, int(p.x)+x, int(p.y)+y, p.image.image[x*p.image.height+y])
			}
		}
	}
	s.Show()
}
