package tree

import "testing"

func TestLineStatementGetLineId(t *testing.T) {
	tests := map[string]struct {
		tags     []string
		expected string
	}{
		"no tags": {
			tags: []string{},
		},
		"several tags but none is line": {
			tags: []string{"tone:sarcastic", "duplicate"},
		},
		"several tags and one is line": {
			tags:     []string{"line:z465f"},
			expected: "z465f",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lineStatement := LineStatement{
				Tags: test.tags,
			}

			if actual := lineStatement.LineId(); actual != test.expected {
				t.Errorf("unexpected result: got [%s], wanted [%s]", actual, test.expected)
			}
		})
	}
}
