package configs

import (
	"go-tokopaedi-microservices/protobuf/genproto"

	"google.golang.org/grpc"
)

type GrpcClientRegistry struct {
	Registry map[string]interface{}
}

func NewGrpcClientRegistry() *GrpcClientRegistry {
	return &GrpcClientRegistry{
		Registry: make(map[string]interface{}),
	}
}

func (grpcClientRegistry *GrpcClientRegistry) Connect(serviceName, endpointTarget string, dialOption grpc.DialOption) error {
	grpcConn, err := grpc.NewClient(endpointTarget, dialOption)
	if err != nil {
		return err
	}
	userGrpcClient := genproto.NewUserServiceClient(grpcConn)
	grpcClientRegistry.Registry[serviceName] = userGrpcClient
	return nil
}

func (grpcClientRegistry *GrpcClientRegistry) Get(serviceName string) (interface{}, error) {
	if _, ok := grpcClientRegistry.Registry[serviceName]; ok {
		return grpcClientRegistry.Registry[serviceName], nil
	}

	return nil, nil
}
