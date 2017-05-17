package apptree

type Type string

const (
	Type_Text               Type = "Text"
	Type_ListItem           Type = "ListItem"
	Type_DateTime           Type = "DateTime"
	Type_Date               Type = "Date"
	Type_Int                Type = "Int"
	Type_Float              Type = "Float"
	Type_Relationship       Type = "Relationship"
	Type_TimeInterval       Type = "TimeInterval"
	Type_Boolean            Type = "Boolean"
	Type_DateRange          Type = "DateRange"
	Type_DateTimeRange      Type = "DateTimeRange"
	Type_Image              Type = "Image"
	Type_Location           Type = "Location"
	Type_SingleRelationship Type = "SingleRelationship"
	Type_Color                   Type = "Color"

	//TODO: Add attachments
)
