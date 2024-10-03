package main

import (
	"fmt"
	"image"
	"log"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 15
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

type ebitenGame struct {
	game *gameData
}

var scroll float64

func getScroll() int {
	return int(scroll)
}

func (g *ebitenGame) Update() error {
	_, sV := ebiten.Wheel()
	if scroll+sV > 0 {
		scroll += sV
	} else {
		scroll = 0
	}
	fmt.Printf("%v\n", getScroll())

	if !g.game.showCursor {
		return nil
	}

	cInputRune = ebiten.AppendInputChars(cInputRune[:0])
	consoleIn += string(cInputRune)

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		newInput <- consoleIn
		consoleIn = ""
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(consoleIn) >= 1 {
			consoleIn = consoleIn[:len(consoleIn)-1]
		}
	}

	return nil
}

func (g *ebitenGame) Draw(screen *ebiten.Image) {
	screen.Fill(colorBG)

	inbuf := strings.Join(consoleOut, "")
	lines := strings.Split(inbuf, "\n")

	numLines := len(lines)
	sLine := numLines - termHeight - getScroll()
	if sLine < 0 {
		sLine = 0
	}
	startLine := max(0, sLine)
	showLines := lines[startLine:numLines]
	buf := strings.Join(showLines, "\n")

	if buf != "" && g.game.showCursor {
		cur := " "
		if time.Now().UnixMilli()/500%2 == 0 {
			cur = string(rune(cursorChar))
		}
		drawText(screen, buf+" "+consoleIn+cur, xMargin/2, yMargin/2)
	} else {
		drawText(screen, buf, xMargin/2, yMargin/2)
	}
}

func drawText(screen *ebiten.Image, buf string, x, y int) {

	var row, col int
	for _, char := range buf {
		col++

		if char == '\n' {
			row++
			col = 0
		} else if int(char) > cursorChar || char < ' ' {
			char = '?'
		}

		start := int(char - 32)
		cx, cy := (start%32)*fontSizeX, (start/32)*fontSizeY

		rect := image.Rect(cx, cy, cx+fontSizeX, cy+fontSizeY)

		subImage := fontImg.SubImage(rect).(*ebiten.Image)

		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterNearest
		op.GeoM.Translate(float64(x)+float64(col*fontSizeX)-float64(fontSizeX),
			float64(y)+float64(row*(fontSizeY+1)))
		screen.DrawImage(subImage, op)
	}
}

func (g *ebitenGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func startEbiten(game *gameData) {
	g := &ebitenGame{game: game}

	//fmt.Printf("%v, %v\n", screenWidth/fontScale, screenHeight/fontScale)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowSize(screenWidth*screenMagnify, screenHeight*screenMagnify)
	ebiten.SetWindowTitle("Market Madness")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
