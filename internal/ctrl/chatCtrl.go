package ctrl

import (
	"ChatClient/internal/log"
	"ChatClient/internal/model"
	"ChatClient/internal/wsClient"
	"fyne.io/fyne/v2/widget"
)

var UserListLabel *widget.Label

var ChatLabel *widget.Label

var ConnStatus *widget.Label

var InputText *widget.Entry


//发起连接,开始聊天

func ChatStart(userName string,addr string) {
		if userName == ""||addr == ""{
			ChatLabel.Text = model.CheckPra
			return
		}
		if wsClient.Conn !=nil{
			ChatLabel.Text = model.RepeatCon
			return
		}
		err := wsClient.ChatCon(userName,addr)
		if err!=nil {
			log.Error.Println(err)
			return
		}
		ConnStatus.Text = model.OK
		ChatLabel.Text = ""
		ChatSend(userName,model.SignIn)
		go ChatReceive()

}
//发送消息

func ChatSend(username string,context string) {
	if wsClient.Conn == nil {
		ChatLabel.Text = model.FisCon
		return
	}
	if wsClient.Conn !=nil{
		if context==model.ExitType {
			message := &model.ChatRequest{
				UserName: username,
				Type: context,
				Content: context,
			}
			wsClient.WriteMessage(message)
		} else {
			log.Info.Print(model.TalkLog,"user:"+username)
			message := &model.ChatRequest{
				UserName: username,
				Type: model.TalkType,
				Content: context,
		}
		wsClient.WriteMessage(message)
		}
	}
	InputText.Text = ""
	InputText.Refresh()
}
//接受消息

func ChatReceive() {
	for {
		msg, err := wsClient.ReadMessage()
		if err != nil {
			break
		}
		//读取用户列表
		var userList string
		if msg.UserList != nil {
			for _, value := range msg.UserList {
				userList += value + "\n"
			}
			UserListLabel.Text = userList
			UserListLabel.Refresh()
		}
		//读取聊天内容
		if msg.Content != ""&& msg.UserName !=""{
			ChatLabel.Text += msg.UserName+":"+msg.Content+"\n"
			ChatLabel.Refresh()
		}
	}
}

//退出连接

func ChatExit() {
	if wsClient.Conn == nil {
		ChatLabel.Text = model.FisCon
		return
	}
	if wsClient.Conn != nil{
		err := wsClient.Exit()
		if err != nil {
			log.Error.Println(err)
		}
		wsClient.Conn = nil
		ConnStatus.Text = model.Fail
		ChatLabel.Text = ""
	}
}
