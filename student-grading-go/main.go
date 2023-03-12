package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

// TODO: Implement stringer method for student type
type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {

	//check if the filepath given is of a csv file
	extension := filepath.Ext(filePath)
	if extension != ".csv" {
		log.Fatal("Given file is not a csv file")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	//skip header
	if scanner.Scan() {
		_ = scanner.Text()
	}

	var students []student

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		// Parse the test scores into integers
		test1Score, _ := strconv.Atoi(fields[3])
		test2Score, _ := strconv.Atoi(fields[4])
		test3Score, _ := strconv.Atoi(fields[5])
		test4Score, _ := strconv.Atoi(fields[6])

		// Create a new student struct and assign values to its fields
		student := student{
			firstName:  fields[0],
			lastName:   fields[1],
			university: fields[2],
			test1Score: test1Score,
			test2Score: test2Score,
			test3Score: test3Score,
			test4Score: test4Score,
		}

		students = append(students, student)
	}

	if err := scanner.Err(); err != nil {
		//log.Fatal(err) vs panic(err)
		log.Fatal(err)
	}

	return students
}

func calculateGrade(students []student) []studentStat {

	var studentStats []studentStat

	for _, student := range students {
		finalScore := float32(
			student.test1Score+
				student.test2Score+
				student.test3Score+
				student.test4Score) / 4.0

		var grade Grade

		switch {
		case finalScore < 35:
			grade = F
		case finalScore < 50:
			grade = C
		case finalScore < 70:
			grade = B
		default:
			grade = A
		}

		studentStat := studentStat{
			student:    student,
			finalScore: finalScore,
			grade:      grade,
		}
		studentStats = append(studentStats, studentStat)
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	top := gradedStudents[0]
	for _, s := range gradedStudents {
		if s.finalScore > top.finalScore {
			top = s
		}
	}
	return top
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	toppers := make(map[string]studentStat)

	for _, s := range gs {
		if currentTopper, ok := toppers[s.university]; !ok || s.finalScore > currentTopper.finalScore {
			toppers[s.university] = s
		}
	}

	return toppers
}

//TODO: find all the toppers having same marks
