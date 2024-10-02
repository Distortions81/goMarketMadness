package main

import (
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	fontScale             = 2
	fontSizeX, fontSizeY  = 16, 16
	termWidth, termHeight = 32, 24

	xMarginPercent = 40
	yMarginPercent = 40

	baseX = (fontSizeX / fontScale) * termWidth
	baseY = (fontSizeY / fontScale) * (termHeight - 1)

	xMargin = (baseX / xMarginPercent)
	yMargin = (baseY / yMarginPercent)

	screenWidth  = (baseX + xMargin) * fontScale
	screenHeight = (baseY + yMargin) * fontScale

	screenScale = 2
)

var (
	colorBG = color.NRGBA{R: 65, G: 232, B: 240, A: 255}
	colorFG = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
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
	screen.Fill(colorBG)

	outputLock.Lock()
	defer outputLock.Unlock()

	inbuf := strings.Join(sOut, "")
	lines := strings.Split(inbuf, "\n")

	sLine := len(lines) - termHeight
	if sLine < 0 {
		sLine = 0
	}
	startLine := max(0, sLine)
	showLines := lines[startLine:]
	buf := strings.Join(showLines, "\n")

	//ebitenutil.DebugPrint(screen, buf+sInBuf)
	drawText(screen, buf, xMargin/2, yMargin/2)
}

func drawText(screen *ebiten.Image, buf string, x, y int) {

	var row, col int
	for _, char := range buf {
		col++

		if char == '\n' {
			row++
			col = 0
		} else if char > '~' || char < ' ' {
			char = '?'
		}

		if col > termWidth {
			continue
		}

		start := int(char - 32)
		cx, cy := (start%32)*fontSizeX, (start/32)*fontSizeY

		// Define the rectangle for the sub-region
		rect := image.Rect(cx, cy, cx+fontSizeX, cy+fontSizeY)

		// Use SubImage and type assert the result to *ebiten.Image
		subImage := fontImg.SubImage(rect).(*ebiten.Image)

		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterNearest
		op.GeoM.Translate(float64(x)+float64(col*fontSizeX)-fontSizeX, float64(y)+float64(row*fontSizeY))
		screen.DrawImage(subImage, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func startEbiten() {
	g := &Game{}

	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowSize(screenWidth*screenScale, screenHeight*screenScale)
	ebiten.SetWindowTitle("Market Madness")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
