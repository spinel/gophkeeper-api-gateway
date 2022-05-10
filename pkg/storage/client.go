package storage

import (
	"fmt"

	"github.com/spinel/gophkeeper-api-gateway/pkg/config"
	"github.com/spinel/gophkeeper-api-gateway/pkg/storage/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.StorageServiceClient
}

func InitServiceClient(c *config.Config) pb.StorageServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.StorageSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewStorageServiceClient(cc)
}
