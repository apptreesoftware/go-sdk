package apptree

import "encoding/json"

type ListItem struct {
	Valid bool
	Id    string `json:"id"`
	Value string `json:"value"`
}

func (l ListItem) IsNull() bool {
	return !l.Valid
}

func (ListItem) ValueType() Type {
	return Type_ListItem
}

func NewListItem(value string) ListItem {
	return ListItem{Id: value, Value: value, Valid: true}
}

func NullListItem() ListItem {
	return ListItem{Valid: false}
}

func (l ListItem) MarshalText() ([]byte, error) {
	if !l.Valid {
		return []byte(`null`), nil
	}
	return json.Marshal(&struct {
		Value string `json:"value"`
		Id    string `json:"id"`
	}{
		Value: l.Value,
		Id:    l.Id,
	})
}

type uListItem ListItem

func (l *ListItem) UnmarshalText(b []byte) error {
	str := string(b)
	if len(str) == 0 {
		l.Valid = false
		return nil
	}
	var uItem uListItem
	err := json.Unmarshal(b, &uItem)
	if err != nil {
		return err
	}
	l.Value = uItem.Value
	l.Id = uItem.Id
	l.Valid = true
	return nil
}
