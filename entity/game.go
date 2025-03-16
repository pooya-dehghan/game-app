package entity

import "time"

type Game struct {
	ID          uint
	Category    string
	QuestionIDs []uint
	PlayerIDs   []uint
	WinnerId    uint
	StartTime   time.Time
}

type Player struct {
	ID      uint
	UserID  uint
	Score   uint
	GameID  uint
	Answers []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	PlayserID  uint
	QuestionID uint
	Choice     PossibleAnswerChoice
}
