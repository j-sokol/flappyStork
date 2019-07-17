package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

var style = tcell.StyleDefault
var row = 0

var screenWidth = 0
var screenHeight = 0
var quit chan struct{}

type world struct {
	name string
	age  int
}

var blockTexture = &image{}

func main() {

	s, e := screenInit()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	screenWidth, screenHeight = s.Size()

	board := createBoard()
	player := createPlayer()
	scoreBoard := createScoreboard()

	// Load block structure
	blockTexture = loadImageStruct(blockImgPath)
	board.generatedCount = 0

	board.generateStartupColumns(screenWidth)

	// Init first board ofset and fill board with blocks
	board.heightOffset = screenHeight / 2
	board.fillBoardWithBlocks()

	board.getPlayerPosition()

	quit = make(chan struct{})
	go pollEvents(s)

	s.Show()

	go func() {
		for {
			board.shift()
			player.move()
			player.gravity()

			// if player.hit(board) {
			// return
			// }
			board.drawScreen(s)
			// board.drawScreen(s, scoreBoard)
			player.renderPlayer(s)
			scoreBoard.renderBackground(s)
			scoreBoard.renderScore(s, board.score)

			time.Sleep(time.Millisecond * tickTime)
		}
	}()

	<-quit
	s.Fini()
}
