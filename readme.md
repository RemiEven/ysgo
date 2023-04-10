# Ysgo

[![Go Reference](https://pkg.go.dev/badge/github.com/remieven/ysgo.svg)](https://pkg.go.dev/github.com/remieven/ysgo)
![build](https://github.com/RemiEven/ysgo/actions/workflows/onPush.yml/badge.svg)
[![codecov](https://codecov.io/gh/RemiEven/ysgo/branch/main/graph/badge.svg)](https://codecov.io/gh/RemiEven/ysgo)

## Introduction

Ysgo is a Go library that is able to run YarnSpinner dialogues.

It hasn't reached V1 yet but its public API shouldn't undergo any massive change now.

## What's YarnSpinner?

[YarnSpinner](https://yarnspinner.dev/) is an open-source tool for creating interactive dialogue and branching narratives in video games and other interactive media. It allows writers and game designers to create complex branching stories, character interactions, and dialogues, and provides an [intuitive visual interface](https://marketplace.visualstudio.com/items?itemName=SecretLab.yarn-spinner) for scripting and organizing the flow of the story. YarnSpinner is often used in the creation of visual novels, adventure games, and other interactive narrative experiences.

## Minimal example

Here's a small program that runs a dialogue, printing it to the standard output:

```go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/remieven/ysgo"
)

func main() {
	script := `
title: flavors
---
Which flavor do you prefer?
-> Chocolate
	<<set $flavor to "chocolate">>
-> Raspberry
	<<set $flavor to "raspberry">>
Nice! Here, get some {$flavor} icecream!
===`

	logFatalIfErr := func(err error) {
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	// NewDialogueRunner parses the dialogue and creates a runner that can execute it
	dr, err := ysgo.NewDialogueRunner(nil, "", strings.NewReader(script))
	logFatalIfErr(err)

	// Next advances the dialogue to the next step
	dialogueElement, err := dr.Next(0)
	logFatalIfErr(err)
	fmt.Println(dialogueElement.Line.Text) // Which flavor do you prefer?

	dialogueElement, err = dr.Next(0)
	logFatalIfErr(err)
	fmt.Println(dialogueElement.Options[0].Line.Text) // Chocolate
	fmt.Println(dialogueElement.Options[1].Line.Text) // Raspberry

	dialogueElement, err = dr.Next(1) // Nice! Here, get some raspberry icecream!
	logFatalIfErr(err)
	fmt.Println(dialogueElement.Line.Text)
}
```

## Features

YarnSpinner's core features are mostly supported by Ysgo, with the exception of localization (which might be added sometime) and type checking (not currently planned).

In addition to these missing features, there are key differences between Ysgo and YarnSpinner:

- No compilation: parsing occurs at runtime
- No built-in integration with any game engines (however, it works well with Ebitengine)
- Dialogue runners are only in "pull" mode, meaning they won't emit events or control UI components on their own.

Ysgo only handles your dialogues when your game is running, during development you're still encouraged to use the vscode plugin and YarnSpinner-Console.

## Ysgo in Action

Despite not having all of YarnSpinner's features, Ysgo still provides enough functionality to recreate sample projects from the YarnSpinner documentation.

Check out the demos:

- *Choose-Your-Path Game*: [doc link](https://docs.yarnspinner.dev/unity-sample-projects/example-project-1), [demo link](https://remieven.itch.io/ysgo-sample-cubeandsphere)

Source code for all demos is available in [the ysgo-examples Github repository](https://github.com/RemiEven/ysgo-examples).

## Getting help

Feel free to open an issue in this repository!

## License

Ysgo is licensed under the MIT license (see the `LICENSE.txt` file).

## Credits

Ysgo wouldn't exist if it weren't for all the people behind Yarn Spinner. Big thanks to them.

Also, [this article](https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/) has been super helpful to figure out how to use antlr4 with Go.
