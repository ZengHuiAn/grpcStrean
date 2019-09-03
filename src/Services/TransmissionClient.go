package Services

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gprcStream/src/protobuf"
)

type TransmissionClient struct {
	getClient protobuf.TCPServiceClient
}

func NewTransmissionClient(rawConn *grpc.ClientConn) *TransmissionClient {
	return &TransmissionClient{getClient: protobuf.NewTCPServiceClient(rawConn)}
}

func (client *TransmissionClient) SendMessage(ctx context.Context, message *protobuf.ProtoMessage) (*protobuf.ProtoMessage, error) {
	fmt.Println("发送数据", message)
	msg, err := client.getClient.ServiceProgress(ctx, message)
	fmt.Println("发送返回值", msg, err)
	return msg, err
}