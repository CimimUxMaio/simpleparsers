package tests

import (
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
