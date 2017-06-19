package apptree

import "fmt"

type Type string

func (t Type) ToAppTreeTypePackageName() string {
	switch t {
	case Type_Text:
		return "apptree.String"
	case Type_ListItem:
		return "apptree.ListItem"
	case Type_Date:
		return "apptree.Date"
	case Type_DateTime:
		return "apptree.DateTime"
	case Type_Int:
		return "apptree.Int"
	case Type_Float:
		return "apptree.Float"
	case Type_Relationship:
		return "apptree.ToManyRelationship"
	case Type_SingleRelationship:
		return "apptree.SingleRelationship"
	case Type_TimeInterval:
		return "apptree.TimeInterval"
	case Type_Boolean:
		return "apptree.Bool"
	case Type_DateRange:
		return "apptree.DateRange"
	case Type_DateTimeRange:
		return "apptree.DateTimeRange"
	case Type_Image:
		return "apptree.Image"
	case Type_Location:
		return "apptree.Location"
	case Type_Color:
		return "apptree.Color"
	}
	return fmt.Sprintf("INVALID TYPE %s", t)
}

const (
	Type_None               Type = ""
	Type_Text               Type = "Text"
	Type_ListItem           Type = "ListItem"
	Type_DateTime           Type = "DateTime"
	Type_Date               Type = "Date"
	Type_Int                Type = "Integer"
	Type_Float              Type = "Float"
	Type_Relationship       Type = "Relationship"
	Type_TimeInterval       Type = "TimeInterval"
	Type_Boolean            Type = "Boolean"
	Type_DateRange          Type = "DateRange"
	Type_DateTimeRange      Type = "DateTimeRange"
	Type_Image              Type = "Image"
	Type_Location           Type = "Location"
	Type_SingleRelationship Type = "SingleRelationship"
	Type_Color              Type = "Color"

	//TODO: Add attachments
)
