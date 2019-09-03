package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gprcStream/src/Services"
	"gprcStream/src/protobuf"
	sig "gprcStream/src/signal"
	"os"
	"time"
)

var (
	targetHost = "127.0.0.1"
	targetPort = "10002"
)

func main() {
	conn, err := grpc.Dial(targetHost+":"+targetPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	client := Services.NewServiceRegistClient(conn)
	streamClient := Services.NewTransmissionClient(conn)
	client.RegisteClient(&protobuf.RegisterServiceRequest{ClientName: "view", RegType: protobuf.RegisterType_Register})
	var oldMsg = &protobuf.ProtoMessage{TargetClient: "view", MessageType: protobuf.SendMessageType_ResponseClient, Body: []byte{1, 2, 3}}
	for {

		oldMsg, err = streamClient.SendMessage(context.Background(), oldMsg)
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
	sig.Wait(os.Interrupt, os.Kill)

	client.RegisteClient(&protobuf.RegisterServiceRequest{ClientName: "view", RegType: protobuf.RegisterType_UnRegister})
}
