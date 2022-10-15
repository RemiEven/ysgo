package terminal

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/RemiEven/ysgo/parser"
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

func NewRunner(filename string) (*Runner, error) {
	inputStream, err := antlr.NewFileStream("/home/remi/projects/ysgo/script.yarn")
	if err != nil {
		return nil, fmt.Errorf("failed to create file input stream: %w", err)
	}

	lexer := parser.NewYarnSpinnerLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewYarnSpinnerParser(stream)

	listener := &tree.ParserListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Dialogue())

	di := listener.Dialogue()

	dr := runner.NewDialogueRunner(di)

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault

	r := &Runner{
		dr:       dr,
		app:      tview.NewApplication(),
		textView: tview.NewTextView(),
		list:     tview.NewList().ShowSecondaryText(false),
	}

	r.textView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			r.next()
		}
	})

	return r, nil
}

func (r *Runner) displayDialogueElement(element *runner.DialogueElement) {
	switch {
	case element.Line != "":
		r.displayLine(element.Line)
	case len(element.Options) != 0:
		r.displayOptions(element.Options)
	default:
		panic("unknown type of dialogue element")
	}
}

func (r *Runner) displayLine(line string) {
	r.textView.SetText(line)
	r.app.SetRoot(r.textView, true)
}

func (r *Runner) displayOptions(options []string) {
	r.list.Clear()
	for i, option := range options {
		r.list = r.list.AddItem(option, "", 0, r.setChoice(i))
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
	element, ok := r.dr.Next(r.choice)
	if !ok {
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
