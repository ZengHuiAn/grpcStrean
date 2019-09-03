package Services

import (
	"context"
	"fmt"
	"google.golang.org/grpc/peer"
	"gprcStream/src/protobuf"
	"net"
	"strings"
)

type ServiceRegistServer struct {
	clients map[string][]string
	TransmissionServer
}

func NewServiceRegistServer() *ServiceRegistServer {
	return &ServiceRegistServer{clients:make(map[string][]string)}
}

func (s *ServiceRegistServer) RegistService(ctx context.Context, proto *protobuf.RegisterServiceRequest) (*protobuf.RegisterServiceResponse, error) {
	var status int32 = 200
	clientInfo, _ := GetClientIP(ctx)
	fmt.Println(proto, clientInfo)
	if proto.RegType == protobuf.RegisterType_UnRegister {
		if s.clients != nil {
			delete(s.clients, proto.ClientName)
		}

		return &protobuf.RegisterServiceResponse{
			Status:  status,
			MsgType: protobuf.MessageType_Response,
		}, nil
	}
	if s.clients == nil {
		s.clients = make(map[string][]string)
	}

	if s.clients[proto.ClientName] != nil {
		status = 400
	} else {
		//
		s.clients[proto.ClientName] = []string{clientInfo[0], clientInfo[1]}
	}
	return &protobuf.RegisterServiceResponse{
		Status:  status,
		MsgType: protobuf.MessageType_Response,
	}, nil
}

func GetClientIP(ctx context.Context) ([]string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("[getClinetIP] invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return nil, fmt.Errorf("[getClientIP] peer.Addr is nil")
	}

	addSlice := strings.Split(pr.Addr.String(), ":")
	if addSlice[0] == "[" {
		//本机地址
		return []string{"localhost", "10002"}, nil
	}
	return addSlice, nil
}
