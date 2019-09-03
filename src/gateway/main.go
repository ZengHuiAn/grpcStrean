package main

import (
	"google.golang.org/grpc"
	"gprcStream/src/Services"
	"gprcStream/src/protobuf"
	"log"
	"net"
)

var (
	host = "127.0.0.1"
	port = "10002"
)

func main() {

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	registerSV := Services.NewServiceRegistServer()
	protobuf.RegisterStreamServiceServer(s, registerSV)
	service := Services.NewTransmissionServer(&OnTransmissonMessage{})
	//client := Services.NewTransmissionClient(lis)
	protobuf.RegisterTCPServiceServer(s, service)
	//protobuf.RegisterStreamServiceServer(s, service)
	_ = s.Serve(lis)
}

type OnTransmissonMessage struct {
}

func (*OnTransmissonMessage) OnMessage(message *protobuf.ProtoMessage) *protobuf.ProtoMessage {
	return &protobuf.ProtoMessage{TargetClient: "null"}
}
