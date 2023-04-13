package cligw

type (
	SelectChoice struct {
		Label string
		Value string
		Data  interface{}
	}

	SelectChoiceList []SelectChoice
)

func NewSelectChoice(label, value string, data interface{}) SelectChoice {
	return SelectChoice{
		Label: label,
		Value: value,
		Data:  data,
	}
}

func NewSimpleSelectChoiceList(vals ...string) SelectChoiceList {
	list := SelectChoiceList{}

	for _, val := range vals {
		list = append(list, SelectChoice{
			Label: val,
			Value: val,
		})
	}

	return list
}

func (sc *SelectChoiceList) Labels() (result []string) {
	for _, choice := range *sc {
		result = append(result, choice.Label)
	}

	return
}

func (sc *SelectChoiceList) GetByLabels(labels ...string) SelectChoiceList {
	list := SelectChoiceList{}

	for _, label := range labels {
		for _, choice := range *sc {
			if choice.Label == label {
				list = append(list, choice)
			}
		}
	}

	return list
}
