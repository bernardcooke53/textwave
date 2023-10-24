package wave

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"strings"
	"unicode"
)

// WaveMaker is responsible for formatting a string
// according to the associated options
type WaveMaker struct {
	// SpongebobMocking gives you sPoNgEBoB MocKiNg TeXt
	SpongebobMocking bool
	// AllCaps gives you all capitalized text and ignores
	// SpongebobMocking
	AllCaps bool
	// ColumnSize is the spacing between each letter
	// horizontally
	ColumnSize int
	// NumberOfColumns allows you to specify how wide the wave
	// is in column numbers
	NumberOfColumns int
}

func (w *WaveMaker) MakeWave(str string) []string {
	var NumberOfColumns int
	if w.NumberOfColumns > 0 {
		NumberOfColumns = w.NumberOfColumns
	} else {
		maxWordLen := maxWordLen(str)
		NumberOfColumns = int(float64(maxWordLen / 2))
	}

	processedChars := w.preprocessChars(str)

	lines := make([]string, 0)

	for index, char := range processedChars {
		indent := w.makeIndent(index, NumberOfColumns)
		lines = append(lines, fmt.Sprintf("%s%c", indent, char))
	}
	return lines
}

func (w *WaveMaker) preprocessChars(str string) string {
	if w.AllCaps {
		return strings.ToUpper(str)
	}
	if w.SpongebobMocking {
		return preprocessCharsForSpongebobMocking(str)
	}
	return str
}

func preprocessCharsForSpongebobMocking(str string) string {
	builder := strings.Builder{}
Loop:
	for index, char := range str {
		if index == 0 {
			builder.WriteRune(char)
			continue Loop
		}
		if unicode.IsSpace(char) {
			builder.WriteRune(char)
			continue Loop
		}
		surprise := rand.Float64() < 0.2 // #nosec: G404

		// If the last letter was lowercase, 20% chance we write a lowercase letter
		if unicode.IsLower(rune(str[index-1])) || (unicode.IsSpace(rune(str[index-1])) && index > 1 && unicode.IsLower(rune(str[index-2]))) {
			if !surprise {
				builder.WriteRune(unicode.ToUpper(char))
				continue Loop
			}
		}

		// 20% chance we write two uppercase letters consecutively
		if unicode.IsUpper(rune(str[index-1])) || (unicode.IsSpace(rune(str[index-1])) && index > 1 && unicode.IsUpper(rune(str[index-2]))) {
			if surprise {
				builder.WriteRune(unicode.ToUpper(char))
				continue Loop
			}
		}
		builder.WriteRune(char)
	}

	return builder.String()
}

func (w *WaveMaker) makeIndent(index, NumberOfColumns int) string {
	indentLevel := NumberOfColumns - int(math.Abs(float64(
		NumberOfColumns-index%(NumberOfColumns*2),
	)))
	return strings.Repeat(" ", indentLevel*w.ColumnSize)
}

func maxWordLen(str string) int {
	fields := strings.Fields(str)
	fieldLengths := arrayMap(fields, func(s string) int { return len(s) })
	return slices.Max(fieldLengths)
}

func arrayMap[T, R any](it []T, f func(T) R) []R {
	ret := make([]R, len(it))
	for _, v := range it {
		ret = append(ret, f(v))
	}
	return ret
}
