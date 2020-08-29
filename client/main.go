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
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/issadarkthing/gochat/structure"
	"github.com/rivo/tview"
)

func main() {

	// very useful for debugging
	file, err := os.OpenFile("./log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	log.SetOutput(file)

	client := Client{
		url: "wss://gochat-app.herokuapp.com/ws",
	}

	gui := newGui()
	bot := Bot{gui: gui, client: &client}

	// for login prompt
	gui.login.AddButton("login", func() {
		getValue := func(label string) string {
			val := gui.login.GetFormItemByLabel(label).(*tview.InputField).GetText()
			return strings.TrimSpace(val)
		}

		client.username = getValue("Username: ")
		client.passphrase = getValue("Passphrase: ")
		label := fmt.Sprintf("[%s] %s :[white] ", client.color, client.username)
		gui.input.SetLabel(label)
		gui.app.SetRoot(gui.window, true)
	})

	// create a connection
	err = client.connect()
	if err != nil {
		panic(err)
	}

	client.receiveHandler(func(data structure.Message) {
		// gets current text from textview and simply append it with incoming message
		currText := gui.text.GetText(false)

		message := fmt.Sprintf("[%s]%s[white]: %s", 
			data.Color, data.Username, data.Message)

		gui.text.SetText(currText+message)
		gui.app.Draw()
	})

	gui.input.SetDoneFunc(func(key tcell.Key) {

		if key != tcell.KeyEnter {
			return
		}

		text := gui.input.GetText()

		// if it was command
		if strings.HasPrefix(text, "/") {
			bot.messageHandler(text)
		} else {
			// send message if enter was pressed
			client.send(text)
		}

		// clear the input bar
		gui.input.SetText("")
	})

	// display login prompt first
	gui.app.SetRoot(center(gui.login, 40, 10), true)
	gui.app.SetFocus(gui.login)
	// main loop goes here
	if err = gui.app.Run(); err != nil {
		panic(err)
	}
}
