// Copyright (C) 2020 Raziman Mahathir

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
		AddInputField("Username: ", "", 20, nil, nil)

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
