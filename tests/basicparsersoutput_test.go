package tests

import (
	"testing"

	"github.com/cimimuxmaio/simpleparsers"
)

var (
	digitParser       simpleparsers.Parser = simpleparsers.NewDigitParser()
	letterParser      simpleparsers.Parser = simpleparsers.NewLetterParser()
	alfaNumCharParser simpleparsers.Parser = simpleparsers.NewAlphanumericParser()
	wordParser        simpleparsers.Parser = simpleparsers.NewWordParser()
	integerParser     simpleparsers.Parser = simpleparsers.NewIntegerParser()
	numberParser      simpleparsers.Parser = simpleparsers.NewNumberParser()
)

func TestDigitParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"1asdasd", newParserOutput("1", "asdasd")},
		stringParserOutputTestCase{"5012", newParserOutput("5", "012")},
		stringParserOutputTestCase{"hello", nil},
		stringParserOutputTestCase{"+2-%2", nil},
		stringParserOutputTestCase{"$3-%2", nil},
		stringParserOutputTestCase{"asd32", nil},
	}

	assertAllEqualsParserOutput(t, digitParser, testCases)
}

func TestLetterParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"a1231", newParserOutput("a", "1231")},
		stringParserOutputTestCase{"hello", newParserOutput("h", "ello")},
		stringParserOutputTestCase{"OlLeH", newParserOutput("O", "lLeH")},
		stringParserOutputTestCase{"12345", nil},
		stringParserOutputTestCase{"+23Abc", nil},
		stringParserOutputTestCase{"$a245", nil},
	}

	assertAllEqualsParserOutput(t, letterParser, testCases)
}

func TestAlfanumericParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"a1231", newParserOutput("a", "1231")},
		stringParserOutputTestCase{"hello", newParserOutput("h", "ello")},
		stringParserOutputTestCase{"12345", newParserOutput("1", "2345")},
		stringParserOutputTestCase{"OlLeH", newParserOutput("O", "lLeH")},
		stringParserOutputTestCase{"+23Abc", nil},
		stringParserOutputTestCase{"$a245", nil},
	}

	assertAllEqualsParserOutput(t, alfaNumCharParser, testCases)
}

func TestWordParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"hello world", newParserOutput("hello", " world")},
		stringParserOutputTestCase{"a1231", newParserOutput("a", "1231")},
		stringParserOutputTestCase{"asd123", newParserOutput("asd", "123")},
		stringParserOutputTestCase{"IusaqE123", newParserOutput("IusaqE", "123")},
		stringParserOutputTestCase{"dSa+DsA.", newParserOutput("dSa", "+DsA.")},
		stringParserOutputTestCase{"12345", nil},
		stringParserOutputTestCase{"1aloha", nil},
		stringParserOutputTestCase{"+23Abc", nil},
	}

	assertAllEqualsParserOutput(t, wordParser, testCases)
}

func TestIntegerParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"123123", newParserOutput("123123", "")},
		stringParserOutputTestCase{"9876asd", newParserOutput("9876", "asd")},
		stringParserOutputTestCase{"1hello", newParserOutput("1", "hello")},
		stringParserOutputTestCase{"+754//", nil},
	}

	assertAllEqualsParserOutput(t, integerParser, testCases)
}

func TestRegexParserEquivalentToWordParser(t *testing.T) {
	regexParser, _ := simpleparsers.NewRegexParser("[a-zA-Z]+") // Equivalent to wordParser

	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"hello world", newParserOutput("hello", " world")},
		stringParserOutputTestCase{"a1231", newParserOutput("a", "1231")},
		stringParserOutputTestCase{"asd123", newParserOutput("asd", "123")},
		stringParserOutputTestCase{"IusaqE123", newParserOutput("IusaqE", "123")},
		stringParserOutputTestCase{"dSa+DsA.", newParserOutput("dSa", "+DsA.")},
		stringParserOutputTestCase{"12345", nil},
		stringParserOutputTestCase{"1aloha", nil},
		stringParserOutputTestCase{"+23Abc", nil},
	}

	assertAllEqualsParserOutput(t, regexParser, testCases)
}

func TestRegexParserCustom(t *testing.T) {
	regexParser, _ := simpleparsers.NewRegexParser("[3-9]+")

	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"3412", newParserOutput("34", "12")},
		stringParserOutputTestCase{"9876asd", newParserOutput("9876", "asd")},
		stringParserOutputTestCase{"123123", nil},
		stringParserOutputTestCase{"+754//", nil},
		stringParserOutputTestCase{"1hello", nil},
		stringParserOutputTestCase{"hello world", nil},
	}

	assertAllEqualsParserOutput(t, regexParser, testCases)
}

func TestCharParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"1+", nil},
		stringParserOutputTestCase{"+", newParserOutput("+", "")},
		stringParserOutputTestCase{"+asd123", newParserOutput("+", "asd123")},
		stringParserOutputTestCase{"a1231", nil},
		stringParserOutputTestCase{"", nil},
	}

	assertAllEqualsParserOutput(t, simpleparsers.NewCharParser('+'), testCases)
}

func TestNumberParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"123123", newParserOutput("123123", "")},
		stringParserOutputTestCase{"9876asd", newParserOutput("9876", "asd")},
		stringParserOutputTestCase{"1.234", newParserOutput("1.234", "")},
		stringParserOutputTestCase{"98.1world", newParserOutput("98.1", "world")},
		stringParserOutputTestCase{"65.hi", newParserOutput("65", ".hi")},
		stringParserOutputTestCase{"+754//", nil},
	}

	assertAllEqualsParserOutput(t, numberParser, testCases)
}
