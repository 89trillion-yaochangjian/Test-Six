package service

import (
	"ChatClient/internal/config"
	"ChatClient/internal/model"
	"ChatClient/internal/status"
	"ChatClient/internal/ws"
)

//exit处理

func ChatExit() {
	if ws.Conn == nil {
		model.ChatLabel.Text = status.CheckPra.Msg
		return
	}
	username := model.UserNameText.Text
	context := "exit"
	if ws.Conn != nil {
		message := &model.ChatRequest{
			UserName: username,
			Type:     context,
			Content:  context,
		}
		ws.WriteMessage(message)
		if ws.Conn == nil {
			model.ChatLabel.Text = status.FisCon.Msg
			return
		}
		err := ws.Exit()
		if err != nil {
			config.Error.Println(err)
		}
		ws.Conn = nil
		model.ConnStatus.Text = model.Fail
		model.ChatLabel.Text = ""
	}
	model.UserListLabel.Text = ""
	model.InputText.Text = ""
	model.InputText.Refresh()

}

//user list处理

func ChatUserList(username string) {
	message := &model.ChatRequest{
		UserName: username,
		Type:     model.UserListType,
	}
	ws.WriteMessage(message)
}

//talk 处理

func ChatTalk(username string, context string) {
	config.Info.Print(model.TalkLog, "user:"+username)
	message := &model.ChatRequest{
		UserName: username,
		Type:     model.TalkType,
		Content:  context,
	}
	ws.WriteMessage(message)
}
