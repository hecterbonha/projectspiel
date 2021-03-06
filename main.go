package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hecterbonha/projectspiel/lib/components/adventurer"
)

const (
	screenWidth  = 240
	screenHeight = 160

	frameOX     = 0
	frameWidth  = 32
	frameHeight = 32
)

var (
	count           = 0
	runnerImage     *ebiten.Image
	positionX       = screenWidth/2 - frameWidth/2
	positionY       = screenHeight/2 - frameHeight/2
	playerIsIdle    = true
	playerIsFlipped = false
	frameNum        = 5
	frameOY         = 0
	speed           = 1
)

func spriteChangeState() {
	if playerIsIdle {
		frameNum = 12
		frameOY = 0
	} else {
		frameNum = 8
		frameOY = 32
	}
}

func characterMovement() {
	var pressed []ebiten.Key
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			pressed = append(pressed, k)
		}
	}
	if len(pressed) == 0 {
		playerIsIdle = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerIsFlipped = false
		playerIsIdle = false
		positionX = positionX + speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerIsFlipped = true
		playerIsIdle = false
		positionX = positionX - speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerIsIdle = false
		positionY = positionY - speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerIsIdle = false
		positionY = positionY + speed
	}
	spriteChangeState()
}

func update(screen *ebiten.Image) error {
	count++
	characterMovement()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	if playerIsFlipped {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(frameWidth, 0)
	} else {
		op.GeoM.Scale(1, 1)
	}
	op.GeoM.Translate(float64(positionX), float64(positionY))
	i := (count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	hr := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)
	screen.DrawImage(runnerImage.SubImage(hr).(*ebiten.Image), op)
	return nil
}

func main() {
	fmt.Println("Welcome to Project Spiel")
	adventurer.Animation("WOW")
	img, _, err := image.Decode(bytes.NewReader(adventurer.Hero_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	title := "Project Spiel"
	if err := ebiten.Run(update, screenWidth, screenHeight, 4, title); err != nil {
		log.Fatal(err)
	}
}
