package entity

type Question struct {
	ID uint
	Text string
	PossibleAnswers []PossibleAnswer
	CorrectAnswer string
	Difficulty QuestionDifficulty
	Category string
}

type PossibleAnswer struct { 
	ID uint
	Answer string
	Choice PossibleAnswerChoice
}

type Answer struct {
  ID uint
	PlayerID uint
	QuestionID uint
}

// a behave on type and column on a table 

type PossibleAnswerChocie uint8

func (p PossibleAnswerChocie) IsValid() bool {
	if(p < 0 || p > PossibleAnswerD){
		return false
	}

	return true
}

const(
	PossibleAnswerA  PossibleAnswerChocie = iota + 1
	PossibleAnswerB 
	PossibleAnswerC
	PossibleAnswerD
)
 
type QuestionDifficulty uint8


const( 
	QuestionDifficultyEasy = iota + 1
	QuestionDifficultyMedium
 	QuestionDifficultyHard
)

func (q QuestionDifficulty) IsValid() bool {
	 if q >= QuestionDifficultyEasy && q <= QuestionDifficultyHard {
			return true
	 }
	 return false
}