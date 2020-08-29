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

import "fmt"



type Bot struct {
	gui  Gui
}


func (b Bot) messageHandler(command string) {
	switch command {
	case "/help":
		b.printHelp()
	}
}

func (b Bot) printHelp() {
	
	helpMessage := `Hi, I'm a simple bot. Here are the commands that may help you in gochat
┌────────────────────────────────────────────────┐
│        command               description       │
├─────────────────────────┬──────────────────────┤
│/members                 │ show number of users │
│/nick       <name>       │ change username      │
│/passphrase <passphrase> │ change passphrase    │
│/help                    │ print help message   │
│/exit                    │ exit gochat          │
└─────────────────────────┴──────────────────────┘
`
	b.print(helpMessage)
}

func (b Bot) print(text string) {
	// get current text on the text panel
	currText := b.gui.text.GetText(false)
	message := fmt.Sprintf("[grey]bot: %s[white]", text)
	b.gui.text.SetText(currText+message)
}
