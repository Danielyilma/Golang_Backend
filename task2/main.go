/*
 * File: main.go
 * Description: Contains functions to remove punctuation, count word frequency,
 *              and check if a string is a palindrome (ignoring punctuation and case).
 */

package main

import (
	"strings"
	"fmt"
	"unicode"
)

/**
 * removePunctuation - Removes all punctuation characters from a string.
 * @input: The input string to process.
 * Return: A new string without punctuation.
 */
func removePunctuation(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if !unicode.IsPunct(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

/**
 * isAlphanumeric - Checks if a rune is a letter or digit.
 * @r: The rune to check.
 * Return: true if alphanumeric, false otherwise.
 */
func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

/**
 * frequencyCounter - Counts the frequency of words in a string.
 * Punctuation is removed and comparison is case-insensitive.
 * @s: Input string to analyze.
 * Return: A map where keys are words and values are their frequencies.
 */
func frequencyCounter(s string) map[string]int {
	mp := make(map[string]int)
	new_s := removePunctuation(s)

	words := strings.Fields(strings.ToLower(new_s))

	for _, word := range words {
		mp[word] += 1
	}

	return mp
}

/**
 * is_palidrom - Checks if a string is a palindrome.
 * Non-alphanumeric characters are ignored, and comparison is case-insensitive.
 * @s: The string to check.
 * Return: true if the string is a palindrome, false otherwise.
 */
func is_palidrom(s string) bool {
	new_s := strings.ToLower(s)
	for i, j := 0, len(new_s)-1; i < j; {
		for i < j && !isAlphanumeric(rune(new_s[i])) {
			i++
		}

		for i < j && !isAlphanumeric(rune(new_s[j])) {
			j--
		}

		if new_s[i] != new_s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

/**
 * main - Entry point of the program.
 * Tests the frequencyCounter and is_palidrom functions with example inputs.
 */
func main() {
	// Q1: Frequency counter example
	s := "Write a Go function! that takes a string as ?input!"
	fmt.Println(frequencyCounter(s))

	// Q2: Palindrome check example
	s = "abbb!a"
	fmt.Println(s, "Is Palindrome:", is_palidrom(s))
}
