# Text Utility: Word Frequency & Palindrome Checker

This Go program provides two utility functions:

1. A word frequency counter that removes punctuation and counts the number of times each word appears.
2. A palindrome checker that determines whether a string is a palindrome, ignoring punctuation and letter casing.

## ✨ Features

* Removes punctuation from input strings
* Converts strings to lowercase for consistent processing
* Counts word occurrences accurately
* Checks whether a string is a valid palindrome
* Ignores non-alphanumeric characters during palindrome checks

## 🛠️ Technologies

* Go (Golang)
* `strings` for text manipulation
* `unicode` for character analysis

## 🚀 How to Run

1. Save the code as `main.go`.
2. Open a terminal and run the following command:

```bash
go run main.go
```

## 🔍 Example Output

```text
map[as:1 function:1 go:1 input:1 string:2 takes:1 that:1 write:1]
abbb!a Is Palindrome: true
```

## 📊 Functions

### removePunctuation(string) string

Removes punctuation characters from the input string.

### isAlphanumeric(rune) bool

Checks if a character is a letter or digit.

### frequencyCounter(string) map\[string]int

Counts word frequency after cleaning and converting to lowercase.

### is\_palidrom(string) bool

Checks if the input string is a palindrome, ignoring case and punctuation.

## 📖 License

This project is provided for educational purposes and is released under the MIT license.
