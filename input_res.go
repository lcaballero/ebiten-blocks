package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputReduce struct {
	keys       chan ebiten.Key
	accum      time.Duration
	resolution time.Duration
}

func (r *InputReduce) Update(elapsed time.Duration, dt float64) {
	r.accum += elapsed
}
