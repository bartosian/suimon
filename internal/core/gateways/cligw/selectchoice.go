package cligw

type (
	SelectChoiceList []SelectChoice
	SelectChoice     struct {
		Data  any
		Label string
		Value string
	}
)

func NewSelectChoiceList(values ...string) SelectChoiceList {
	options := make(SelectChoiceList, 0, len(values))

	for _, val := range values {
		option := SelectChoice{
			Label: val,
			Value: val,
		}

		options = append(options, option)
	}

	return options
}

func (choiceList *SelectChoiceList) Labels() (result []string) {
	for _, option := range *choiceList {
		result = append(result, option.Label)
	}

	return
}

func (choiceList *SelectChoiceList) GetByLabels(labels ...string) SelectChoiceList {
	options := make(SelectChoiceList, 0)

	for _, label := range labels {
		for _, option := range *choiceList {
			if option.Label == label {
				options = append(options, option)
			}
		}
	}

	return options
}
