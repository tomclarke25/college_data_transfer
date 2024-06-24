package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Teacher struct {
	FirstName string
	Surname   string
	Email     string
	StaffID   int
}

type Class struct {
	OfferingCode string
	Title        string
	StaffID      int
	Staff        string
}

type Student struct {
	Forename string
	Surname  string
	Email    string
	Email2   string
	//OfferingCode string
	RefNo int
}

type StudentOffering struct {
	StudentRefNo int
	OfferingCode string
}

func main() {

	//teachers := readTeachers("data/Teachers.csv")
	////teacherIDMap := make(map[string]int)
	//for _, teacher := range teachers {
	//	fmt.Println(teacher)
	//}

	//classes := readClasses("data/Classes.csv")
	////classIDMap := make(map[string]int)
	//for _, class := range classes {
	//	fmt.Println(class)
	//}

	//students, _ := readStudents("data/Students.csv")
	//fmt.Println("Students")
	//for _, student := range students {
	//	fmt.Println(student)
	//}

	//fmt.Println("Offerings")
	//for _, studentOffering := range studentOfferings {
	//	fmt.Println(studentOffering)
	//}

	fmt.Println("Data imported successfully")
}

func readTeachers(filename string) []Teacher {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var teachers []Teacher
	for _, record := range records[1:] { // Skip header
		staffID, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatalf("Error converting StaffID to int: %v", err)
		}
		teachers = append(teachers, Teacher{
			FirstName: record[0],
			Surname:   record[1],
			Email:     record[2],
			StaffID:   staffID,
		})
	}

	return teachers
}

func readClasses(filename string) []Class {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var classes []Class
	for _, record := range records[1:] { // Skip header
		staffID, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatalf("Error converting StaffID to int: %v", err)
		}
		classes = append(classes, Class{
			OfferingCode: record[0],
			Title:        record[1],
			StaffID:      staffID,
			Staff:        record[3],
		})
	}

	return classes
}

func readStudents(filename string) ([]Student, []StudentOffering) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var students []Student
	var studentOfferings []StudentOffering
	studentMap := make(map[int]Student)
	for _, record := range records[1:] {
		refNo, err := strconv.Atoi(record[5])
		if err != nil {
			log.Fatalf("Error converting StaffID to int: %v", err)
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
		studentOfferings = append(studentOfferings, StudentOffering{
			StudentRefNo: refNo,
			OfferingCode: record[4],
		})
	}

	for _, student := range studentMap {
		students = append(students, student)
	}

	return students, studentOfferings
}
