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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/spf13/viper"
)

const (
	screenWidth  = 320
	screenHeight = 240
	sceenZoom    = 3
	tileSize     = 16
	windowTitle  = "Fantasy!"
	DefaultTPS   = 5
)

type Game struct {
	keys []ebiten.Key
}

var img *ebiten.Image
var player *ebiten.Image
var Tiles map[string]interface{}
var Maps *viper.Viper
var Player *viper.Viper

func main() {

	g := &Game{}

	ebiten.SetWindowSize(screenWidth*sceenZoom, screenHeight*sceenZoom)
	ebiten.SetWindowTitle(windowTitle)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Draw(screen *ebiten.Image) {

	layers := Maps.GetStringSlice("layers")

	for l, layer := range layers {
		rows := strings.Split(layer, "\n")
		for y, row := range rows {
			col := strings.Split(row, "")
			for x, char := range col {

				if char == "S" { // S is the maps spawn point
					Player.Set("x", x*tileSize)
					Player.Set("y", y*tileSize)
					char = " "
				}
				drawTile(screen, char, x, y, l)

			}
		}
	}

	drawSprite(screen, Player)
	ebitenutil.DebugPrint(screen, fmt.Sprintf(" TPS: %0.2f", ebiten.CurrentTPS()))

	keyStrs := []string{}
	for _, p := range g.keys {
		keyStrs = append(keyStrs, p.String())
	}
	ebitenutil.DebugPrint(screen, "\n "+strings.Join(keyStrs, ", "))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	g.keys = inpututil.PressedKeys()
	return nil
}

func init() {

	loadMap("main")
	loadPlayer()

}

func drawSprite(screen *ebiten.Image, sprite *viper.Viper) {

	sprite.Set("isSpawned", true)

	var step = sprite.GetString("step")
	var dir = sprite.GetString("direction")

	var steps = sprite.GetStringMap(dir)

	if steps[step] == nil {
		step = "0"
	}

	var coordstring = steps[step]

	cors := strings.Split(coordstring.(string), ",")

	_ = cors

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(sprite.GetInt("x")), float64(sprite.GetInt("y")))

	x0, _ := strconv.Atoi(cors[0])
	y0, _ := strconv.Atoi(cors[1])
	x1, _ := strconv.Atoi(cors[2])
	y1, _ := strconv.Atoi(cors[3])

	screen.DrawImage(player.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image), op)

	stepInt, _ := strconv.Atoi(step)

	stepInt++
	sprite.Set("step", stepInt)

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
	}
}

func loadMap(mapName string) {

	Maps = viper.New()

	Maps.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	Maps.AddConfigPath("maps")
	Maps.SetConfigName(mapName)

	err := Maps.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	img, _, _ = ebitenutil.NewImageFromFile(Maps.GetString("tileFile"))
	Tiles = Maps.GetStringMap("tiles")

}

func loadPlayer() {
	Player = viper.New()
	Player.SetConfigFile("sprites/player.yaml")
	Player.ReadInConfig()
	Player.Set("isSpawned", false)
	Player.Set("direction", "l")
	Player.Set("step", 0)
	player, _, _ = ebitenutil.NewImageFromFile(Player.GetString("spritesheet"))

}
