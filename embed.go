package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	//go:embed data
	f       embed.FS
	fontImg *ebiten.Image
)

func init() {
	file, err := f.Open("data/fonts/font.png")
	if err != nil {
		log.Fatal("Can't read ti font.")
	}
	fontImg, _, err = ebitenutil.NewImageFromReader(file)
	if err != nil {
		log.Fatal("Can't parse ti font image.")
	}
}
