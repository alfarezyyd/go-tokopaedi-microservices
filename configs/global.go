package configs

import "google.golang.org/grpc"

const CustomTimeFormat = "02 Jan 2006 15:04:05"

type GrpcConnectionClient struct {
	EndpointTarget string
	DialOption     grpc.DialOption
}

type ClientInitiation func()
