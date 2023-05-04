package cligw

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

type SelectOneOpts struct{ PageLimit *int }

func (gateway *Gateway) SelectOne(question string, choices SelectChoiceList) (*SelectChoice, error) {
	return gateway.SelectOneWithOpts(question, choices, SelectOneOpts{})
}

func (gateway *Gateway) SelectOneWithOpts(
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

	err := survey.AskOne(prompt, rawResult, gateway.icons)
	if err != nil {
		return nil, err
	}

	results := choices.GetByLabels(*rawResult)

	if len(results) == 0 {
		return nil, errors.New("no result selected")
	}

	return &results[0], nil
}
