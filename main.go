package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/spf13/viper"
)

const (
	screenWidth  = 320
	screenHeight = 240
	sceenZoom    = 3
	tileSize     = 16
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

	// const xNum = screenWidth / tileSize
	// const yNum = screenHeight / tileSize

	layers := viper.GetStringSlice("layers")

	for l, layer := range layers {
		rows := strings.Split(layer, "\n")
		for y, row := range rows {
			col := strings.Split(row, "")
			for x, char := range col {
				drawTile(screen, char, x, y, l)
			}

		}
	}

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
	viper.AddConfigPath("./world")
	// viper.SetConfigName("config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	img, _, _ = ebitenutil.NewImageFromFile("world/tiles.png")
	Tiles = viper.GetStringMap("tiles")

}

func drawTile(screen *ebiten.Image, s string, x int, y int, l int) {

	if Tiles[s] != nil {
		cors := strings.Split(Tiles[s].(string), ",")

		tilex, _ := strconv.Atoi(cors[0])
		tiley, _ := strconv.Atoi(cors[1])

		tilex = tilex * tileSize
		tiley = tiley * tileSize

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))

		if l != 1 || s != " " {
			screen.DrawImage(img.SubImage(image.Rect(tilex, tiley, tilex+tileSize, tiley+tileSize)).(*ebiten.Image), op)
		}
		// screen.DrawImage(img, nil)

	}

}

var img *ebiten.Image
var Tiles map[string]interface{}
