package view

import (
	"ChatClient/internal/ctrl"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func View(){
	myApp := app.New()
	w := myApp.NewWindow("Chat")
	UserNameText := widget.NewEntry()
	AddrText := widget.NewEntry()
	ctrl.ConnStatus = widget.NewLabel("fail")
	ctrl.InputText = widget.NewEntry()
	ConButton := widget.NewButton("ConnectionBtn", func() {
		username := UserNameText.Text
		addr := AddrText.Text
		if username ==""||addr=="" {

		}
		ctrl.ChatStart(username,addr)

	})

	ExitButton := widget.NewButton("ExitBtn", func() {
		username := UserNameText.Text
		context := "exit"
		ctrl.ChatSend(username,context)
		ctrl.ChatExit()
		ctrl.UserListLabel.Text =""
		ctrl.InputText.Text = ""
		ctrl.InputText.Refresh()
	})
	//用户列表
	ctrl.UserListLabel = widget.NewLabel("")

	//聊天内容
	ctrl.ChatLabel = widget.NewLabel("")
	SendButton := widget.NewButton("SendBtn", func() {
		username := UserNameText.Text
		context := ctrl.InputText.Text
		ctrl.ChatSend(username,context)
	})
	newForm := widget.NewForm(
		widget.NewFormItem("UserName:",UserNameText),
		widget.NewFormItem("Address:",AddrText),
		widget.NewFormItem("Connection:",ConButton),
		widget.NewFormItem("ConnStatus:",ctrl.ConnStatus),
		widget.NewFormItem("",ExitButton),
		widget.NewFormItem("UserList:",ctrl.UserListLabel),
		widget.NewFormItem("Chat:",ctrl.ChatLabel),
		widget.NewFormItem("Input:",ctrl.InputText),
		widget.NewFormItem("",SendButton),
	)
	w.SetContent(newForm)
	w.Resize(fyne.Size{
		Width: 500,
		Height: 400,
	})
	w.ShowAndRun()
}