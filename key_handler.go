package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyHandler struct {
	handler chan ebiten.Key
}

func NewKeyHandler() *KeyHandler {
	return &KeyHandler{
		handler: make(chan ebiten.Key, 1),
	}
}

func (b *KeyHandler) pushKey(key ebiten.Key) {
	isPressed := inpututil.IsKeyJustPressed(key)
	if !isPressed {
		return
	}
	select {
	case b.handler <- key:
	default:
	}
}

func (h *KeyHandler) Update(paused bool) {
	if !paused {
		h.pushKey(ebiten.KeyJ)
		h.pushKey(ebiten.KeyL)
		h.pushKey(ebiten.KeySpace)
		h.pushKey(ebiten.KeyR)
		h.pushKey(ebiten.KeyK)
	}
	h.pushKey(ebiten.KeyQ)
	h.pushKey(ebiten.KeyP)
	h.pushKey(ebiten.Key1)
	h.pushKey(ebiten.Key0)
}
