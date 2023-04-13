package cligw

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

type SelectOneOpts struct{ PageLimit *int }

func (ucg *Gateway) SelectOne(question string, choices SelectChoiceList) (*SelectChoice, error) {
	return ucg.SelectOneWithOpts(question, choices, SelectOneOpts{})
}

func (ucg *Gateway) SelectOneWithOpts(
	question string,
	choices SelectChoiceList,
	opts SelectOneOpts,
) (*SelectChoice, error) {
	rawResult := new(string)
	labels := choices.Labels()
	pageSize := len(labels)

	if opts.PageLimit != nil && (*opts.PageLimit) < pageSize {
		pageSize = *opts.PageLimit
	}

	prompt := &survey.Select{
		Message:  question,
		Options:  labels,
		PageSize: pageSize,
	}

	err := survey.AskOne(prompt, rawResult, ucg.surveyIcons)
	results := choices.GetByLabels(*rawResult)

	if len(results) == 0 {
		return nil, errors.New("no result selected")
	}

	result := results[0]

	return &result, err
}
