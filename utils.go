package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/mattn/go-runewidth"
)

func getNewBlockDistance(prevHeight int) (int, int) {
	widthOffset := rand.Intn(100) + 50
	heightOffset := rand.Intn(1)
	if heightOffset+verticalDistance > screenHeight || heightOffset-verticalDistance < 0 {
		heightOffset = -heightOffset
	}
	return widthOffset, prevHeight + heightOffset
}

func puts(s tcell.Screen, style tcell.Style, x, y int, str string) {
	i := 0
	var deferred []rune
	dwidth := 0
	zwj := false
	for _, r := range str {
		if r == '\u200d' {
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
			deferred = append(deferred, r)
			zwj = true
			continue
		}
		if zwj {
			deferred = append(deferred, r)
			zwj = false
			continue
		}
		switch runewidth.RuneWidth(r) {
		case 0:
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
		case 1:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 1
		case 2:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 2
		}
		deferred = append(deferred, r)
	}
	if len(deferred) != 0 {
		s.SetContent(x+i, y, deferred[0], deferred[1:], style)
		i += dwidth
	}
}

func putln(s tcell.Screen, str string) {

	puts(s, style, 1, row, str)
	row++
}

// Renders one pixel
func renderPixel(s tcell.Screen, x, y int, color uint32) {
	const gl = '▄'                         // Fill with this pixel
	st := tcell.StyleDefault               // Get default style
	st = st.Foreground(tcell.Color(color)) // Lower part of pixel

	st = st.Background(tcell.Color(color)) // Upper
	s.SetCell(x, y, st, gl)
}

func renderText(s tcell.Screen, x, y int, text string, bgColor, color uint32) {
	st := tcell.StyleDefault
	st = st.Foreground(tcell.Color(color))   // Lower part of pixel
	st = st.Background(tcell.Color(bgColor)) // Upper
	// for i := 0; i < len(text); i++ {

	// 	s.SetCell(x+i, y, st, text[i])
	// }
	puts(s, st, x, y, text)
}

// Get events from keyboard
func pollEvents(s tcell.Screen) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				close(quit)
				return

			case tcell.KeyRune:
				switch ev.Rune() {

				case 'q':
					close(quit)
					return

				case ' ':
					jumpImpulse = true
					// return
				}
				//s.Sync()
				// case tcell.KeyUp:
				// 	step := (vp.y0 - vp.y1) / 10
				// 	vp.y0 += step
				// 	vp.y1 += step
			}
		case *tcell.EventResize:
			s.Sync()
		}
	}
}

func screenInit() (tcell.Screen, error) {
	// Set tcell defaults
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	encoding.Register()

	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorSkyblue))
	s.EnableMouse()
	s.Clear()
	return s, nil
}
