[![Build Status](https://travis-ci.org/CimimUxMaio/simpleparsers.svg?branch=master)](https://travis-ci.org/CimimUxMaio/simpleparsers)
[![BCH compliance](https://bettercodehub.com/edge/badge/CimimUxMaio/simpleparsers?branch=master)](https://bettercodehub.com/)

# simpleparsers

A package with functions to create and combine parsers.

---

## Built-in parsers

- [RegexParser](#regexparser)
- [AnyCharParser](#anycharparser)
- [DigitParser](#digitparser)
- [LetterParser](#letterparser)
- [AlphanumericParser](#alphanumericparser)
- [CharParser](#charparser)
- [WordParser](#wordparser)
- [IntegerParser](#integerparser)
- [NumberParser](#numberparser)

## Parser operations

- [Sequence/2](#sequence)
- [Either/2](#either)
- [KleenePlus/1](#kleenePlus)
- [KleeneStar/1](#kleeneStar)
- [Optional/1](#optional)
- [Consume/1](#consume)
- [Conditional/2](#conditional)

---

## Built-in parsers

### RegexParser

Created with `NewRegexParser(<regex>)`. Parses a string matching a prefix that matches the given regex.

`NewRegexParser/2` may return an error if the given regex has not a valid format.

##### Examples:

```
regexParser, _ := NewRegexParser("hello (world )+")


regexParser.Parse("hello world")
> &ParserOutput{ Match: "hello world", Remainder: ""}, nil


regexParser.Parse("hello world world world bananas")
> &ParserOutput{ Match: "hello world world world ", Remainder: "bananas" }, nil


regexParser.Parse("A creative sentence.")
> nil, err

err.Error()
> No match found for regex: "hello (world )+" and input: "A creative sentence.".
```

### AnyCharParser

Created with `NewAnyCharParser`. Parses a string matching any character.

##### Examples:

```
anyParser := NewAnyCharParser()


anyParser.Parse("aloha")
> &ParserOutput{ Match: "a", Remainder: "loha"}, nil


anyParser.Parse("9876")
> &ParserOutput{ Match: "9", Remainder: "876"}, nil


anyParser.Parse("!=(&/($")
> &ParserOutput{ Match: "!", Remainder: "=(&/($"}, nil


anyParser.Parse("")
> nil, err

err.Error()
> No character match for empty string.

```

### DigitParser

Created with `NewDigitParser()`. Parses a string matching the first character if it is a _digit_ according to `unicode.IsDigit`.

It is equivalent to: [Conditional(NewAnyCharParser(), startsWithDigit)](#conditional)

where `startsWithDigit` returns true if the head of the string is in fact a _digit_.

##### Examples:

```
digitParser := NewDigitParser()

digitParser.Parse("123hello")
> &ParserOutput{ Match: "1", Remainder: "23hello" }, nil


digitParser.Parse("hello")
> nil, err

err.Error()
> No match found for character with condition: "github.com/cimimuxmaio/simpleparsers.startsWithDigit" and input: "hello".
```

### LetterParser

Created with `NewLetterParser()`. Parses a string matching the first character if it is a _letter_ according to `unicode.IsLetter`.

It is equivalent to: [Conditional(NewAnyCharParser(), startsWithLetter)](#conditional)

where `startsWithLetter` returns true if the head of the string is in fact a _letter_.

##### Examples:

```
letterParser := NewLetterParser()

letterParser.Parse("bananas!")
> &ParserOutput{ Match: "b", Remainder: "ananas!" }, nil


letterParser.Parse("1 this should fail")
> nil, err

err.Error()
> No match found for character with condition: "github.com/cimimuxmaio/simpleparsers.startsWithLetter" and input: "1 this should fail".
```

### AlphanumericParser

Created with `NewAlphanumericParser()`. Parses a string matching the first character if it is either a _letter_ or a _digit_ according to `unicode.IsLetter` and `unicode.IsDigit`.

It is equivalent to: [`Either(NewLetterParser(), NewDigitParser())`](#either)

##### Examples:

```
alphaParser := NewAlphanumericParser()

alphaParser.Parse("2020")
> &ParserOutput{ Match: "2", Remainder: "020" }, nil


alphaParser.Parse("abcdefg...")
> &ParserOutput{ Match: "a", Remainder: "bcdefg..." }, nil


alphaParser.Parse("!!! :(")
> nil, err

err.Error()
> No match found for character with condition: "github.com/cimimuxmaio/simpleparsers.startsWithDigit" and input: "!!! :(".
```

### CharParser

Created with `NewCharParser(<a_character>)`. Parses a string matching the first character only if it is the same as the given character.

It is equivalent to:
`Conditional(NewAnyCharParser(), func(match string) bool { return startsWithChar(match, char) })`

where `startsWithChar` returns true if `match` starts with `char`.

##### Examples:

```
charParser := NewCharParser('-')

charParser.Parse("-+/*")
> &ParserOutput{ Match: "-", Remainder: "+/*" }, nil


letterParser.Parse("sample text")
> nil, err

err.Error()
> No match found for character with condition: "github.com/cimimuxmaio/simpleparsers.NewCharParser.func1.1" and input: "sample text". // Should be improved
```

### WordParser

Created with `NewLetterParser()`. Parses a string matching the first word (sequence of _letters_).

It is equivalent to: [`KleenePlus(NewLetterParser())`](#kleenePlus)

##### Examples:

```
wordParser := NewWordParser()

wordParser.Parse("hello world!")
> &ParserOutput{ Match: "hello", Remainder: " world!" }


wordParser.Parse("Hi987 Wo789!")
> &ParserOutput{ Match: "Hi", Remainder: "987 Wo789!" }

wordParser.Parse("123.0")

err.Error()
> No match found for character with condition: "github.com/cimimuxmaio/simpleparsers.startsWithLetter" and input: "123.0".

```

### IntegerParser

Created with `NewIntegerParser()`. Parses a string matching the first integer number (sequence of _digits_).

It is equivalent to: [`KleenePlus(NewDigitParser())`](#kleenePlus)

##### Examples:

```
integerParser := NewIntegerParser()

integerParser.Parse("2020")
> &ParserOutput{ Match: "2020", Remainder: "" }


integerParser.Parse("123.0")
> &ParserOutput{ Match: "123", Remainder: ".0" }

integerParser.Parse("Hola mundo")
> nil, err

err.Error()
> No match found for character with condition: "github.com/cimimuxmaio/simpleparsers.startsWithDigit" and input: "Hola mundo".

```

### NumberParser

Created with `NewNumberParser()`. Parses a string matching the first **number** (either an _integer_ or a _floating point number_).

It is equivalent to:

`Sequence(NewIntegerParser(), Optional(Sequence(NewCharParser('.'), NewIntegerParser())))`

##### See:

- [Sequence/1](#sequence)
- [Optional/1](#optional)

##### Examples:

```
numberParser := NewNumberParser()

numberParser.Parse("2020")
> &ParserOutput{ Match: "2020", Remainder: "" }


numberParser.Parse("123.0")
> &ParserOutput{ Match: "123.0", Remainder: "" }


numberParser.Parse("999.Nice!")
> &ParserOutput{ Match: "999", Remainder: ".Nice!" }


numberParser.Parse("Parsers!")
> nil, err

err.Error()
> No match found for character with condition: "github.com/cimimuxmaio/simpleparsers.startsWithDigit" and input: "Parsers!".

```

---

## Parser operations

### Sequence

Given two parsers, returns a new one that parses as the two parsers consecutively, using the remainder of the first one as input for the second.

If for a certain input either parser fails, the resulting parser of sequencing both will also fail for this input.

##### Examples:

**TODO**

### Either

Given two parsers, returns a new one that parses a certain input as the first one. If it fails, parses as the second. If both fail then this parser will also fail.

##### Examples:

**TODO**

### KleenePlus

Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.

If there is **no** matches, the parser returns an error.

##### Examples:

**TODO**

### KleeneStar

Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.

If there is **no** matches, the parser matches the empty string (`""`).

##### Examples:

**TODO**

### Optional

Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.

If there is **no** matches, the parser returns an error.

##### Examples:

**TODO**

### Consume

Returns a parser that parses a certain input as the given parser.

If it matches, ignores the matching string, returning: `&ParserOutput{ Match: "", Remainder: remainder }`
If there is **no** match, returns an error.

##### Examples:

**TODO**

### Conditional

Returns a parser that parses a certain input as the given parser.

It will only match with strings that satisfy the given condition.

##### Examples:

**TODO**
