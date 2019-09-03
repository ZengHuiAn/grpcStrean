package Services

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gprcStream/src/protobuf"
)

type IMessage interface {
	OnMessage(message *protobuf.ProtoMessage) *protobuf.ProtoMessage
}

type TransmissionServer struct {
	readBuff chan *protobuf.ProtoMessage
	onMsg    IMessage
}

func NewTransmissionServer(callFunc IMessage) *TransmissionServer {
	return &TransmissionServer{readBuff: make(chan *protobuf.ProtoMessage, 1), onMsg: callFunc}
}

// 客户端向服务器发送数据
func (self *TransmissionServer) ServiceProgress(ctx context.Context, message *protobuf.ProtoMessage) (*protobuf.ProtoMessage, error) {
	self.readBuff <- message

	var response *protobuf.ProtoMessage = nil
	select {
	case value, ok := <-self.readBuff:
		if ok {
			if self.onMsg != nil && self.onMsg.OnMessage != nil {
				response = self.onMsg.OnMessage(value)
			}
		}
	}
	fmt.Println("服务器收到数据", message, "返回数据：", response)
	return response, nil
}


