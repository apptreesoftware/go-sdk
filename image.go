package apptree

import (
	"encoding/json"
)

type Image struct {
	Valid    bool
	ImageURL string `json:"imageURL"`
}

func (Image) ValueType() Type {
	return Type_Image
}

func (i Image) IsNull() bool {
	return !i.Valid
}

func NewImage(imageURL string) Image {
	return Image{ImageURL: imageURL, Valid: true}
}

func NullImage() Image {
	return Image{Valid: false}
}

func (i Image) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte(`null`), nil
	}
	return json.Marshal(&struct {
		ImageURL string `json:"imageURL"`
	}{
		ImageURL: i.ImageURL,
	})
}

type uImage Image

func (i *Image) UnmarshalText(b []byte) error {
	str := string(b)
	if len(str) == 0 {
		i.Valid = false
		return nil
	}
	var uItem uImage
	err := json.Unmarshal(b, &uItem)
	if err != nil {
		return err
	}
	i.ImageURL = uItem.ImageURL
	i.Valid = true
	return nil
}
