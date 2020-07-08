package cal8

import (
	"strings"
)

// Center returns a centered version of the string given, padded on either side by padStr.
// If the totalSize is not even, then the left side will have the extra padding.
func Center(str string, padStr string, totalSize int) string {
	right := (totalSize - len(str)) / 2 // integer floor division
	left := totalSize - right - len(str)

	if left < 0 || right < 0 {
		panic("invalid arguments to center")
	}

	return strings.Repeat(padStr, left) + str + strings.Repeat(padStr, right)
}

// HorizontalAppend takes two strings and returns a new string where the arguments are next to each other.
// They are seperated by the longest line in the first argument. space is the string to seperate them with, usually a space.
func HorizontalAppend(left, right, space string) string {
	leftLines := strings.Split(left, "\n")
	rightLines := strings.Split(right, "\n")

	var maxLeft int
	var out string

	for _, l := range leftLines {
		if len(l) > maxLeft {
			maxLeft = len(l)
		}
	}

	for i := range leftLines {
		l := leftLines[i]
		r := ""

		if !(i >= len(rightLines)) {
			r = rightLines[i]
		}

		out += l + strings.Repeat(" ", maxLeft-len(l)) + space + r

		if i != len(leftLines)-1 {
			out += "\n"
		}
	}

	if len(rightLines) <= len(leftLines) {
		return out
	}

	out += "\n"

	for i := len(leftLines); i < len(rightLines); i++ {
		out += strings.Repeat(" ", maxLeft) + space + rightLines[i]

		if i != len(rightLines)-1 {
			out += "\n"
		}
	}

	return out
}
