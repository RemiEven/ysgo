package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/remieven/ysgo/internal/tree"
)

func main() {
	writeDialogue(parseDialogue())
}

func parseDialogue() *tree.Dialogue {
	dirEntries, err := os.ReadDir(".")
	if err != nil {
		slog.Error("failed to read dir entries", "error", err)
		os.Exit(1)
	}

	readers := make([]io.Reader, 0, len(dirEntries))

	for _, entry := range dirEntries {
		entryName := entry.Name()
		isYarnFile := entry.Type().IsRegular() && strings.HasSuffix(entryName, ".yarn")
		if !isYarnFile {
			continue
		}
		reader, err := os.Open(entryName)
		if err != nil {
			slog.Error("failed to open file", "file", entryName, "error", err)
			os.Exit(1)
		}
		defer reader.Close()
		readers = append(readers, reader)
	}

	d, err := tree.FromReader(io.MultiReader(readers...))
	if err != nil {
		slog.Error("failed to parse dialogue", "error", err)
		os.Exit(1)
	}
	return d
}

func writeDialogue(d *tree.Dialogue) {
	dialogueFile, err := os.Create("dialogue.json")
	if err != nil {
		slog.Error("failed open destination file", "error", err)
		os.Exit(1)
	}
	defer dialogueFile.Close()

	writeDialogueToWriter(d, dialogueFile)
}

func writeDialogueToWriter(d *tree.Dialogue, writer io.Writer) {
	if err := json.NewEncoder(writer).Encode(d); err != nil {
		slog.Error("failed to encode dialogue", "error", err)
		os.Exit(1)
	}
}
