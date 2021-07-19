package service

import (
	"ChatClient/internal/log"
	"ChatClient/internal/model"
	"ChatClient/internal/wsClient"
)

//exit处理

func ChatExit(username string, context string) {
	message := &model.ChatRequest{
		UserName: username,
		Type:     context,
		Content:  context,
	}
	wsClient.WriteMessage(message)
	if wsClient.Conn == nil {
		model.ChatLabel.Text = model.FisCon
		return
	}
	if wsClient.Conn != nil {
		err := wsClient.Exit()
		if err != nil {
			log.Error.Println(err)
		}
		wsClient.Conn = nil
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
	wsClient.WriteMessage(message)
}

//talk 处理

func ChatTalk(username string, context string) {
	log.Info.Print(model.TalkLog, "user:"+username)
	message := &model.ChatRequest{
		UserName: username,
		Type:     model.TalkType,
		Content:  context,
	}
	wsClient.WriteMessage(message)
}
