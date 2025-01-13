package notification

import (
	"errors"
	pb "github.com/azeezlala/assessment/shared/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type Client struct {
	connection *grpc.ClientConn
	client     pb.NotificationServiceClient
}

func New() (*Client, error) {
	c := new(Client)
	err := c.Start()
	if err != nil {
		return nil, errors.New("could not start client")
	}

	return c, nil
}

func (c *Client) Start() error {
	var err error

	c.connection, err = grpc.Dial(os.Getenv("GRPC_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	c.client = pb.NewNotificationServiceClient(c.connection)
	return err
}

func (c *Client) Stop() error {
	return c.connection.Close()
}
