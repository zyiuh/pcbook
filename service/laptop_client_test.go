package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zyiuh/pcbook/pb"
	"github.com/zyiuh/pcbook/sample"
	"github.com/zyiuh/pcbook/serializer"
	"github.com/zyiuh/pcbook/service"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// 检查数据是否存储在内存中
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// 比较两个对象是否一致
	requireSameLaptop(t, laptop, other)
}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	// 注册服务
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // 随机可用端口
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
