package example

import "time"

type Issue struct {
	Id   string    `apptree:"name=Key;type=Int"`
	Date time.Time `apptree:"index=5;name=SomeDate;require;canUpdate"`
}
