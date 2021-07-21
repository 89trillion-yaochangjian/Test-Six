package ctrl

import (
	"ChatClient/internal/config"
	"ChatClient/internal/model"
	"ChatClient/internal/service"
	"ChatClient/internal/status"
	"ChatClient/internal/ws"
)

//发起连接,开始聊天

func ChatStart() {
	userName := model.UserNameText.Text
	addr := model.AddrText.Text
	if userName == "" || addr == "" {
		model.ChatLabel.Text = status.CheckPra.Msg
		return
	}
	if ws.Conn != nil {
		model.ChatLabel.Text = status.RepeatCon.Msg
		return
	}
	err := ws.ChatCon(userName, addr)
	if err != nil {
		config.Error.Println(err)
		return
	}
	model.ConnStatus.Text = model.OK
	model.ChatLabel.Text = ""
	model.InputText.Text = status.SignIn.Msg
	ChatSend()
	go ChatReceive()
}

//发送消息

func ChatSend() {
	username := model.UserNameText.Text
	context := model.InputText.Text
	model.InputText.Refresh()
	if ws.Conn == nil {
		model.ChatLabel.Text = status.FisCon.Msg
		return
	}
	if ws.Conn != nil {
		if context == model.ExitType {
			service.ChatExit()
		} else if context == model.UserListType {
			service.ChatUserList(username)
		} else {
			service.ChatTalk(username, context)
		}
	}
	model.InputText.Text = ""
	model.InputText.Refresh()
}

//接受消息

func ChatReceive() {
	for {
		msg, err := ws.ReadMessage()
		if err != nil {
			model.ConnStatus.Text = "fail"
			model.UserListLabel.Text = ""
			model.ChatLabel.Text = ""
			model.ConnStatus.Refresh()
			model.UserListLabel.Refresh()
			model.ChatLabel.Refresh()
			ws.Conn = nil
			break
		}
		//读取用户列表
		var userList string
		if msg.UserList != nil {
			for _, value := range msg.UserList {
				userList += value + "\n"
			}
			model.UserListLabel.Text = userList
			model.UserListLabel.Refresh()
		}
		//读取聊天内容
		if msg.Content != "" && msg.UserName != "" {
			model.ChatLabel.Text += msg.UserName + ":" + msg.Content + "\n"
			model.ChatLabel.Refresh()
		}
	}
}
