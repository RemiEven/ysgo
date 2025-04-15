package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/remieven/ysgo/internal/tree"
)

func readSampleDialogue() ([]byte, error) {
	fileReader1, err := os.Open("testdata/ghosty_lads.yarn")
	if err != nil {
		return nil, fmt.Errorf("failed to open dialogue file: %w", err)
	}
	defer fileReader1.Close()
	fileReader2, err := os.Open("testdata/space_journey.yarn")
	if err != nil {
		return nil, fmt.Errorf("failed to open dialogue file: %w", err)
	}
	defer fileReader2.Close()
	return io.ReadAll(io.MultiReader(fileReader1, fileReader2))
}

func BenchmarkParse(b *testing.B) {
	sampleDialogue, err := readSampleDialogue()
	if err != nil {
		b.Errorf("failed to read sample dialogue: %v", err)
		return
	}

	for b.Loop() {
		if _, err := tree.FromReader(bytes.NewBuffer(sampleDialogue)); err != nil {
			b.Errorf("failed: %v", err)
			return
		}
	}
}

func BenchmarkParseJson(b *testing.B) {
	sampleDialogue, err := readSampleDialogue()
	if err != nil {
		b.Errorf("failed to read sample dialogue: %v", err)
		return
	}
	d, err := tree.FromReader(bytes.NewBuffer(sampleDialogue))
	if err != nil {
		b.Errorf("failed: %v", err)
		return
	}
	buffer := bytes.NewBuffer([]byte{})
	writeDialogueToWriter(d, buffer)
	dialogueBytes := buffer.Bytes()

	for b.Loop() {
		if _, err := tree.FromMarshaledReader(bytes.NewBuffer(dialogueBytes)); err != nil {
			b.Errorf("failed: %v", err)
			return
		}
		buffer.Reset()
	}
}
