package Services

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gprcStream/src/protobuf"
)

type ServiceRegistClient struct {
	getClient protobuf.StreamServiceClient
}

func NewServiceRegistClient(rawConn *grpc.ClientConn) *ServiceRegistClient {
	return &ServiceRegistClient{getClient: protobuf.NewStreamServiceClient(rawConn)}
}

func (client *ServiceRegistClient) RegisteClient(req *protobuf.RegisterServiceRequest) {
	fmt.Println(req)
	resp, _ := client.getClient.RegistService(context.Background(), req)
	fmt.Println("结果:", resp)
}
