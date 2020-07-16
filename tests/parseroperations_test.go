package tests

import (
	"strconv"
	"testing"

	"github.com/cimimuxmaio/simpleparsers"
)

func TestSequence(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"12312hola++", newParserOutput("12312hola", "++")},
		stringParserOutputTestCase{"1a123", newParserOutput("1a", "123")},
		stringParserOutputTestCase{"hola123", nil},
		stringParserOutputTestCase{"", nil},
	}

	assertAllEqualsParserOutput(t, simpleparsers.Sequence(integerParser, wordParser), testCases)
}

func TestEither(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"1a23", newParserOutput("1", "a23")},
		stringParserOutputTestCase{"world2!5", newParserOutput("world", "2!5")},
		stringParserOutputTestCase{"", nil},
		stringParserOutputTestCase{"+&%$$asd123", nil},
	}

	assertAllEqualsParserOutput(t, simpleparsers.Either(digitParser, wordParser), testCases)
}

func TestKleenePlus(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"1a23", newParserOutput("1a23", "")},
		stringParserOutputTestCase{"world2!5", newParserOutput("world2", "!5")},
		stringParserOutputTestCase{"h1e2l3l4o@", newParserOutput("h1e2l3l4o", "@")},
		stringParserOutputTestCase{"", nil},
		stringParserOutputTestCase{"+&%$$asd123", nil},
	}

	assertAllEqualsParserOutput(t, simpleparsers.KleenePlus(alfaNumCharParser), testCases)
}

func TestKleeneStar(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"1a23", newParserOutput("1a23", "")},
		stringParserOutputTestCase{"world2!5", newParserOutput("world2", "!5")},
		stringParserOutputTestCase{"h1e2l3l4o@", newParserOutput("h1e2l3l4o", "@")},
		stringParserOutputTestCase{"", newParserOutput("", "")},
		stringParserOutputTestCase{"+&%$$asd123", newParserOutput("", "+&%$$asd123")},
	}

	assertAllEqualsParserOutput(t, simpleparsers.KleeneStar(alfaNumCharParser), testCases)
}

func TestOptional(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"1a23", newParserOutput("", "1a23")},
		stringParserOutputTestCase{"world2!5", newParserOutput("w", "orld2!5")},
		stringParserOutputTestCase{"h1e2l3l4o@", newParserOutput("h", "1e2l3l4o@")},
		stringParserOutputTestCase{"", newParserOutput("", "")},
		stringParserOutputTestCase{"+&%$$asd123", newParserOutput("", "+&%$$asd123")},
	}

	assertAllEqualsParserOutput(t, simpleparsers.Optional(letterParser), testCases)
}

func TestConsume(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"1a2", newParserOutput("1", "2")},
		stringParserOutputTestCase{"9xBye", newParserOutput("9", "Bye")},
		stringParserOutputTestCase{"456", nil},
		stringParserOutputTestCase{"he1l2l3o", nil},
		stringParserOutputTestCase{"", nil},
	}

	parser := simpleparsers.Sequence(digitParser, simpleparsers.Consume(letterParser))
	assertAllEqualsParserOutput(t, parser, testCases)
}

func TestConditional(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"9201", newParserOutput("9", "201")},
		stringParserOutputTestCase{"6xBye", newParserOutput("6", "xBye")},
		stringParserOutputTestCase{"123321", nil},
		stringParserOutputTestCase{"banana", nil},
		stringParserOutputTestCase{"", nil},
		stringParserOutputTestCase{"5asd", nil},
	}

	greaterThanFive := func(match string) bool {
		intValue, err := strconv.Atoi(match)
		if err != nil {
			return false
		}

		return intValue > 5
	}

	parser := simpleparsers.Conditional(digitParser, greaterThanFive)
	assertAllEqualsParserOutput(t, parser, testCases)
}
