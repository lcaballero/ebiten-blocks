package main

import (
	"bytes"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const sampleRate = 48000

type Sound struct {
	ctx    *audio.Context
	player *audio.Player
	stream *wav.Stream
}

func (a *Sound) SetVolume(lvl float64) {
	a.SetVolume(lvl)
}

func (a *Sound) Volume() float64 {
	return a.player.Volume()
}

func (a *Sound) Play() error {
	// As audioPlayer has one stream and remembers the playing position,
	// rewinding is needed before playing when reusing audioPlayer.
	if err := a.player.Rewind(); err != nil {
		return err
	}
	a.player.Play()
	return nil
}

type Audio struct {
	jab *Sound
}

func MustLoadAudio() *Audio {
	bin, err := ioutil.ReadFile("./jab.wav")
	if err != nil {
		panic(err)
	}
	context := audio.NewContext(sampleRate)
	stream, err := wav.DecodeWithoutResampling(bytes.NewReader(bin))
	if err != nil {
		panic(err)
	}
	player, err := context.NewPlayer(stream)
	if err != nil {
		panic(err)
	}
	player.SetVolume(.5)
	return &Audio{
		jab: &Sound{
			ctx:    context,
			stream: stream,
			player: player,
		},
	}
}
