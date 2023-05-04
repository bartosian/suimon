package cligw

import "github.com/AlecAivazis/survey/v2"

type MsgOpts struct {
	Indent int
}

type Gateway struct {
	icons survey.AskOpt
}

func NewGateway() *Gateway {
	icons := survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Text = "❔"
	})

	return &Gateway{
		icons: icons,
	}
}
