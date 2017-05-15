package apptree

type ListElement struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

func NewListItem(value string) ListElement {
	return ListElement{Id: value, Value: value}
}
