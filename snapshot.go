package ysgo

import "github.com/remieven/ysgo/variable"

// Snapshot holds data that represents the current state of a dialogue, so it can be restored later.
type Snapshot struct {
	Variables      map[string]variable.Value
	SmartVariables map[string]variable.Expression

	CurrentNode string

	VisitedNodes map[string]int
}
