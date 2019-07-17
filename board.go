package main

import (
	"fmt"

	"github.com/gdamore/tcell"
)

type board struct {
	firstRendered  *boardColumn
	hasPlayer      *boardColumn
	lastGenerated  *boardColumn
	generatedCount int
	heightOffset   int
	score          int
}

func createBoard() *board {
	return &board{
		score: 0,
	}
}

func (b *board) generateStartupColumns(screenWidth int) error {
	col1 := make([]int, screenHeight)

	for index := 0; index < (screenWidth*2)/4; index++ {
		b.addCollumn(col1)
		b.generatedCount++
	}
	return nil
}

func (b *board) generateBlockColumns(height, width int) error {

	colSide := make([]int, screenHeight)
	fillCollumn(colSide, 1, 0, height-verticalDistance)
	fillCollumn(colSide, 1, height+verticalDistance, screenHeight)

	b.addCollumn(colSide)
	b.generatedCount++

	for index := 0; index < blockWidth; index++ {
		b.addCollumn(colSide)
		b.generatedCount++

	}
	b.addCollumn(colSide)
	b.generatedCount++

	b.generateStartupColumns(width)
	return nil

}
func (b *board) addCollumn(column []int) error {
	col := &boardColumn{
		column: column,
	}
	if b.firstRendered == nil {
		b.firstRendered = col
	} else {
		currentNode := b.lastGenerated
		currentNode.next = col
		col.previous = b.lastGenerated
	}
	b.lastGenerated = col
	return nil
}

func (b *board) printStruct() error {
	currentNode := b.firstRendered
	if currentNode == nil {
		fmt.Println("Board is empty.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

// Checks if board is full of blocks
func (b *board) fillBoardWithBlocks() {
	if screenWidth+5 > b.generatedCount {
		width, height := getNewBlockDistance(b.heightOffset)

		b.generateBlockColumns(height, width)
		b.heightOffset = height
	}
}

// Returns pointer to column where player sits
func (b *board) getPlayerPosition() *boardColumn {
	tmpCol := b.firstRendered
	for index := 0; index <= playerOffset; index++ {
		tmpCol = tmpCol.next
		if index == playerOffset {
			b.hasPlayer = tmpCol
		}
	}
	return tmpCol

}
func (b *board) shift() {
	b.fillBoardWithBlocks()
	b.firstRendered = b.firstRendered.next
	b.hasPlayer = b.hasPlayer.next
	b.generatedCount--
	b.score++
}

// Renders achieved score
func (b *board) renderScore(s tcell.Screen) {
	scoreText := fmt.Sprintf("Score: %d", b.score)
	puts(s, style, 0, screenWidth-len(scoreText)-10, scoreText)
	s.Show()
}

// Renders window blocks
func (b *board) drawScreen(s tcell.Screen) {
	// func (b *board) drawScreen(s tcell.Screen, sb *scoreboard) {
	screenWidth, screenHeight = s.Size()

	if screenWidth == 0 || screenHeight == 0 {
		return
	}
	currentCol := b.firstRendered
	if currentCol != nil {
		for x := 0; x < screenWidth; x++ {
			for y := 0; y < screenHeight-scoreboardHeight; y++ {
				// Render only of point contains block
				if len(currentCol.column) > y && currentCol.column[y] == 1 {
					// Compute color from texture
					colorLocation := ((b.score+x)%blockTexture.width)*blockTexture.height + (y % blockTexture.height)
					renderPixel(s, x, y, blockTexture.image[colorLocation])

				} else {
					// Otherwise put empty texture
					puts(s, style, x, y, " ")
				}
			}
			// Shift to next collumn
			if currentCol.next != nil {
				currentCol = currentCol.next
			}

		}
	}
	// for x := 0; x < screenWidth; x++ {
	// 	for y := 0; y < scoreboardHeight; y++ {
	// 		colorLocation := (x%sb.texture.width)*sb.texture.height + (y % sb.texture.height)
	// 		renderPixel(s, x, screenHeight-y, blockTexture.image[colorLocation])
	// 	}
	// }

	s.Show()
}
