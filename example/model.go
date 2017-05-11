package example

import "time"

type Issue struct {
	Id    string     `apptree:"name=Key;type=Int"`
	Date  time.Time  `apptree:"index=5;name=SomeDate;require;canUpdate"`
	Tasks []Task     `apptree:"index=6"`
	Code  StatusCode `apptree:"type=ListItem"`
}

//Relationship
type Task struct {
	Name        string
	Description string
}

//ListItem
type StatusCode struct {
	Code        int
	Description string
}
