package terminal

import (
	"testing"

	"github.com/remieven/ysgo/markup"
)

func TestToStyledText(t *testing.T) {
	tests := map[string]struct {
		input          *markup.ParseResult
		expectedOutput string
	}{
		"no markup": {
			input: &markup.ParseResult{
				Text: "Hello",
			},
			expectedOutput: "Hello",
		},
		"simple color": {
			input: &markup.ParseResult{
				Text: "Hello",
				Attributes: []markup.Attribute{
					{
						Position: 0,
						Length:   5,
						Name:     "color",
						Properties: map[string]markup.Value{
							"color": {StringValue: "red"},
						},
					},
				},
			},
			expectedOutput: "[red]Hello[-]",
		},
		"simple italic": {
			input: &markup.ParseResult{
				Text: "Hello there",
				Attributes: []markup.Attribute{
					{
						Position: 6,
						Length:   5,
						Name:     "i",
					},
				},
			},
			expectedOutput: "Hello [::i]there[::-]",
		},
		"mixed colors and styles": {
			input: &markup.ParseResult{
				Text: "Hello there. General Kenobi.",
				Attributes: []markup.Attribute{
					{
						Position: 6,
						Length:   14,
						Name:     "color",
						Properties: map[string]markup.Value{
							"color": {StringValue: "teal"},
						},
					},
					{
						Position: 6,
						Length:   14,
						Name:     "i",
					},
					{
						Position: 13,
						Length:   14,
						Name:     "u",
					},
				},
			},
			expectedOutput: "Hello [teal][::i]there. [::iu]General[-][::u] Kenobi[::-].",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualOutput := toStyledText(test.input)
			if actualOutput != test.expectedOutput {
				t.Errorf("unexpected outout: wanted [%s], got [%s]", test.expectedOutput, actualOutput)
			}
		})
	}
}
