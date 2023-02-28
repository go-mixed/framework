package grpc

import (
	"context"
	"google.golang.org/grpc"
	"gopkg.in/go-mixed/framework.v1/container"
	cgrpc "gopkg.in/go-mixed/framework.v1/contracts/grpc"
)

func getGrpc() cgrpc.IGrpc {
	return container.MustMake[cgrpc.IGrpc]("grpc")
}

func Run(host string) error {
	return getGrpc().Run(host)
}
func Server() *grpc.Server {
	return getGrpc().Server()
}
func Client(ctx context.Context, name string) (*grpc.ClientConn, error) {
	return getGrpc().Client(ctx, name)
}
func UnaryServerInterceptors(v []grpc.UnaryServerInterceptor) {
	getGrpc().UnaryServerInterceptors(v)
}
func UnaryClientInterceptorGroups(v map[string][]grpc.UnaryClientInterceptor) {
	getGrpc().UnaryClientInterceptorGroups(v)
}
