package models

//Qualifying randed game
type AnswerLog struct {
	TCom
	PlayerId    int64
	QuestionId   int64
	PvpId       string
	Answer      string
	Failed       int
	Right       int
	Hinted      bool
	FirstFail  int64
}

func NewAnswerLog(playerId, questionId int64) *AnswerLog {
	return &AnswerLog{
		PlayerId: playerId,
		QuestionId: questionId,
	}
}

func (t *AnswerLog) TableName() string {
	return "answer_logs"
}


func (t *AnswerLog) X(){

}