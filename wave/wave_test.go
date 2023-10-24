package wave

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type maxWordLenTestCase struct {
	input    string
	expected int
}

var maxWordLenTestCases = []maxWordLenTestCase{
	{"a b c d", 1},
	{"a bc def gh ijk word", 4},
	{"punctuation's very tricky", 13},
}

func TestMaxWordLen(t *testing.T) {
	for nr, testCase := range maxWordLenTestCases {
		l := maxWordLen(testCase.input)
		assert.Equal(t, testCase.expected, l, "Case %d failed", nr)
	}
}

type makeIndentTestCase struct {
	index, NumberOfColumns int
	expected               string
}

var makeIndentTestCases = []makeIndentTestCase{
	{0, 1, ""},
	{1, 1, " "},
	{2, 1, ""},
	{1, 2, " "},
	{2, 2, "  "},
	{4, 2, ""},
	{6, 2, "  "},
	{6, 5, "    "},
}

var makeIndentTestCaseColumnSizes = []int{1, 2, 3}

func TestMakeIndent(t *testing.T) {
	for _, ColumnSize := range makeIndentTestCaseColumnSizes {
		w := WaveMaker{
			SpongebobMocking: false,
			AllCaps:          false,
			ColumnSize:       ColumnSize,
			NumberOfColumns:  0,
		}

		for index, testCase := range makeIndentTestCases {
			indent := w.makeIndent(testCase.index, testCase.NumberOfColumns)
			assert.Equal(
				t,
				strings.Repeat(testCase.expected, ColumnSize),
				indent,
				"test case %d failed with ColumnSize %d",
				index,
				ColumnSize,
			)
		}
	}
}
