package main

import "time"

const seed0 = int64(12231)

func Seed(seed int64) {
	if seed < 0 {
		seed = seed0
	}
	if seed == 0 {
		seed = time.Now().UnixMilli()
	}
}
