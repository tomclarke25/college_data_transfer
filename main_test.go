package main

import (
	"encoding/csv"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestReadTeachers(t *testing.T) {
	filename := "data/Teachers.csv"

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open CSV file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatalf("Failed to close CSV file: %v", err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV file: %v", err)
	}

	teachers := readTeachers(filename)

	if len(teachers) != len(records)-1 {
		t.Errorf("readTeachers() returned %d teachers, want %d", len(teachers), len(records)-1)
	}

	for i, record := range records[1:] {
		staffID, err := strconv.Atoi(record[3])
		if err != nil {
			t.Fatalf("Error converting StaffID to int: %v", err)
		}
		expected := Teacher{
			FirstName: record[0],
			Surname:   record[1],
			Email:     record[2],
			StaffID:   staffID,
		}
		if !reflect.DeepEqual(teachers[i], expected) {
			t.Errorf("Teacher at index %d = %v, want %v", i, teachers[i], expected)
		}
	}
}

func TestReadClasses(t *testing.T) {
	filename := "data/Classes.csv"

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open CSV file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatalf("Failed to close file: %v", err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV file: %v", err)
	}

	classes := readClasses(filename)

	if len(classes) != len(records)-1 {
		t.Errorf("readClasses() returned %d classes, want %d", len(classes), len(records)-1)
	}

	for i, record := range records[1:] {
		staffID, err := strconv.Atoi(record[2])
		if err != nil {
			t.Fatalf("Error converting StaffID to int: %v", err)
		}
		expected := Class{
			OfferingCode: record[0],
			Title:        record[1],
			StaffID:      staffID,
			Staff:        record[3],
		}
		if !reflect.DeepEqual(classes[i], expected) {
			t.Errorf("Class at index %d = %v, want %v", i, classes[i], expected)
		}
	}
}

func TestReadStudents(t *testing.T) {
	filename := "data/Students.csv"

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open CSV file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatalf("Failed to close CSV file: %v", err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV file: %v", err)
	}

	students, studentOfferings := readStudents(filename)

	studentMap := make(map[int]Student)
	for _, record := range records[1:] {
		refNo, err := strconv.Atoi(record[5])
		if err != nil {
			t.Fatalf("Error converting RefNo to int: %v", err)
		}
		if _, exists := studentMap[refNo]; !exists {
			studentMap[refNo] = Student{
				Forename: record[0],
				Surname:  record[1],
				Email:    record[2],
				Email2:   record[3],
				RefNo:    refNo,
			}
		}
	}

	if len(students) != len(studentMap) {
		t.Errorf("readStudents() returned %d students, want %d", len(students), len(studentMap))
	}

	if len(studentOfferings) != len(records)-1 {
		t.Errorf("readStudents() returned %d student offerings, want %d", len(studentOfferings), len(records)-1)
	}

	for _, record := range records[1:] {
		refNo, err := strconv.Atoi(record[5])
		if err != nil {
			t.Fatalf("Error converting RefNo to int: %v", err)
		}
		expectedStudent := Student{
			Forename: record[0],
			Surname:  record[1],
			Email:    record[2],
			Email2:   record[3],
			RefNo:    refNo,
		}
		if !reflect.DeepEqual(studentMap[refNo], expectedStudent) {
			t.Errorf("Student with RefNo %d = %v, want %v", refNo, studentMap[refNo], expectedStudent)
		}
	}

	for i, record := range records[1:] {
		refNo, err := strconv.Atoi(record[5])
		if err != nil {
			t.Fatalf("Error converting RefNo to int: %v", err)
		}
		expectedOffering := StudentOffering{
			StudentRefNo: refNo,
			OfferingCode: record[4],
		}
		if !reflect.DeepEqual(studentOfferings[i], expectedOffering) {
			t.Errorf("StudentOffering at index %d = %v, want %v", i, studentOfferings[i], expectedOffering)
		}
	}
}
