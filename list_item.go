package apptree

import "encoding/json"

type ListItem struct {
	Valid bool
	Id    string `json:"id"`
	Value string `json:"value"`
	Attribute01 string `json:"attribute01,omitempty"`
	Attribute02 string `json:"attribute02,omitempty"`
	Attribute03 string `json:"attribute03,omitempty"`
	Attribute04 string `json:"attribute04,omitempty"`
	Attribute05 string `json:"attribute05,omitempty"`
	Attribute06 string `json:"attribute06,omitempty"`
	Attribute07 string `json:"attribute07,omitempty"`
	Attribute08 string `json:"attribute08,omitempty"`
	Attribute09 string `json:"attribute09,omitempty"`
	Attribute10 string `json:"attribute10,omitempty"`
	Attribute11 string `json:"attribute11,omitempty"`
	Attribute12 string `json:"attribute12,omitempty"`
	Attribute13 string `json:"attribute13,omitempty"`
	Attribute14 string `json:"attribute14,omitempty"`
	Attribute15 string `json:"attribute15,omitempty"`
	Attribute16 string `json:"attribute16,omitempty"`
	Attribute17 string `json:"attribute17,omitempty"`
	Attribute18 string `json:"attribute18,omitempty"`
	Attribute19 string `json:"attribute19,omitempty"`
	Attribute20 string `json:"attribute20,omitempty"`
	Attribute21 string `json:"attribute21,omitempty"`
	Attribute22 string `json:"attribute22,omitempty"`
	Attribute23 string `json:"attribute23,omitempty"`
	Attribute24 string `json:"attribute24,omitempty"`
	Attribute25 string `json:"attribute25,omitempty"`
	Attribute26 string `json:"attribute26,omitempty"`
	Attribute27 string `json:"attribute27,omitempty"`
	Attribute28 string `json:"attribute28,omitempty"`
	Attribute29 string `json:"attribute29,omitempty"`
	Attribute30 string `json:"attribute30,omitempty"`
	Attribute31 string `json:"attribute31,omitempty"`
	Attribute32 string `json:"attribute32,omitempty"`
	Attribute33 string `json:"attribute33,omitempty"`
	Attribute34 string `json:"attribute34,omitempty"`
	Attribute35 string `json:"attribute35,omitempty"`
	Attribute36 string `json:"attribute36,omitempty"`
	Attribute37 string `json:"attribute37,omitempty"`
	Attribute38 string `json:"attribute38,omitempty"`
	Attribute39 string `json:"attribute39,omitempty"`
	Attribute40 string `json:"attribute40,omitempty"`
	Attribute41 string `json:"attribute41,omitempty"`
	Attribute42 string `json:"attribute42,omitempty"`
	Attribute43 string `json:"attribute43,omitempty"`
	Attribute44 string `json:"attribute44,omitempty"`
	Attribute45 string `json:"attribute45,omitempty"`
	Attribute46 string `json:"attribute46,omitempty"`
	Attribute47 string `json:"attribute47,omitempty"`
	Attribute48 string `json:"attribute48,omitempty"`
	Attribute49 string `json:"attribute49,omitempty"`
	Attribute50 string `json:"attribute50,omitempty"`
	Attribute51 string `json:"attribute51,omitempty"`
	Attribute52 string `json:"attribute52,omitempty"`
	Attribute53 string `json:"attribute53,omitempty"`
	Attribute54 string `json:"attribute54,omitempty"`
	Attribute55 string `json:"attribute55,omitempty"`
	Attribute56 string `json:"attribute56,omitempty"`
	Attribute57 string `json:"attribute57,omitempty"`
	Attribute58 string `json:"attribute58,omitempty"`
	Attribute59 string `json:"attribute59,omitempty"`
	Attribute60 string `json:"attribute60,omitempty"`
	Attribute61 string `json:"attribute61,omitempty"`
	Attribute62 string `json:"attribute62,omitempty"`
	Attribute63 string `json:"attribute63,omitempty"`
	Attribute64 string `json:"attribute64,omitempty"`
	Attribute65 string `json:"attribute65,omitempty"`
	Attribute66 string `json:"attribute66,omitempty"`
	Attribute67 string `json:"attribute67,omitempty"`
	Attribute68 string `json:"attribute68,omitempty"`
	Attribute69 string `json:"attribute69,omitempty"`
	Attribute70 string `json:"attribute70,omitempty"`
	Attribute71 string `json:"attribute71,omitempty"`
	Attribute72 string `json:"attribute72,omitempty"`
	Attribute73 string `json:"attribute73,omitempty"`
	Attribute74 string `json:"attribute74,omitempty"`
	Attribute75 string `json:"attribute75,omitempty"`
	Attribute76 string `json:"attribute76,omitempty"`
	Attribute77 string `json:"attribute77,omitempty"`
	Attribute78 string `json:"attribute78,omitempty"`
	Attribute79 string `json:"attribute79,omitempty"`
	Attribute80 string `json:"attribute80,omitempty"`
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
		Attribute01 string `json:"attribute01,omitempty"`
		Attribute02 string `json:"attribute02,omitempty"`
		Attribute03 string `json:"attribute03,omitempty"`
		Attribute04 string `json:"attribute04,omitempty"`
		Attribute05 string `json:"attribute05,omitempty"`
		Attribute06 string `json:"attribute06,omitempty"`
		Attribute07 string `json:"attribute07,omitempty"`
		Attribute08 string `json:"attribute08,omitempty"`
		Attribute09 string `json:"attribute09,omitempty"`
		Attribute10 string `json:"attribute10,omitempty"`
		Attribute11 string `json:"attribute11,omitempty"`
		Attribute12 string `json:"attribute12,omitempty"`
		Attribute13 string `json:"attribute13,omitempty"`
		Attribute14 string `json:"attribute14,omitempty"`
		Attribute15 string `json:"attribute15,omitempty"`
		Attribute16 string `json:"attribute16,omitempty"`
		Attribute17 string `json:"attribute17,omitempty"`
		Attribute18 string `json:"attribute18,omitempty"`
		Attribute19 string `json:"attribute19,omitempty"`
		Attribute20 string `json:"attribute20,omitempty"`
		Attribute21 string `json:"attribute21,omitempty"`
		Attribute22 string `json:"attribute22,omitempty"`
		Attribute23 string `json:"attribute23,omitempty"`
		Attribute24 string `json:"attribute24,omitempty"`
		Attribute25 string `json:"attribute25,omitempty"`
		Attribute26 string `json:"attribute26,omitempty"`
		Attribute27 string `json:"attribute27,omitempty"`
		Attribute28 string `json:"attribute28,omitempty"`
		Attribute29 string `json:"attribute29,omitempty"`
		Attribute30 string `json:"attribute30,omitempty"`
		Attribute31 string `json:"attribute31,omitempty"`
		Attribute32 string `json:"attribute32,omitempty"`
		Attribute33 string `json:"attribute33,omitempty"`
		Attribute34 string `json:"attribute34,omitempty"`
		Attribute35 string `json:"attribute35,omitempty"`
		Attribute36 string `json:"attribute36,omitempty"`
		Attribute37 string `json:"attribute37,omitempty"`
		Attribute38 string `json:"attribute38,omitempty"`
		Attribute39 string `json:"attribute39,omitempty"`
		Attribute40 string `json:"attribute40,omitempty"`
		Attribute41 string `json:"attribute41,omitempty"`
		Attribute42 string `json:"attribute42,omitempty"`
		Attribute43 string `json:"attribute43,omitempty"`
		Attribute44 string `json:"attribute44,omitempty"`
		Attribute45 string `json:"attribute45,omitempty"`
		Attribute46 string `json:"attribute46,omitempty"`
		Attribute47 string `json:"attribute47,omitempty"`
		Attribute48 string `json:"attribute48,omitempty"`
		Attribute49 string `json:"attribute49,omitempty"`
		Attribute50 string `json:"attribute50,omitempty"`
		Attribute51 string `json:"attribute51,omitempty"`
		Attribute52 string `json:"attribute52,omitempty"`
		Attribute53 string `json:"attribute53,omitempty"`
		Attribute54 string `json:"attribute54,omitempty"`
		Attribute55 string `json:"attribute55,omitempty"`
		Attribute56 string `json:"attribute56,omitempty"`
		Attribute57 string `json:"attribute57,omitempty"`
		Attribute58 string `json:"attribute58,omitempty"`
		Attribute59 string `json:"attribute59,omitempty"`
		Attribute60 string `json:"attribute60,omitempty"`
		Attribute61 string `json:"attribute61,omitempty"`
		Attribute62 string `json:"attribute62,omitempty"`
		Attribute63 string `json:"attribute63,omitempty"`
		Attribute64 string `json:"attribute64,omitempty"`
		Attribute65 string `json:"attribute65,omitempty"`
		Attribute66 string `json:"attribute66,omitempty"`
		Attribute67 string `json:"attribute67,omitempty"`
		Attribute68 string `json:"attribute68,omitempty"`
		Attribute69 string `json:"attribute69,omitempty"`
		Attribute70 string `json:"attribute70,omitempty"`
		Attribute71 string `json:"attribute71,omitempty"`
		Attribute72 string `json:"attribute72,omitempty"`
		Attribute73 string `json:"attribute73,omitempty"`
		Attribute74 string `json:"attribute74,omitempty"`
		Attribute75 string `json:"attribute75,omitempty"`
		Attribute76 string `json:"attribute76,omitempty"`
		Attribute77 string `json:"attribute77,omitempty"`
		Attribute78 string `json:"attribute78,omitempty"`
		Attribute79 string `json:"attribute79,omitempty"`
		Attribute80 string `json:"attribute80,omitempty"`
	}{
		Value: l.Value,
		Id:    l.Id,
		Attribute01: l.Attribute01,
		Attribute02: l.Attribute02,
		Attribute03: l.Attribute03,
		Attribute04: l.Attribute04,
		Attribute05: l.Attribute05,
		Attribute06: l.Attribute06,
		Attribute07: l.Attribute07,
		Attribute08: l.Attribute08,
		Attribute09: l.Attribute09,
		Attribute10: l.Attribute10,
		Attribute11: l.Attribute11,
		Attribute12: l.Attribute12,
		Attribute13: l.Attribute13,
		Attribute14: l.Attribute14,
		Attribute15: l.Attribute15,
		Attribute16: l.Attribute16,
		Attribute17: l.Attribute17,
		Attribute18: l.Attribute18,
		Attribute19: l.Attribute19,
		Attribute20: l.Attribute20,
		Attribute21: l.Attribute21,
		Attribute22: l.Attribute22,
		Attribute23: l.Attribute23,
		Attribute24: l.Attribute24,
		Attribute25: l.Attribute25,
		Attribute26: l.Attribute26,
		Attribute27: l.Attribute27,
		Attribute28: l.Attribute28,
		Attribute29: l.Attribute29,
		Attribute30: l.Attribute30,
		Attribute31: l.Attribute31,
		Attribute32: l.Attribute32,
		Attribute33: l.Attribute33,
		Attribute34: l.Attribute34,
		Attribute35: l.Attribute35,
		Attribute36: l.Attribute36,
		Attribute37: l.Attribute37,
		Attribute38: l.Attribute38,
		Attribute39: l.Attribute39,
		Attribute40: l.Attribute40,
		Attribute41: l.Attribute41,
		Attribute42: l.Attribute42,
		Attribute43: l.Attribute43,
		Attribute44: l.Attribute44,
		Attribute45: l.Attribute45,
		Attribute46: l.Attribute46,
		Attribute47: l.Attribute47,
		Attribute48: l.Attribute48,
		Attribute49: l.Attribute49,
		Attribute50: l.Attribute50,
		Attribute51: l.Attribute51,
		Attribute52: l.Attribute52,
		Attribute53: l.Attribute53,
		Attribute54: l.Attribute54,
		Attribute55: l.Attribute55,
		Attribute56: l.Attribute56,
		Attribute57: l.Attribute57,
		Attribute58: l.Attribute58,
		Attribute59: l.Attribute59,
		Attribute60: l.Attribute60,
		Attribute61: l.Attribute61,
		Attribute62: l.Attribute62,
		Attribute63: l.Attribute63,
		Attribute64: l.Attribute64,
		Attribute65: l.Attribute65,
		Attribute66: l.Attribute66,
		Attribute67: l.Attribute67,
		Attribute68: l.Attribute68,
		Attribute69: l.Attribute69,
		Attribute70: l.Attribute70,
		Attribute71: l.Attribute71,
		Attribute72: l.Attribute72,
		Attribute73: l.Attribute73,
		Attribute74: l.Attribute74,
		Attribute75: l.Attribute75,
		Attribute76: l.Attribute76,
		Attribute77: l.Attribute77,
		Attribute78: l.Attribute78,
		Attribute79: l.Attribute79,
		Attribute80: l.Attribute80,
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
	l.Attribute01 = uItem.Attribute01
	l.Attribute02 = uItem.Attribute02
	l.Attribute03 = uItem.Attribute03
	l.Attribute04 = uItem.Attribute04
	l.Attribute05 = uItem.Attribute05
	l.Attribute06 = uItem.Attribute06
	l.Attribute07 = uItem.Attribute07
	l.Attribute08 = uItem.Attribute08
	l.Attribute09 = uItem.Attribute09
	l.Attribute10 = uItem.Attribute10
	l.Attribute11 = uItem.Attribute11
	l.Attribute12 = uItem.Attribute12
	l.Attribute13 = uItem.Attribute13
	l.Attribute14 = uItem.Attribute14
	l.Attribute15 = uItem.Attribute15
	l.Attribute16 = uItem.Attribute16
	l.Attribute17 = uItem.Attribute17
	l.Attribute18 = uItem.Attribute18
	l.Attribute19 = uItem.Attribute19
	l.Attribute20 = uItem.Attribute20
	l.Attribute21 = uItem.Attribute21
	l.Attribute22 = uItem.Attribute22
	l.Attribute23 = uItem.Attribute23
	l.Attribute24 = uItem.Attribute24
	l.Attribute25 = uItem.Attribute25
	l.Attribute26 = uItem.Attribute26
	l.Attribute27 = uItem.Attribute27
	l.Attribute28 = uItem.Attribute28
	l.Attribute29 = uItem.Attribute29
	l.Attribute30 = uItem.Attribute30
	l.Attribute31 = uItem.Attribute31
	l.Attribute32 = uItem.Attribute32
	l.Attribute33 = uItem.Attribute33
	l.Attribute34 = uItem.Attribute34
	l.Attribute35 = uItem.Attribute35
	l.Attribute36 = uItem.Attribute36
	l.Attribute37 = uItem.Attribute37
	l.Attribute38 = uItem.Attribute38
	l.Attribute39 = uItem.Attribute39
	l.Attribute40 = uItem.Attribute40
	l.Attribute41 = uItem.Attribute41
	l.Attribute42 = uItem.Attribute42
	l.Attribute43 = uItem.Attribute43
	l.Attribute44 = uItem.Attribute44
	l.Attribute45 = uItem.Attribute45
	l.Attribute46 = uItem.Attribute46
	l.Attribute47 = uItem.Attribute47
	l.Attribute48 = uItem.Attribute48
	l.Attribute49 = uItem.Attribute49
	l.Attribute50 = uItem.Attribute50
	l.Attribute51 = uItem.Attribute51
	l.Attribute52 = uItem.Attribute52
	l.Attribute53 = uItem.Attribute53
	l.Attribute54 = uItem.Attribute54
	l.Attribute55 = uItem.Attribute55
	l.Attribute56 = uItem.Attribute56
	l.Attribute57 = uItem.Attribute57
	l.Attribute58 = uItem.Attribute58
	l.Attribute59 = uItem.Attribute59
	l.Attribute60 = uItem.Attribute60
	l.Attribute61 = uItem.Attribute61
	l.Attribute62 = uItem.Attribute62
	l.Attribute63 = uItem.Attribute63
	l.Attribute64 = uItem.Attribute64
	l.Attribute65 = uItem.Attribute65
	l.Attribute66 = uItem.Attribute66
	l.Attribute67 = uItem.Attribute67
	l.Attribute68 = uItem.Attribute68
	l.Attribute69 = uItem.Attribute69
	l.Attribute70 = uItem.Attribute70
	l.Attribute71 = uItem.Attribute71
	l.Attribute72 = uItem.Attribute72
	l.Attribute73 = uItem.Attribute73
	l.Attribute74 = uItem.Attribute74
	l.Attribute75 = uItem.Attribute75
	l.Attribute76 = uItem.Attribute76
	l.Attribute77 = uItem.Attribute77
	l.Attribute78 = uItem.Attribute78
	l.Attribute79 = uItem.Attribute79
	l.Attribute80 = uItem.Attribute80
	return nil
}
