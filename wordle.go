package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	title  string = "Wordle"
	width  int    = 435
	height int    = 600
	rows   int    = 6
	cols   int    = 5
)

var (
	fontSize        int = 24
	mplusNormalFont font.Face

	bkg       = color.White
	lightgrey = color.RGBA{R: 0x99, G: 0x99, B: 0x99, A: 0xff}
	grey      = color.RGBA{R: 0x66, G: 0x66, B: 0x66, A: 0xff}
	yellow    = color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
	green     = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
	fontColor = color.Black

	edge = false

	available_chars = "abcdefghijklmnopqrstuvwxyz"
	grid            [cols * rows]string
	dict            []string
	check           [cols * rows]int
	loc             int = 0
	won                 = false
	answer          string
)

type Game struct {
	runes []rune
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay = 30
		interval = 3
	)

	d := inpututil.KeyPressDuration(key)
	if d ==1 {
		return true
	}
	return false
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bkg)
	for w := 0; w < cols; w++ {
		for h := 0; h < rows; h++ {
			rect := ebiten.NewImage(75,75)
			rect.Fill(lightgrey)
			fontColor = color.Black

			//converts the grid to a linear index and controls color
			if check[w+(h*cols)] != 0 {
				if check[w+(h*cols)] == 1 {rec.Fill(green)}
				if check[w+(h*cols)] == 2 {rec.Fill(yellow)}
				if check[w+(h*cols)] == 3 {rec.Fill(grey)}
				font.Color = color.White
			}

			//check current cursor location
			if w+(h*cols) == loc && check[w+(h*cols)] == 0 {
				rect.Fill(grey)
				fontColor = color.White
			}

			//draws the rectangle
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*85+10),float64(h*85+10))
			screen.DrawImage(rect,op)
			if check[w+(cols*h)] == 0 {
				rect2 := ebiten.NewImage(73,73)
				rect2.Fill(color.White)
				op2 := &ebiten.DrawImageOptions{}
				op2.GeoM.Translate(float64(w*85+11),float64(h*85+11))
				screen.DrawImage(rec2,op2)
			}
			if grid[w+(h*cols)] != "" {
				msg := fmt.Sprintf(string.ToUpper(grid[w+(h*cols)])) 
			}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int,int) {
	return width, height
}

func main() {

	game := &Game{}

	font, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	nplusNormalFont, err = opentype.NewFace(font, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     72,
		Hinting: font.HitingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	eachWord, err := ioutil.ReadFile("dict.txt")

	rand.Seed(time.Now().UnixNano())
	answer = dict[rand.Intn(len(dict))]
}
