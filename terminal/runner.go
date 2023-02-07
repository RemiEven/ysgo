// Package terminal uses the tcell library to provide a tool that can run YarnSpinner scripts
// in a terminal.
package terminal

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/RemiEven/ysgo/internal/tree"
	"github.com/RemiEven/ysgo/markup"
	"github.com/RemiEven/ysgo/runner"
)

// Runner can run and display a dialogue in a terminal.
type Runner struct {
	dr       *runner.DialogueRunner
	app      *tview.Application
	textView *tview.TextView
	list     *tview.List
	choice   int
}

// NewRunner creates a new runner that will execute the dialogue contained in the
// given filename, with randomness based on rngSeed (if the dialogue uses it).
func NewRunner(filename string, rngSeed string) (*Runner, error) {
	di, err := tree.FromFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dialogue from file: %w", err)
	}

	dr, err := runner.NewDialogueRunner(di, nil, rngSeed)
	if err != nil {
		return nil, fmt.Errorf("failed to create dialogue runner: %w", err)
	}

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault

	r := &Runner{
		dr:       dr,
		app:      tview.NewApplication(),
		textView: tview.NewTextView(),
		list:     tview.NewList().ShowSecondaryText(false),
	}

	r.textView.SetDynamicColors(true)
	r.textView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			r.next()
		}
	})

	return r, nil
}

func (r *Runner) displayDialogueElement(element *runner.DialogueElement) {
	switch {
	case element.Line != nil:
		r.displayLine(element.Line)
	case len(element.Options) != 0:
		r.displayOptions(element.Options)
	default:
		panic("unknown type of dialogue element")
	}
}

func (r *Runner) displayLine(line *markup.ParseResult) {
	text := toStyledText(line)
	if characterAttr, ok := line.Attribute("character"); ok {
		text = "[yellow::i]" + text[:characterAttr.Length] + "[-::-]\n\t" + text[characterAttr.Length:]
	}
	r.textView.SetText(text)
	r.app.SetRoot(r.textView, true)
}

func (r *Runner) displayOptions(options []runner.DialogueOption) {
	r.list.Clear()
	for i, option := range options {
		if !option.Disabled {
			r.list = r.list.AddItem(option.Line.Text, "", 0, r.setChoice(i))
		}
	}
	r.app.SetRoot(r.list, true)
}

func (r *Runner) setChoice(choice int) func() {
	return func() {
		r.choice = choice
		r.next()
	}
}

func (r *Runner) next() {
	element, err := r.dr.Next(r.choice)
	if errors.Is(err, runner.ErrWaitingForCommandCompletion) {
		// do nothing, check again in 1/20th of a second
		time.Sleep(50 * time.Millisecond)
		r.next()
	} else if err != nil {
		r.app.Stop()
		log.Fatalf("error: %v", err)
	} else if element == nil {
		r.app.Stop()
	} else {
		r.displayDialogueElement(element)
	}
}

// Run starts the execution of the dialogue and its display.
// It returns once the dialogue has ended or when an error is encountered.
func (r *Runner) Run() error {
	r.next()
	if err := r.app.Run(); err != nil {
		return fmt.Errorf("failed to run application: %w", err)
	}
	return nil
}
