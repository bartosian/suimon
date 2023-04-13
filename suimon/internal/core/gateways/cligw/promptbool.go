package cligw

import "github.com/AlecAivazis/survey/v2"

type PromptBoolOpts struct {
	Default bool
}

func (ucg *Gateway) PromptBool(question string) (bool, error) {
	return ucg.PromptBoolWithOpts(question, PromptBoolOpts{})
}

func (ucg *Gateway) PromptBoolWithOpts(question string, opts PromptBoolOpts) (bool, error) {
	result := new(bool)
	prompt := &survey.Confirm{
		Message: question,
		Default: opts.Default,
	}

	err := survey.AskOne(prompt, result, survey.WithValidator(survey.Required), ucg.surveyIcons)

	return *result, err
}
