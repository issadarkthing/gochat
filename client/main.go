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

	goChat.login.AddButton("login", func() {
		getValue := func(label string) string {
			val := goChat.login.GetFormItemByLabel(label).(*tview.InputField).GetText()
			return strings.TrimSpace(val)
		}

		client.username = getValue("Username: ")
		client.email = getValue("Email: ")
		goChat.input.SetLabel(fmt.Sprintf(" %s : ", client.username))
		goChat.app.SetRoot(goChat.window, true)
	})

	// websockets
	err := client.connect()
	if err != nil {
		panic(err)
	}

	client.receiveHandler(func(data structure.Message) {
		currText := goChat.text.GetText(true)
		message := fmt.Sprintf("%s: %s\n", data.Username, data.Message)
		goChat.text.SetText(currText+message)
		goChat.app.Draw()
	})

	goChat.input.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}
		client.send(goChat.input.GetText())
		goChat.input.SetText("")
	})

	if err := goChat.app.SetRoot(goChat.login, true).Run(); err != nil {
		panic(err)
	}
}
