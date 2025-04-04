package grpcserver

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
	plan_feature "mis-plan-features-hub/proto/plan-feature"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	Port       string
	Logger     interfaces.Logger
	Controller plan_feature.PlanFeatureServer
}

func (g GrpcServer) InitGrpcServer() {
	lis, err := net.Listen("tcp", ":"+g.Port)
	if err != nil {
		g.Logger.Error("failed to listen: " + err.Error())
	}

	grpcServer := grpc.NewServer()
	plan_feature.RegisterPlanFeatureServer(grpcServer, g.Controller)

	reflection.Register(grpcServer)

	g.Logger.Info("Servidor gRPC rodando na porta: " + g.Port)
	if err := grpcServer.Serve(lis); err != nil {
		g.Logger.Error("failed to serve: " + err.Error())
	}
}
