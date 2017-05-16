package apptree

type Type string

const (
	Text         Type = "Text"
	ListItem     Type = "ListItem"
	DateTime     Type = "DateTime"
	Date         Type = "Date"
	Int          Type = "Int"
	Float        Type = "Float"
	Relationship Type = "Relationship"

	//TODO: Check to make sure we have all types
)
