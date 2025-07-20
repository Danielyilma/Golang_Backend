# Student Grade Calculator

This Go program allows users to input a student's name, number of subjects, and the grades for each subject. It then calculates and displays the average grade along with a summary of all input data.

## 📋 Features

- Input student name and number of subjects
- Enter multiple subjects and their corresponding grades
- Validates grade inputs (0–100)
- Calculates the average grade
- Displays a well-formatted report

## 🛠️ Technologies

- Go (Golang)
- `bufio` for input handling
- `strconv` for parsing integers
- `strings` for string manipulation

## 🚀 How to Run

1. Make sure you have Go installed.
2. Save the file as `main.go`.
3. Run the program:

```bash
go run main.go
```

## 📈 Example Output

```text
go run task1.go
Enter student name (or 'exit' to quit):
John
Enter number of subjects:
2
Enter subject 1 name: Math
Enter grade for subject test: 30
Enter subject 2 name: Physics
Enter grade for subject re: 45
--------------------------------------------------



--------------------------------------------------
Student Name: John
Number of Subjects: 2
Subject: Math, Grade: 30
Subject: Physics, Grade: 45
--------------------------------------------------
Average Grade: 37.50
--------------------------------------------------
```

