package main

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

func isWordSeparator(r rune) bool {

	return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsMark(r) ||
		unicode.IsControl(r) || unicode.IsSymbol(r)
}

func patchedScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])

		// catch the invalid / binary input
		if r == utf8.RuneError {
			return start, data[start:width], errors.New("invalid input: binary format")
		}

		if !isWordSeparator(r) {
			break
		}
	}

	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		// catch the invalid / binary input
		if r == utf8.RuneError {
			return start, data[start:width], errors.New("invalid input: binary format")
		}

		if isWordSeparator(r) {
			return i + width, data[start:i], nil
		}

	}

	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}
