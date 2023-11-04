package internal

import (
	"errors"
	"fmt"
)

func AddStudent(student Student) error {
	err := Session.Query("INSERT INTO students(part, id, firstname, lastname, age) VALUES(1, ?, ?, ?, ?)",
		student.ID, student.Firstname, student.Lastname, student.Age).Exec()
	if err != nil {
		fmt.Println("Error while inserting")
	}
	return err
}

func MaxStudentID() (int, error) {
	m := map[string]interface{}{}
	count := 0
	iter := Session.Query("select max(id) from students where part=1").Iter()
	for iter.MapScan(m) {
		var a any
		a, ok := m["system.max(id)"]
		if !ok {
			return 0, errors.New("wrong field system.max(id)")
		}
		count, ok = a.(int)
		if !ok {
			return 0, errors.New("wrong type int system.max(id)")
		}
		return count, nil
	}
	return 0, errors.New("no rows from cassandra")
}
