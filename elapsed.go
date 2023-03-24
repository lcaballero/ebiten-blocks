package main

import (
	"fmt"
	"time"
)

var id int = 0

type Elapsed struct {
	id    int
	msg   string
	start time.Time
}

func Start(msg string) Elapsed {
	id++
	return Elapsed{
		id:    id,
		msg:   msg,
		start: time.Now(),
	}
}

func (e Elapsed) End() Elapsed {
	elapsed := time.Since(e.start)
	fmt.Printf("elapsed: %v, message: %s\n", elapsed, e.msg)
	return e
}
