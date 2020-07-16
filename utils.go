package simpleparsers

import (
	"unicode"
)

func startsWith(match string, check func(rune) bool) bool {
	if len(match) == 0 {
		return false
	}

	return check(rune(match[0]))
}

func startsWithDigit(match string) bool {
	return startsWith(match, unicode.IsDigit)
}

func startsWithLetter(match string) bool {
	return startsWith(match, unicode.IsLetter)
}

func startsWithChar(match string, char rune) bool {
	return startsWith(match, func(head rune) bool { return head == char })
}
