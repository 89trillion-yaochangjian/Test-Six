package wsClient

import (
	"ChatClient/internal/config"
	"ChatClient/internal/model"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
)

// Conn 用户连接
var Conn *websocket.Conn

//发起连接

func ChatCon(userName string, addr string) error {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	var dialer *websocket.Dialer
	header := http.Header{}
	header.Set("username", userName)
	//通过Dialer连接websocket服务器
	conn, _, err := dialer.Dial(u.String(), header)
	if err != nil {
		config.Error.Println(err)
		return err
	}
	Conn = conn
	return err
}

//写消息

func WriteMessage(msg *model.ChatRequest) error {
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		config.Error.Println(err)
		return err
	}
	err1 := Conn.WriteMessage(websocket.TextMessage, msgMarshal)
	return err1
}

//读消息

func ReadMessage() (*model.ChatRequest, error) {
	if Conn != nil {
		_, rMsg, err := Conn.ReadMessage()
		if err != nil {
			return nil, err
		}
		message := &model.ChatRequest{}
		err1 := proto.Unmarshal(rMsg, message)
		return message, err1
	}
	return nil, nil
}

//退出连接

func Exit() error {
	err := Conn.WriteMessage(websocket.CloseMessage, nil)
	if err != nil {
		config.Error.Println(err)
		return err
	}
	err = Conn.Close()
	if err != nil {
		config.Error.Println(err)
		return err
	}
	return err
}
