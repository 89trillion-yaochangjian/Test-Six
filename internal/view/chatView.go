package view

import (
	"ChatClient/internal/ctrl"
	"ChatClient/internal/model"
	"ChatClient/internal/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func InitView() {
	myApp := app.New()
	w := myApp.NewWindow("Chat")
	model.UserNameText = widget.NewEntry()
	model.AddrText = widget.NewEntry()
	//连接状态OK：连接 fail：为连接
	model.ConnStatus = widget.NewLabel("fail")
	model.InputText = widget.NewEntry()
	ConButton := widget.NewButton("ConnectionBtn", func() {
		ctrl.ChatStart()
	})
	ExitButton := widget.NewButton("ExitBtn", func() {
		service.ChatExit()
	})
	//用户列表
	model.UserListLabel = widget.NewLabel("")
	//聊天内容
	model.ChatLabel = widget.NewLabel("")
	SendButton := widget.NewButton("SendBtn", func() {
		ctrl.ChatSend()
	})
	newForm := widget.NewForm(
		//用户名
		widget.NewFormItem("UserName:", model.UserNameText),
		//地址
		widget.NewFormItem("Address:", model.AddrText),
		//连接按钮
		widget.NewFormItem("Connection:", ConButton),
		//连接状态
		widget.NewFormItem("ConnStatus:", model.ConnStatus),
		//退出按钮
		widget.NewFormItem("", ExitButton),
		//用户列表
		widget.NewFormItem("UserList:", model.UserListLabel),
		//聊天框
		widget.NewFormItem("Chat:", model.ChatLabel),
		//输入框
		widget.NewFormItem("Input:", model.InputText),
		//发送
		widget.NewFormItem("", SendButton),
	)
	w.SetContent(newForm)
	w.Resize(fyne.Size{
		Width:  500,
		Height: 400,
	})
	w.ShowAndRun()
}
