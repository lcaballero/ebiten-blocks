package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const keyResolution = time.Second / 10

type KBHandler struct {
	handler    chan ebiten.Key
	right      *KeyHandler
	left       *KeyHandler
	down       *KeyHandler
	rotate     *KeyHandler
	resetPeice *KeyHandler
	quit       *KeyHandler
	pause      *KeyHandler
	playJab    *KeyHandler
	restart    *KeyHandler
}

func NewKBHandler() *KBHandler {
	out := make(chan ebiten.Key, 1)
	res := keyResolution
	return &KBHandler{
		handler:    out,
		left:       NewKeyHandler(ebiten.KeyJ, res, out),
		right:      NewKeyHandler(ebiten.KeyL, res, out),
		down:       NewKeyHandler(ebiten.KeyK, res, out),
		rotate:     NewKeyHandler(ebiten.KeySpace, res, out),
		resetPeice: NewKeyHandler(ebiten.KeyR, res, out),
		quit:       NewKeyHandler(ebiten.KeyQ, res, out),
		pause:      NewKeyHandler(ebiten.KeyP, res, out),
		playJab:    NewKeyHandler(ebiten.Key1, res, out),
		restart:    NewKeyHandler(ebiten.Key0, res, out),
	}
}

func (h *KBHandler) Update(paused bool, elapsed time.Duration) {
	if !paused {
		h.right.Update(elapsed)
		h.left.Update(elapsed)
		h.down.Update(elapsed)
		h.rotate.Update(elapsed)
		h.resetPeice.Update(elapsed)
	}
	h.quit.Update(elapsed)
	h.pause.Update(elapsed)
	h.playJab.Update(elapsed)
	h.restart.Update(elapsed)
}

type KeyHandler struct {
	key        ebiten.Key
	resolution time.Duration
	last       time.Duration
	accum      time.Duration
	out        chan ebiten.Key
}

func NewKeyHandler(key ebiten.Key, res time.Duration, out chan ebiten.Key) *KeyHandler {
	return &KeyHandler{
		key:        key,
		resolution: res,
		out:        out,
	}
}

func (k *KeyHandler) pushKey(out chan ebiten.Key, key ebiten.Key) {
	isPressed := inpututil.IsKeyJustPressed(key)
	if !isPressed {
		return
	}
	select {
	case out <- key:
		k.last = k.accum
	default:
	}
}

func (k *KeyHandler) Update(elapsed time.Duration) {
	k.accum += elapsed
	ds := k.accum - k.last
	if ds > k.resolution {
		k.pushKey(k.out, k.key)
	}
}
