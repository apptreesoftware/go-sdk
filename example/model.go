package example

import "time"

type Issue struct {
	Id    string     `apptree:"index=1;"`
	Date  time.Time  `apptree:"index=2;name=SomeDate;require;canUpdate"`
	Tasks []Task     `apptree:"index=3"`
	Code  StatusCode `apptree:"index=4;type=ListItem"`
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
