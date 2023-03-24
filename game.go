package main

import (
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lcaballero/ebiten-01/rand"
	"github.com/lcaballero/ebiten-01/shapes"
)

const keyResolution = time.Second / 10

type Game struct {
	board      *Board
	background *Background
	keys       *KeyHandler

	pieces   *Pieces
	prev     time.Time
	last     time.Duration
	accum    time.Duration
	seconds  time.Duration
	frames   int
	keyAccum time.Duration
	paused   bool

	// game pieces
	current *Tetromino
	next    *Tetromino
	rnd     rand.Rnd

	audio *Audio
}

func NewGame(opts NewGameOpts) *Game {
	p := NewPieces()
	seed := Seed(opts.Seed())
	audio := MustLoadAudio()
	rnd := rand.NewRnd(seed)
	bg := NewBackground()
	game := &Game{
		background: bg,
		board:      NewBoard(bg.board),
		keys:       NewKeyHandler(),
		prev:       time.Now(),
		pieces:     p,
		rnd:        rnd,
		audio:      audio,
	}
	game.createStartPeice()
	game.createNextPeice()
	return game
}

func (b *Game) velocity() shapes.Vec {
	return b.background.scoring.Velocity()
}

func (b *Game) top() shapes.Vec {
	//return b.background.board.Pos.Add(b.background.board.Dim.VecX().Half())
	return shapes.Vec{20 + 40, 20}
}

func (b *Game) createStartPeice() {
	b.current = &Tetromino{
		image:    b.pieces.blocks.Pick(b.rnd),
		pos:      b.top(),
		tetro:    RandTetro(b.rnd),
		rot:      R1,
		velocity: b.velocity(),
		size:     10,
	}
}

func (b *Game) createNextPeice() {
	next := &Tetromino{
		image:    b.pieces.blocks.Pick(b.rnd),
		tetro:    RandTetro(b.rnd),
		rot:      R1,
		velocity: b.velocity(),
		size:     10,
	}
	score := shapes.Rect{
		Pos: shapes.Vec{170, 20},
		Dim: shapes.Vec{60, 60},
	}
	b.next = next.MoveCenterTo(score.Center())
}

func (b *Game) rotateInNextPeice() {
	b.next.pos = b.top()
	b.current = b.next
	b.createNextPeice()
}

func (b *Game) restart() {
	b.createStartPeice()
	b.createNextPeice()
	b.background.reset()
	b.board.reset()
}

func (b *Game) step(now time.Time, elapsed time.Duration) {
	b.prev = now
	b.last = elapsed
	b.accum += elapsed
	b.keyAccum += elapsed
	ds := b.accum - b.seconds
	for b.keyAccum > keyResolution {
		select {
		case key := <-b.keys.handler:
			switch key {
			case ebiten.KeyL:
				if b.board.CanGoRight(b.current) {
					b.current.MoveRight()
				}
			case ebiten.KeyJ:
				if b.board.CanGoLeft(b.current) {
					b.current.MoveLeft()
				}
			case ebiten.KeySpace:
				b.current.RotateRight()
			case ebiten.KeyR:
				b.current.pos = b.top()
				b.current.isFrozen = false
			case ebiten.KeyP:
				b.paused = !b.paused
			case ebiten.KeyK:
				b.current.Accelerate()
			case ebiten.Key1:
				b.audio.jab.Play()
			case ebiten.Key0:
				b.restart()
			}
		default:
			b.keyAccum -= keyResolution
		}
	}
	hasTics := ds > time.Second
	for ds > time.Second {
		b.seconds += time.Second
		ds -= time.Second
	}
	if hasTics {
		//log.Printf("fps: %d", b.frames)
		b.frames = 0
	}
}

func (b *Game) Update() error {
	b.step(time.Now(), time.Since(b.prev))
	ds := float64(b.last) / float64(time.Second)
	if !b.paused {
		b.current.Update(b.last, ds)
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(1)
	}
	b.keys.Update(b.paused)
	b.board.CheckBounds(b.current)
	if b.current.isFrozen {
		rows := b.board.ClearFullRows(b.current)
		b.background.scoring = b.background.scoring.Add(len(rows))
		b.rotateInNextPeice()
		b.audio.jab.Play()
	}
	return nil
}

func (b *Game) Draw(screen *ebiten.Image) {
	b.background.Draw(screen)
	b.board.Draw(screen)
	b.current.Draw(screen)
	b.next.Draw(screen)
	b.frames++
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
