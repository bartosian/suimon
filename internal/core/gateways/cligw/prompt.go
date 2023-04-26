package cligw

import "github.com/AlecAivazis/survey/v2"

type PromptOpts struct {
	Default *string
}

func (ucg *Gateway) Prompt(question string) (*string, error) {
	return ucg.PromptWithOpts(question, PromptOpts{})
}

func (ucg *Gateway) PromptWithOpts(question string, opts PromptOpts) (*string, error) {
	result := new(string)
	prompt := &survey.Input{
		Message: question,
	}

	if opts.Default != nil {
		prompt.Default = *opts.Default
	}

	err := survey.AskOne(prompt, result, survey.WithValidator(survey.Required), ucg.surveyIcons)

	return result, err
}
