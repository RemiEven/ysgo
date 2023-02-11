package ysgo_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	stepTypeLine = iota
	stepTypeOption
	stepTypeSelect
	stepTypeCommand
	stepTypeStop
)

type TestPlanStep struct {
	stepType            int
	stringValue         string
	intValue            int
	expectOptionEnabled bool
}

func newTestPlanStep(line string) (*TestPlanStep, error) {
	step := &TestPlanStep{}
	line = strings.TrimSpace(line)
	if len(line) == 0 || strings.HasPrefix(line, "#") {
		return nil, nil
	}

	switch {
	case strings.HasPrefix(line, "line:"):
		step.stepType = stepTypeLine
		step.stringValue = line[6:]
	case strings.HasPrefix(line, "option:"):
		step.stepType = stepTypeOption
		step.stringValue = line[8:]
		step.expectOptionEnabled = true
		if strings.HasSuffix(step.stringValue, " [disabled]") {
			step.expectOptionEnabled = false
			step.stringValue = strings.TrimSuffix(step.stringValue, " [disabled]")
		}
	case strings.HasPrefix(line, "command:"):
		step.stepType = stepTypeCommand
		step.stringValue = line[9:]
	case strings.HasPrefix(line, "select:"):
		step.stepType = stepTypeSelect
		i, err := strconv.Atoi(line[8:])
		if err != nil {
			return nil, fmt.Errorf("failed to parse select index: %w", err)
		}
		step.intValue = i
	case strings.HasPrefix(line, "stop:"):
		step.stepType = stepTypeStop
	}

	return step, nil
}

type TestPlan []*TestPlanStep

func parseTestPlan(path string) (TestPlan, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read test plan file: %w", err)
	}
	lines := strings.Split(string(fileContent), "\n")
	plan := make(TestPlan, 0, len(lines))
	for i, line := range lines {
		if step, err := newTestPlanStep(line); err != nil {
			return nil, fmt.Errorf("failed to parse line [%v] of test plan file: %w", i, err)
		} else if step != nil {
			plan = append(plan, step)
		}
	}
	return plan, nil
}
