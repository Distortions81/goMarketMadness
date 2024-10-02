package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const fontPath = "data/fonts/font.png"

var (
	//go:embed data
	f       embed.FS
	fontImg *ebiten.Image
)

func init() {
	file, err := f.Open(fontPath)
	if err != nil {
		log.Fatal("Could not read: " + fontPath)
	}
	fontImg, _, err = ebitenutil.NewImageFromReader(file)
	if err != nil {
		log.Fatal("Could not parse: " + fontPath)
	}
}
