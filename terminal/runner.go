package terminal

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/RemiEven/ysgo/markup"
	"github.com/RemiEven/ysgo/runner"
	"github.com/RemiEven/ysgo/tree"
)

type Runner struct {
	dr       *runner.DialogueRunner
	app      *tview.Application
	textView *tview.TextView
	list     *tview.List
	choice   int
}

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
	if characterAttr, ok := line.GetAttributeWithName("character"); ok {
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
	if err != nil {
		r.app.Stop()
		log.Fatalf("error: %v", err)
	} else if element == nil {
		r.app.Stop()
	} else {
		r.displayDialogueElement(element)
	}
}

func (r *Runner) Run() error {
	r.next()
	if err := r.app.Run(); err != nil {
		return fmt.Errorf("failed to run application: %w", err)
	}
	return nil
}
