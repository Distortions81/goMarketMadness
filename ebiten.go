package main

import (
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 256 * 2
	screenHeight = 192 * 2
	screenLines  = 24
)

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
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

type Game struct {
}

func (g *Game) Update() error {
	inputLock.Lock()
	defer inputLock.Unlock()

	sRune = ebiten.AppendInputChars(sRune[:0])
	sInBuf += string(sRune)

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		sLine = sInBuf
		sInBuf = ""
		sDirty = true
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(sInBuf) >= 1 {
			sInBuf = sInBuf[:len(sInBuf)-1]
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	outputLock.Lock()
	defer outputLock.Unlock()

	inbuf := strings.Join(sOut, "")
	lines := strings.Split(inbuf, "\n")

	sLine := len(lines) - screenLines
	if sLine < 0 {
		sLine = 0
	}
	startLine := max(0, sLine)
	showLines := lines[startLine:]
	buf := strings.Join(showLines, "\n")

	ebitenutil.DebugPrint(screen, buf+sInBuf)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func startEbiten() {
	g := &Game{}

	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowTitle("Market Madness")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
