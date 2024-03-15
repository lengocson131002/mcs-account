package grpc

import (
	pb "github.com/lengocson131002/mcs-account/presentation/grpc/pb"
	"github.com/lengocson131002/mcs-account/presentation/grpc/server"
)

func (g *GrpcServer) WithAccountServer() GrpcServerStartOption {
	return func(s *GrpcServer) error {
		tSrv := server.NewAccountServer()
		pb.RegisterAccountServiceServer(s.gSrv, tSrv)
		return nil
	}
}
