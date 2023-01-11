package markup

import "testing"

func TestReplacePlaceholders(t *testing.T) {
	tests := map[string]struct {
		replacement, value string
		expected           string
	}{
		"no placeholder": {
			replacement: "many pies",
			value:       "3",
			expected:    "many pies",
		},
		"one placeholder": {
			replacement: "% pies",
			value:       "3",
			expected:    "3 pies",
		},
		"several placeholders": {
			replacement: "% pies... % pies... % pies!",
			value:       "5",
			expected:    "5 pies... 5 pies... 5 pies!",
		},
		"escaped percent signs": {
			replacement: "Your % pies qualify you for a 10\\% discount!",
			value:       "10",
			expected:    "Your 10 pies qualify you for a 10% discount!",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if actual := replacePlaceholders(test.replacement, test.value); actual != test.expected {
				t.Errorf("unexpected result: got [%s], wanted [%s]", actual, test.expected)
			}
		})
	}
}
