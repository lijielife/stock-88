package task

import "testing"

func TestReport(t *testing.T) {
	task := NewBasicsTask()
	task.CreateTables()
	task.Report(2014, 1)
}

func TestCreateTables(t *testing.T) {
	task := NewBasicsTask()
	err := task.CreateTables()
	if err != nil {
		t.Error(err)
	}
}
