/*
 * File: main.go
 * Description: This program collects student details, subjects, and grades,
 * calculates the average grade, and displays the result.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * struct Subjects - Represents a subject and its grade.
 * @Name: Name of the subject.
 * @Grade: Grade received in the subject (0–100).
 */
type Subjects struct {
	Name  string
	Grade int
}

/*
 * struct Student - Represents a student and their academic information.
 * @Name: Student's name.
 * @No_of_subjects: Number of subjects the student is enrolled in.
 * @Subjects: Slice of Subjects the student is taking.
 * @Average_grade: Computed average grade across all subjects.
 */
type Student struct {
	Name           string
	No_of_subjects int
	Subjects       []Subjects
	Average_grade  float64
}

/**
 * CalculateAverage - Calculates the average grade for the student.
 * This method updates the Average_grade field.
 */
func (s *Student) CalculateAverage() {
	if s.No_of_subjects == 0 || len(s.Subjects) == 0 {
		s.Average_grade = 0
		return
	}

	total := 0
	for _, subject := range s.Subjects {
		total += subject.Grade
	}

	s.Average_grade = float64(total) / float64(s.No_of_subjects)
}

/**
 * displayStudentInfo - Prints student information and grades.
 * @student: The student whose data is to be displayed.
 */
func displayStudentInfo(student Student) {
	fmt.Println("--------------------------------------------------\n\n\n")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Student Name: %s\n", student.Name)
	fmt.Printf("Number of Subjects: %d\n", student.No_of_subjects)
	for _, subject := range student.Subjects {
		fmt.Printf("Subject: %s, Grade: %d\n", subject.Name, subject.Grade)
	}
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Average Grade: %.2f\n", student.Average_grade)
	fmt.Println("--------------------------------------------------\n\n\n")
}

/**
 * isValidGrade - Validates if a grade is between 0 and 100.
 * @grade: Grade to validate.
 * Return: true if valid, false otherwise.
 */
func isValidGrade(grade int) bool {
	return grade >= 0 && grade <= 100
}

/**
 * main - Entry point of the program.
 * Repeatedly collects student details and grades, and shows calculated results.
 */
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var student Student

		fmt.Println("Enter student name (or 'exit' to quit):")
		nameLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading name:", err)
			return
		}
		name := strings.TrimSpace(nameLine)
		if strings.EqualFold(name, "exit") {
			break
		}
		if name == "" {
			fmt.Println("Name cannot be empty.")
			continue
		}
		student.Name = name

		fmt.Println("Enter number of subjects: ")
		numLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading number of subjects:", err)
			return
		}
		numLine = strings.TrimSpace(numLine)
		num, err := strconv.Atoi(numLine)
		if err != nil || num < 1 || num > 10 {
			fmt.Println("Invalid number; please enter a number between 1 and 10.")
			continue
		}
		student.No_of_subjects = num

		for i := 0; i < student.No_of_subjects; i++ {
			var subject Subjects

			fmt.Printf("Enter subject %d name: ", i+1)
			subjectName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading subject name:", err)
				return
			}
			subject.Name = strings.TrimSpace(subjectName)
			if subject.Name == "" {
				fmt.Println("Subject name cannot be empty.")
				i--
				continue
			}

			fmt.Printf("Enter grade for subject %s: ", subject.Name)
			gradeInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading grade:", err)
				return
			}
			gradeInput = strings.TrimSpace(gradeInput)
			grade, err := strconv.Atoi(gradeInput)
			if err != nil || !isValidGrade(grade) {
				fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
				i--
				continue
			}
			subject.Grade = grade

			student.Subjects = append(student.Subjects, subject)
		}

		student.CalculateAverage()
		displayStudentInfo(student)
	}
}
