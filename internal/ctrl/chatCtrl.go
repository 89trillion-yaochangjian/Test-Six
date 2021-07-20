package ctrl

import (
	"ChatClient/internal/config"
	"ChatClient/internal/model"
	"ChatClient/internal/service"
	"ChatClient/internal/wsClient"
)

//发起连接,开始聊天

func ChatStart(userName string, addr string) {
	if userName == "" || addr == "" {
		model.ChatLabel.Text = model.CheckPra
		return
	}
	if wsClient.Conn != nil {
		model.ChatLabel.Text = model.RepeatCon
		return
	}
	err := wsClient.ChatCon(userName, addr)
	if err != nil {
		config.Error.Println(err)
		return
	}
	model.ConnStatus.Text = model.OK
	model.ChatLabel.Text = ""
	ChatSend(userName, model.SignIn)
	go ChatReceive()
	//wsClient.Sender(wsClient.Conn)
}

//发送消息

func ChatSend(username string, context string) {
	if wsClient.Conn == nil {
		model.ChatLabel.Text = model.FisCon
		return
	}
	if wsClient.Conn != nil {
		if context == model.ExitType {
			service.ChatExit(username, context)
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
		msg, err := wsClient.ReadMessage()
		if err != nil {
			model.ConnStatus.Text = "fail"
			model.UserListLabel.Text = ""
			model.ChatLabel.Text = ""
			model.ConnStatus.Refresh()
			model.UserListLabel.Refresh()
			model.ChatLabel.Refresh()
			wsClient.Conn = nil
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
