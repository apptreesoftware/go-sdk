package apptree

import (
	"encoding/json"
)

type Color struct {
	Valid bool `json:"-"`
	Red   int  `json:"r"`
	Green int  `json:"g"`
	Blue  int  `json:"b"`
	Alpha int  `json:"a"`
}

func (c Color) IsNull() bool {
	return !c.Valid
}

func (Color) ValueType() Type {
	return Type_Color
}

func NewColor(red, green, blue, alpha int) Color {
	return Color{
		Valid: true,
		Red:   red,
		Green: green,
		Blue:  blue,
		Alpha: alpha,
	}
}

func NullColor() Color {
	return Color{Valid: false}
}

type uColor Color

func (c Color) MarshalText() ([]byte, error) {
	if !c.Valid {
		return []byte{}, nil
	}
	tmp := uColor(c)
	b, err := json.Marshal(&tmp)
	if err != nil {
		return nil, err
	}
	return b, err
}

func (c *Color) UnmarshalText(b []byte) error {
	str := string(b)
	if len(str) == 0 {
		c.Valid = false
		return nil
	}
	var col uColor
	err := json.Unmarshal(b, &col)
	if err != nil {
		return err
	}
	c.Alpha = col.Alpha
	c.Blue = col.Blue
	c.Green = col.Green
	c.Red = col.Red
	c.Valid = true
	return nil
}
