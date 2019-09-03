package manager

import "gprcStream/src/protobuf"

type MessageManager struct {
	registerFunc map[int32]func(message *protobuf.ProtoMessage) *protobuf.ProtoMessage
}

func (mgr *MessageManager) OnMessage(message *protobuf.ProtoMessage) *protobuf.ProtoMessage {
	callFunc := mgr.registerFunc[message.MessageID]
	if callFunc != nil {
		return callFunc(message)
	}

	if message.MessageType == protobuf.SendMessageType_RequestClient {
		//
	}

	return &protobuf.ProtoMessage{TargetClient: "null", Sn: message.Sn + 1, Body: nil}
}
