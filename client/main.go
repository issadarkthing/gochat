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
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/issadarkthing/gochat/structure"
	"github.com/rivo/tview"
)

var (
	goChat Gui
)

func main() {

	client := Client{
		url: "ws://localhost:8080/ws",
	}

	goChat = newGui()

	// for login prompt
	goChat.login.AddButton("login", func() {
		getValue := func(label string) string {
			val := goChat.login.GetFormItemByLabel(label).(*tview.InputField).GetText()
			return strings.TrimSpace(val)
		}

		client.username = getValue("Username: ")
		label := fmt.Sprintf("[%s] %s :[white] ", client.color, client.username)
		goChat.input.SetLabel(label)
		goChat.app.SetRoot(goChat.window, true)
	})

	// create a connection
	err := client.connect()
	if err != nil {
		panic(err)
	}

	client.receiveHandler(func(data structure.Message) {
		// gets current text from textview and simply append it with incoming message
		currText := goChat.text.GetText(false)
		message := fmt.Sprintf("[%s]%s[white]: %s", 
			data.Color, data.Username, data.Message)
		goChat.text.SetText(currText+message)
		goChat.app.Draw()
	})

	goChat.input.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}
		// send message if enter was pressed
		client.send(goChat.input.GetText())
		// clear the input bar
		goChat.input.SetText("")
	})

	// display login prompt first
	goChat.app.SetRoot(center(goChat.login, 40, 10), true)
	goChat.app.SetFocus(goChat.login)
	// main loop goes here
	if err = goChat.app.Run(); err != nil {
		panic(err)
	}
}
