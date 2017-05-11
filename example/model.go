package example

import "time"

type Issue struct {
	Id   string    `apptree:"index=1;name=Key;type=Int"`
	Date time.Time `apptree:"index=2;name=SomeDate;type=DateTime;require;canUpdate"`
}
