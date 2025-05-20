package grpc_client

import (
	"context"
	pf "github.com/kristrex/protos/gen/go/profile"
	"log"

	pb "github.com/kristrex/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	AuthClient    pb.AuthClient
	ProfileClient pf.ProfileClient
}

func NewGRPCClient(grpcAddr, profAddr string) *GrpcClient {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC service: %v", err)
	}
	connProf, err := grpc.Dial(profAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC service: %v", err)
	}
	return &GrpcClient{
		AuthClient:    pb.NewAuthClient(conn),
		ProfileClient: pf.NewProfileClient(connProf),
	}
}

func (c *GrpcClient) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return c.AuthClient.Register(ctx, req)
}

func (c *GrpcClient) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return c.AuthClient.Login(ctx, req)
}

func (c *GrpcClient) IsAdmin(ctx context.Context, req *pb.IsAdminRequest) (*pb.IsAdminResponse, error) {
	return c.AuthClient.IsAdmin(ctx, req)
}

func (c *GrpcClient) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return c.AuthClient.RefreshToken(ctx, req)
}

func (c *GrpcClient) Logout(ctx context.Context, req *pb.LogoutRequest) error {
	_, err := c.AuthClient.Logout(ctx, req)
	return err
}

func (c *GrpcClient) GetProfile(ctx context.Context, req *pf.GetProfileByIDRequest) error {
	_, err := c.ProfileClient.GetProfileByID(ctx, req)
	return err
}

func (c *GrpcClient) GetCatalog(ctx context.Context, req *pf.GetCatalogRequest) (*pf.GetCatalogResponse, error) {
	return c.ProfileClient.GetCatalog(ctx, req)
}
