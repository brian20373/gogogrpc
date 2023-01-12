package abcd

import "github.com/brian20373/gogogrpc/gen/pb/gogogrpc"

func FixedPayload() *gogogrpc.HelloRequest {
	return &gogogrpc.HelloRequest{
		Name: "Brian",
	}
}
