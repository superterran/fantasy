package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/spf13/viper"
)

const (
	screenWidth  = 240
	screenHeight = 240
	sceenZoom    = 3
	tileSize     = 16
	tileXNum     = 25
)

type Game struct {
	layers [][]int
}

var tilesImage *ebiten.Image

func main() {

	g := &Game{}

	ebiten.SetWindowSize(screenWidth*sceenZoom, screenHeight*sceenZoom)
	ebiten.SetWindowTitle(viper.GetString("name"))

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Draw(screen *ebiten.Image) {

	const xNum = screenWidth / tileSize
	const yNum = screenHeight / tileSize

	t := viper.GetStringSlice("layers")

	// fmt.Println(t)
	for _, v := range t {
		for _, l := range v {
			drawTile(screen, string(l))
		}
		// fmt.Printf("Index: %d, Value: %v\n", i, v)
	}

	// ebitenutil.DebugPrint(screen, "test")

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	return nil
}

func init() {
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	// viper.SetConfigName("config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func drawTile(screen *ebiten.Image, s string) {

	// fmt.Println(s)
	// ebitenutil.DebugPrint(screen, s)

	tile, _, _ := ebitenutil.NewImageFromFile("tiles.png")
	// op := &ebiten.DrawImageOptions{}

	screen.DrawImage(tile, nil)
}
