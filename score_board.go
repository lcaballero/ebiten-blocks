package main

type ScoreBoard struct {
	Score int
	Lines int
	Level int
}

func (s ScoreBoard) Add(n int) ScoreBoard {
	more := s.Level * n
	return ScoreBoard{
		Score: s.Score + more,
		Lines: s.Lines + n,
		Level: s.Level,
	}
}
