package apptree

type TypedValue interface {
	ValueType() Type
	IsNull() bool
}
