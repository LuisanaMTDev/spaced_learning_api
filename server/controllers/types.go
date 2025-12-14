package controllers

type AddLessonRequest struct {
	Topic            string `json:"topic"`
	AmounOfCards     int64  `json:"amoun_of_cards "`
	RepetitionsDates string `json:"repetitions_dates"`
}
