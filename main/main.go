package main

import (
	"fmt"
	"someInterview/mapper"
)

/*
	1 . Write a function with this signature:

	func CapitalizeEveryThirdAlphanumericChar(s string) string {
		// your code goes here
	}

	that capitalizes *only* every third alphanumeric character of a string, so that if I hand you
	Aspiration.com
	You hand me back
	asPirAtiOn.cOm
	Please note:
	- Characters other than each third should be downcased
	- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
	  characters, ie [a-z][A-Z][0-9].
	2. Do the same problem, but this time create a "mapper" package that looks like this:
*/

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	runeArray := []rune(s)
	count := 1
	for i := range runeArray {
		if (count)%3 == 0 {
			if runeArray[i] >= 'a' && runeArray[i] <= 'z' {
				runeArray[i] = runeArray[i] - 32
			}
		} else {
			if runeArray[i] >= 'A' && runeArray[i] <= 'Z' {
				runeArray[i] = runeArray[i] + 32
			}
		}

		if (runeArray[i] >= 'A' && runeArray[i] <= 'Z') || (runeArray[i] >= 'a' && runeArray[i] <= 'z') {
			count++
		}
	}

	return string(runeArray)
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)

	// This is a implementation of the solution withou having to modify the original code
	// of the assignment
	s2 := NewSkipString(3, "Aspiration.com")
	mapper.MapString2(&s2)
	fmt.Println(s2)
}

func NewSkipString(skip int, value string) mapper.SkipString {
	ss := mapper.SkipString{
		Value: value,
		Skip:  skip,
	}

	return ss
}
