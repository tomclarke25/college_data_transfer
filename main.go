package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type (
	Student struct {
		RegisterNumberShort string `json:"register_number_short"`
		RefNum              int    `json:"ref_num"`
		LastName            string `json:"last_name"`
		FirstName           string `json:"first_name"`
		Email               string `json:"email"`
		StaffID             int    `json:"staff_id"`
	}

	Class struct {
		RegisterNumberShort string `json:"register_number_short"`
		RegisterTitle       string `json:"register_title"`
		StaffID             int    `json:"staff_id"`
	}

	Teacher struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		StaffID   int    `json:"staff_id"`
	}
)

const (
	teacherCSVFields = 4
	studentCSVFields = 7
	classCSVFields   = 5
)

func main() {
	err := convertTeachersData()
	if err != nil {
		log.Fatalf("failed to process teachers: %v", err)
	}

	err = convertStudentsData()
	if err != nil {
		log.Fatalf("failed to process students: %v", err)
	}

	err = convertClassesData()
	if err != nil {
		log.Fatalf("failed to process classes: %v", err)
	}
}

func convertTeachersData() error {
	return convertCSVToJSON("data/Teachers.csv", "teachers.json", teacherCSVFields, createTeacher)
}

func createTeacher(record []string) (Teacher, error) {
	staffID, err := parseID(record[3])
	if err != nil {
		return Teacher{}, fmt.Errorf("failed to parse staff id: %w", err)
	}
	return Teacher{
		FirstName: record[0],
		LastName:  record[1],
		Email:     record[2],
		StaffID:   staffID,
	}, nil
}

func convertStudentsData() error {
	return convertCSVToJSON("data/Students.csv", "students.json", studentCSVFields, createStudent)
}

func createStudent(record []string) (Student, error) {
	refNum, err := parseID(record[1])
	if err != nil {
		return Student{}, fmt.Errorf("failed to parse ref number: %w", err)
	}
	staffID, err := parseID(record[6])
	if err != nil {
		return Student{}, fmt.Errorf("failed to parse staff id: %w", err)
	}
	return Student{
		RegisterNumberShort: record[0],
		RefNum:              refNum,
		LastName:            record[2],
		FirstName:           record[3],
		Email:               record[5],
		StaffID:             staffID,
	}, nil
}

func convertClassesData() error {
	return convertCSVToJSON("data/Classes.csv", "classes.json", classCSVFields, createClass)
}

func createClass(record []string) (Class, error) {
	staffID, err := parseID(record[2])
	if err != nil {
		return Class{}, fmt.Errorf("failed to parse staff id: %w", err)
	}
	return Class{
		RegisterNumberShort: record[0],
		RegisterTitle:       record[1],
		StaffID:             staffID,
	}, nil
}

func parseID(id string) (int, error) {
	convertedID, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("failed converting to Int: %w", err)
	}

	return convertedID, nil
}

func writeJSON[T any](data []T, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return fmt.Errorf("failed creating json: %w", err)
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("failed writing to file: %w", err)
	}

	return nil
}

func readCSVFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}(f)
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}

	return data, nil
}

func parseCSVRecords[T any](data []byte, createFunc func([]string) (T, error)) ([]T, error) {
	reader := csv.NewReader(bytes.NewReader(data))

	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	estimatedRecords := bytes.Count(data, []byte{'\n'})
	records := make([]T, 0, estimatedRecords)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		item, err := createFunc(record)
		if err != nil {
			return nil, fmt.Errorf("failed to create record: %w", err)
		}
		records = append(records, item)
	}
	return records, nil
}

func convertCSVToJSON[T any](inputFile, outputFile string, expectedFields int, createFunc func([]string) (T, error)) error {
	log.Printf("Processing file: %s", inputFile)
	data, err := readCSVFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	records, err := parseCSVRecords(data, func(record []string) (T, error) {
		if len(record) != expectedFields {
			return *new(T), fmt.Errorf("expected %d fields, got %d", expectedFields, len(record))
		}

		return createFunc(record)
	})
	if err != nil {
		return fmt.Errorf("failed to process CSV: %w", err)
	}

	err = writeJSON(records, outputFile)
	if err != nil {
		return fmt.Errorf("failed to create and write JSON: %w", err)
	}
	log.Printf("Successfully processed file: %s", inputFile)
	return nil
}
