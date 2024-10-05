package main

import (
	"image"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	scroll          int
	screenDirty     bool
	screenDirtyLock sync.Mutex

	cursorState bool
)

const linesPerScroll = 4

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 10
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

func (g *ebitenGame) Update() error {
	var sV int
	_, s := ebiten.Wheel()
	if s != 0.0 {
		setScreenDirty(true)
		sV += int(s * 4)
		if scroll+sV > 0 {
			scroll += sV
		} else {
			scroll = 0
		}
	}

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyPageUp) {
		sV += linesPerScroll
		setScreenDirty(true)
	} else if repeatingKeyPressed(ebiten.KeyPageDown) {
		sV -= linesPerScroll
		setScreenDirty(true)
	}

	if !g.game.showCursor {
		return nil
	}

	cInputRune = ebiten.AppendInputChars(cInputRune[:0])
	if len(cInputRune) > 0 {
		consoleIn += string(cInputRune)
		setScreenDirty(true)
	}

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		newInput <- consoleIn
		consoleIn = ""
		setScreenDirty(true)
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(consoleIn) >= 1 {
			consoleIn = consoleIn[:len(consoleIn)-1]
			setScreenDirty(true)
		}
	}

	return nil
}

func (g *ebitenGame) Draw(screen *ebiten.Image) {

	if !getScreenDirty() {
		return
	}

	screen.Fill(colorBG)

	inbuf := strings.Join(consoleOut, "")
	lines := strings.Split(inbuf, "\n")

	numLines := len(lines)
	fLine := numLines - termHeight
	if scroll > fLine {
		scroll = fLine
	}
	sLine := fLine - scroll
	if sLine < 0 {
		sLine = 0
	}
	startLine := max(0, sLine)
	showLines := lines[startLine:numLines]
	buf := strings.Join(showLines, "\n")
	bLen := len(buf) - 1

	if buf != "" && g.game.showCursor {
		cur := " "
		if cursorState {
			cur = string(rune(cursorChar))
		}
		drawText(screen, buf[:bLen]+" "+consoleIn+cur, xMargin/2, yMargin/2)
	} else {
		drawText(screen, buf, xMargin/2, yMargin/2)
	}

	if g.game.showSplash {
		for i := 1; i < 17; i++ {
			vector.DrawFilledRect(screen, (float32(i) * 16.0), float32(yMargin), 16, 24, barColors[i-1], false)
		}
	}

	setScreenDirty(false)
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
	ebiten.SetScreenClearedEveryFrame(false)

	go blinkCursor(game)
	setScreenDirty(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

func setScreenDirty(set bool) {
	screenDirtyLock.Lock()
	screenDirty = set
	screenDirtyLock.Unlock()
}

func getScreenDirty() bool {
	screenDirtyLock.Lock()
	defer screenDirtyLock.Unlock()

	return screenDirty
}

func blinkCursor(game *gameData) {
	for {
		time.Sleep(time.Millisecond * 333)
		if game.showCursor {
			cursorState = !cursorState
			setScreenDirty(true)
		}
	}
}
