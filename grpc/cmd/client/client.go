package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "thinknetica/grpc/messenger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewMessengerClient(conn)
	client.Send(context.Background(), &pb.Message{Id: 1, Text: "You wanna know the secret of pain?", DeliveryDate: "2023-08-25"})
	client.Send(context.Background(), &pb.Message{Id: 2, Text: "If you just stop feeling it, you can start using it", DeliveryDate: "2023-08-26"})
	client.Send(context.Background(), &pb.Message{Id: 3, Text: "I'll get you, my pretty! And your little soul, too!", DeliveryDate: "2023-08-27"})
	err = getAllMessagesFromserver(context.Background(), client)
	if err != nil {
		fmt.Println("error=", err)
	}
}

func getAllMessagesFromserver(ctx context.Context, client pb.MessengerClient) error {
	stream, err := client.Messages(context.Background(), &pb.Empty{})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			book, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			fmt.Printf("Message: %v\n", book)
		}
	}
}
