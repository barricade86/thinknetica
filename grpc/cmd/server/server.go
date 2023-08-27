package main

import (
	"context"
	"log"
	"net"
	pb "thinknetica/grpc/messenger"

	"google.golang.org/grpc"
)

type Messages struct {
	Data []pb.Message

	pb.UnimplementedMessengerServer
}

func (m *Messages) Messages(_ *pb.Empty, stream pb.Messenger_MessagesServer) error {
	for i := range m.Data {
		stream.Send(&m.Data[i])
	}
	return nil
}

func (m *Messages) Send(_ context.Context, message *pb.Message) (*pb.Empty, error) {
	m.Data = append(m.Data, *message)
	return new(pb.Empty), nil
}

func main() {
	msg := &Messages{}
	listener, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessengerServer(grpcServer, msg)
	grpcServer.Serve(listener)
}
