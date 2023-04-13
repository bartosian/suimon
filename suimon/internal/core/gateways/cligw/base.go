package cligw

import "github.com/AlecAivazis/survey/v2"

type Gateway struct {
	surveyIcons survey.AskOpt
}

func NewGateway() *Gateway {
	surveyIcons := survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Text = "❔"
	})

	return &Gateway{
		surveyIcons: surveyIcons,
	}
}
