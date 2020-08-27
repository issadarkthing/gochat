package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Gui struct {
	app    *tview.Application
	input  *tview.InputField
	text   *tview.TextView
	login  *tview.Form
	window *tview.Flex
}

func newGui() Gui {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
	gui := Gui{
		app: tview.NewApplication(),
		input: newInputField(),
		text: newTextPanel(),
		login: newLoginForm(),
		window: createWindow(),
	}

	gui.window.
		AddItem(gui.text, 0, 1, false).
		AddItem(gui.input, 1, 1, true)

	gui.window.SetBorderPadding(0, 0, 1, 1)
	return gui
}

func newLoginForm() *tview.Form {
	form := tview.NewForm()
	form.
		AddInputField("Username: ", "", 20, nil, nil).
		AddInputField("Email: ", "", 20, nil, nil)

	return form
}

func createWindow() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	return flex
}

func newTextPanel() *tview.TextView {
	text := tview.NewTextView()
	text.SetBorder(true)
	return text
}

func newInputField() *tview.InputField {
	input := tview.NewInputField()
	input.SetFieldBackgroundColor(tcell.ColorDefault)
	return input
}
