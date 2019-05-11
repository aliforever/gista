package models

type StoryQuestionResponderInfos struct {
	QuestionId                 string      `json:"question_id"`
	Question                   string      `json:"question"`
	QuestionType               string      `json:"question_type"`
	BackgroundColor            string      `json:"background_color"`
	TextColor                  string      `json:"text_color"`
	Responders                 []Responder `json:"responders"`
	MaxId                      interface{} `json:"max_id"`
	MoreAvailable              bool        `json:"more_available"`
	QuestionResponseCount      int         `json:"question_response_count"`
	LatestQuestionResponseTime int         `json:"latest_question_response_time"`
}
