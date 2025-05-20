package auth

import (
	"auth/internal/services/auth"
	"context"
	"errors"
	"github.com/google/uuid"
	ssov1 "github.com/kristrex/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (accessToken, refreshToken string, err error)
	Register(
		ctx context.Context,
		name string,
		lastname string,
		email string,
		goal string,
		password string,
	) (userID uuid.UUID, err error)
	IsAdmin(
		ctx context.Context,
		userID string,
	) (bool, error)
	Logout(
		ctx context.Context,
		refreshToken string,
	) error
	UpdateTokens(
		ctx context.Context,
		refreshTokenID string,
		appID int,
	) (string, string, error)
	Session(
		ctx context.Context,
		accessToken string,
	) (string, string, string, string, bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

const (
	emptyValue       = 0
	emptyValueString = ""
)

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}

	// implement login via auth service
	accessToken, refreshToken, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &ssov1.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *serverAPI) Register(ctx context.Context, in *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if err := validateRegister(in); err != nil {
		return nil, err
	}

	uid, err := s.auth.Register(ctx, in.GetName(), in.GetLastName(), in.GetEmail(), in.GetGoal(), in.GetPassword())
	if err != nil {
		if strings.Contains(err.Error(), "user already exists") {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "failed to register user")
	}
	return &ssov1.RegisterResponse{UserId: uid.String()}, nil
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	if err := validateIsAdmin(req); err != nil {
		return nil, err
	}
	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to check admin status")
	}

	return &ssov1.IsAdminResponse{IsAdmin: isAdmin}, nil
}

func (s *serverAPI) RefreshToken(ctx context.Context, req *ssov1.RefreshTokenRequest) (*ssov1.RefreshTokenResponse, error) {
	if err := validateRefreshToken(req); err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := s.auth.UpdateTokens(ctx, req.RefreshToken, 1)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update tokens")
	}
	return &ssov1.RefreshTokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *serverAPI) Logout(ctx context.Context, req *ssov1.LogoutRequest) (*emptypb.Empty, error) {
	if err := validateLogout(req); err != nil {
		return nil, err
	}

	if err := s.auth.Logout(ctx, req.GetAccessToken()); err != nil {
		return nil, status.Error(codes.Internal, "failed to logout")
	}
	return &emptypb.Empty{}, nil
}

func (s *serverAPI) Session(ctx context.Context, req *ssov1.SessionRequest) (*ssov1.SessionResponse, error) {
	if err := validateSession(req); err != nil {
		return nil, err
	}
	id, name, lastName, role, isAdmin, err := s.auth.Session(ctx, req.GetAccessToken())
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		return nil, status.Error(codes.Internal, "failed to check session")
	}

	return &ssov1.SessionResponse{UserId: id, FirstName: name, LastName: lastName, Role: role, IsAdmin: isAdmin}, nil
}

func validateLogin(req *ssov1.LoginRequest) error {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "requires email or password")
	}

	if req.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "login requires app id")
	}
	return nil
}

func validateRegister(req *ssov1.RegisterRequest) error {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "requires email or password")
	}

	if req.GetGoal() != "scout" && req.GetGoal() != "model" {
		return status.Error(codes.InvalidArgument, "requires goal")
	}

	if req.GetName() == "" || req.GetLastName() == "" {
		return status.Error(codes.InvalidArgument, "requires name")
	}
	if req.GetLastName() == "" {
		return status.Error(codes.InvalidArgument, "requires last name")
	}
	return nil
}

func validateIsAdmin(req *ssov1.IsAdminRequest) error {
	if req.GetUserId() == emptyValueString {
		return status.Error(codes.InvalidArgument, "requires user id")
	}
	return nil
}

func validateRefreshToken(req *ssov1.RefreshTokenRequest) error {
	if req.GetRefreshToken() == emptyValueString {
		return status.Error(codes.InvalidArgument, "requires refresh token")
	}
	return nil
}

func validateLogout(req *ssov1.LogoutRequest) error {
	if req.GetAccessToken() == emptyValueString {
		return status.Error(codes.InvalidArgument, "requires refresh token")
	}
	return nil
}

func validateSession(req *ssov1.SessionRequest) error {
	if req.GetAccessToken() == emptyValueString {
		return status.Error(codes.InvalidArgument, "requires session token")
	}
	return nil
}
