package tree_test

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"

	"github.com/remieven/ysgo/internal/tree"
)

func TestParserListener(t *testing.T) {
	tests := []string{
		"boolean_expressions",
		"example_project_1",
		"function_calls",
		"jump_statements",
		"number_expressions",
		"space_journey",
		"condition_statements",
		"ghosty_lads",
		"nested_shortcut_options",
		"shortcut_options",
		"string_expressions",
		"commands",
		"call_statement",
		"tags",
		"detours",
		"line_groups",
	}

	for _, name := range tests {
		t.Run(name, func(t *testing.T) {
			dialogue, err := tree.FromFile("testdata/" + name + ".yarn")
			if err != nil {
				t.Errorf("failed to create dialogue from testdata file: %v", err)
				return
			}

			cupaloy.SnapshotT(t, dialogue)
		})
	}
}
