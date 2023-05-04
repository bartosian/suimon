package cligw

import "github.com/AlecAivazis/survey/v2"

type SelectManyOpts struct{}

func (gateway *Gateway) SelectMany(question string, choices SelectChoiceList) ([]SelectChoice, error) {
	return gateway.SelectManyWithOpts(question, choices, SelectManyOpts{})
}

func (gateway *Gateway) SelectManyWithOpts(question string, choices SelectChoiceList, _ SelectManyOpts) ([]SelectChoice, error) {
	rawResult := new([]string)
	labels := choices.Labels()
	prompt := &survey.MultiSelect{
		Message:  question,
		Options:  labels,
		PageSize: len(labels),
	}

	err := survey.AskOne(prompt, rawResult, gateway.icons)
	if err != nil {
		return nil, err
	}

	return choices.GetByLabels(*rawResult...), nil
}
