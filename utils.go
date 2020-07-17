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

func reduceParser(operation func(Parser, Parser) Parser, parsers ...Parser) Parser {
	result, otherParsers := parsers[0], parsers[1:]

	for _, p := range otherParsers {
		result = operation(result, p)
	}

	return result
}
