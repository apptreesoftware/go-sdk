package apptree

import "testing"

func TestParseModel(t *testing.T) {
	var model TestModel
	sampleRec := sampleRecord(t)
	sampleRec.WriteInto(&model)

	expectedStr, _ := sampleRec.GetString(0)
	if model.SomeString != expectedStr {
		t.Fail()
	}

	expectedFlt, _ := sampleRec.GetFloat(1)
	if model.SomeFloat != expectedFlt {
		t.Fail()
	}

	expectedInt, _ := sampleRec.GetInt(2)
	if model.SomeInt != expectedInt {
		t.Fail()
	}

	expectedBool, _ := sampleRec.GetBool(3)
	if model.SomeBool != expectedBool {
		t.Fail()
	}

	expectedColor, _ := sampleRec.GetColor(4)
	if model.SomeColor != expectedColor {
		t.Fail()
	}

	expectedListItem, _ := sampleRec.GetListItem(5)
	if model.SomeListItem != expectedListItem {

	}

	if len(model.Children) != 2 {
		t.Fail()
	}

	config := sampleConfig()
	rec2 := NewItem(&config)
	rec2.ReadFrom(&model)

	rec2.PrimaryKey = sampleRec.PrimaryKey
	if !sampleRec.IsEqual(&rec2) {
		t.Fail()
	}

	/*
		record.SetBool(3, NewBool(true))
		record.SetColor(4, NewColor(76, 175, 80, 10))
		record.SetListItem(5, NewListItem("Test Item"))
		record.SetDate(6, NewDate(testDate))
		record.SetDateTime(7, NewDateTime(testDateTime))
		record.SetDateRange(8, NewDateRange(testDate, testDate2))
		record.SetDateTimeRange(9, NewDateTimeRange(testDateTime, testDateTime2))
		record.SetTimeInterval(10, NewTimeInterval(100))
		record.SetImage(11, NewImage("http://fakeImage.com"))
		record.SetLocation(12, testLocation)
	*/

}

type TestModel struct {
	SomeString        String           `index:"0"`
	SomeFloat         Float            `index:"1"`
	SomeInt           Int              `index:"2"`
	SomeBool          Bool             `index:"3"`
	SomeColor         Color            `index:"4"`
	Children          []TestChildModel `index:"26"`
	SomeListItem      ListItem         `index:"5"`
	SomeDate          Date             `index:"6"`
	SomeDateTime      DateTime         `index:"7"`
	SomeDateRange     DateRange        `index:"8"`
	SomeDateTimeRange DateTimeRange    `index:"9"`
	TimeInterval      TimeInterval     `index:"10"`
	SomeImage         Image            `index:"11"`
	SomeLocation      Location         `index:"12"`
	Child             TestChildModel   `index:"27"`
}

type TestChildModel struct {
	ChildString       String           `index:"0"`
	ChildFlaot        Float            `index:"1"`
	SomeInt           Int              `index:"2"`
	SomeBool          Bool             `index:"3"`
	SomeColor         Color            `index:"4"`
	Children          []TestChildModel `index:"-"`
	SomeListItem      ListItem         `index:"5"`
	SomeDate          Date             `index:"6"`
	SomeDateTime      DateTime         `index:"7"`
	SomeDateRange     DateRange        `index:"8"`
	SomeDateTimeRange DateTimeRange    `index:"9"`
	TimeInterval      TimeInterval     `index:"10"`
	SomeImage         Image            `index:"11"`
	SomeLocation      Location         `index:"12"`
}
