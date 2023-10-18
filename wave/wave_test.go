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
	maxWordLenTestCase{"a b c d", 1},
	maxWordLenTestCase{"a bc def gh ijk word", 4},
	maxWordLenTestCase{"punctuation's very tricky", 13},
}

func TestMaxWordLen(t *testing.T) {
	for nr, testCase := range maxWordLenTestCases {
		l, err := maxWordLen(testCase.input)
		assert.Nil(t, err, "Error should not occur")
		assert.Equal(t, testCase.expected, l, "Case %d failed", nr)
	}
}

type makeIndentTestCase struct {
	index, NumberOfColumns int
	expected               string
}

var makeIndentTestCases = []makeIndentTestCase{
	makeIndentTestCase{0, 1, ""},
	makeIndentTestCase{1, 1, " "},
	makeIndentTestCase{2, 1, ""},
	makeIndentTestCase{1, 2, " "},
	makeIndentTestCase{2, 2, "  "},
	makeIndentTestCase{4, 2, ""},
	makeIndentTestCase{6, 2, "  "},
	makeIndentTestCase{6, 5, "    "},
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
