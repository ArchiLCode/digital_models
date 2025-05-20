package profile

import (
	"context"
	"errors"
	"profile/internal/domain/models"
	"time"

	profile1 "github.com/kristrex/protos/gen/go/profile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Profile interface {
	CreateProfile(
		ctx context.Context,
		name string,
		lastName string,
		userId string,
	) error

	GetProfileByID(
		ctx context.Context,
		id string,
	) (models.GetProfileByIDResponse, error)

	UpdateProfile(
		ctx context.Context,
		req models.UpdateProfileRequest,
	) error

	DeleteProfile(
		ctx context.Context,
		id string,
	) error

	GetCatalog(
		ctx context.Context,
		req models.GetCatalogRequest,
	) (models.GetCatalogResponse, error)
}

type serverAPI struct {
	profile1.UnimplementedProfileServer
	profile Profile
}

func Register(gRPC *grpc.Server, profile Profile) {
	profile1.RegisterProfileServer(gRPC, &serverAPI{profile: profile})
}

func (s *serverAPI) CreateProfile(ctx context.Context, req *profile1.CreateProfileRequest) (*emptypb.Empty, error) {
	if err := validateCreateProfile(req); err != nil {
		return nil, err
	}
	if err := s.profile.CreateProfile(ctx, req.GetUserId(), req.GetName(), req.GetLastName()); err != nil {
		return nil, status.Error(codes.Internal, "failed to create profile")
	}
	return &emptypb.Empty{}, nil
}

func (s *serverAPI) GetProfileByID(ctx context.Context, req *profile1.GetProfileByIDRequest) (*profile1.GetProfileByIDResponse, error) {
	var err error
	if err = validateGetProfileByID(req); err != nil {
		return nil, err
	}
	tmpResp := models.GetProfileByIDResponse{}
	if tmpResp, err = s.profile.GetProfileByID(ctx, req.GetUserId()); err != nil {
		return nil, status.Error(codes.Internal, "failed to get profile")
	}

	resp := &profile1.GetProfileByIDResponse{
		Name:        tmpResp.Name,
		LastName:    tmpResp.LastName,
		Sex:         tmpResp.Sex,
		DateOfBirth: tmpResp.DateOfBirth,
		City:        tmpResp.City,
		Height:      tmpResp.Height,
		ChestSize:   tmpResp.ChestSize,
		WaistSize:   tmpResp.WaistSize,
		HipSize:     tmpResp.HipSize,
		Weight:      tmpResp.Weight,
		Photos:      tmpResp.Photos,
		AvatarUrl:   tmpResp.AvatarURL,
	}
	return resp, nil
}

func (s *serverAPI) UpdateProfile(ctx context.Context, req *profile1.UpdateProfileRequest) (*emptypb.Empty, error) {
	if err := validateUpdateProfile(req); err != nil {
		return nil, err
	}

	data := profileRequestToModel(req)

	if err := s.profile.UpdateProfile(ctx, data); err != nil {
		return nil, status.Error(codes.Internal, "failed to update profile")
	}
	return &emptypb.Empty{}, nil
}

func (s *serverAPI) DeleteProfile(ctx context.Context, req *profile1.DeleteProfileRequest) (*emptypb.Empty, error) {
	if err := validateDeleteProfile(req); err != nil {
		return nil, err
	}
	if err := s.profile.DeleteProfile(ctx, req.GetUserId()); err != nil {
		return nil, status.Error(codes.Internal, "failed to delete profile")
	}

	return &emptypb.Empty{}, nil
}

func (s *serverAPI) GetCatalog(ctx context.Context, req *profile1.GetCatalogRequest) (*profile1.GetCatalogResponse, error) {
	if err := validateGetCatalog(req); err != nil {
		return nil, err
	}

	data := catalogRequestToModel(req)

	tmpResp, err := s.profile.GetCatalog(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get catalog")
	}

	resp := catalogResponseModelToNeedResponse(tmpResp)

	return resp, nil
}

func validateCreateProfile(req *profile1.CreateProfileRequest) error {
	if req == nil {
		return errors.New("req is nil")
	}
	if req.GetUserId() == "" {
		return errors.New("user id is empty")
	}
	if req.GetName() == "" {
		return errors.New("name is empty")
	}
	if req.GetLastName() == "" {
		return errors.New("last name is empty")
	}
	return nil
}

func validateDeleteProfile(req *profile1.DeleteProfileRequest) error {
	if req == nil {
		return errors.New("req is nil")
	}
	if req.GetUserId() == "" {
		return errors.New("user id is empty")
	}
	return nil
}

func validateUpdateProfile(req *profile1.UpdateProfileRequest) error {
	if req == nil {
		return errors.New("req is nil")
	}
	if req.GetUserId() == "" {
		return errors.New("user id is empty")
	}

	if req.GetName() == "" {
		return errors.New("name is empty")
	}
	if req.GetLastName() == "" {
		return errors.New("last name is empty")
	}

	return nil
}

func validateGetProfileByID(req *profile1.GetProfileByIDRequest) error {
	if req == nil {
		return errors.New("req is nil")
	}

	if req.GetUserId() == "" {
		return errors.New("user id is empty")
	}
	return nil
}

func validateGetCatalog(req *profile1.GetCatalogRequest) error {
	if req == nil {
		return errors.New("req is nil")
	}

	if req.Page < 0 {
		return errors.New("page must be greater than zero")
	}

	if req.Limit < 1 {
		return errors.New("limit must be greater than one")
	}
	return nil
}

func profileRequestToModel(req *profile1.UpdateProfileRequest) models.UpdateProfileRequest {
	var dateOfBirth time.Time
	if req.GetDateOfBirth() != nil {
		dateOfBirth = req.GetDateOfBirth().AsTime()
	}

	return models.UpdateProfileRequest{
		Name:        req.GetName(),
		LastName:    req.GetLastName(),
		Sex:         req.GetSex(),
		DateOfBirth: dateOfBirth,
		City:        req.GetCity(),
		Height:      req.GetHeight(),
		ChestSize:   req.GetChestSize(),
		WaistSize:   req.GetWaistSize(),
		HipSize:     req.GetHipSize(),
		Weight:      req.GetWeight(),
		Photos:      req.GetPhotos(),
		AvatarURL:   req.GetAvatarUrl(),
		UserID:      req.GetUserId(),
	}
}

func catalogRequestToModel(req *profile1.GetCatalogRequest) models.GetCatalogRequest {
	resp := models.GetCatalogRequest{}
	now := time.Now()

	if req.AgeFrom != nil {
		// Дата "от" = текущая дата - возраст (максимальная дата рождения)
		resp.DateFrom = time.Date(now.Year()-int(req.AgeFrom.GetValue()), now.Month(), now.Day(),
			0, 0, 0, 0, time.UTC)
	}

	if req.AgeTo != nil {
		// Дата "до" = текущая дата - возраст - 1 год (минимальная дата рождения)
		resp.DateTo = time.Date(now.Year()-int(req.AgeTo.GetValue())-1, now.Month(), now.Day(),
			23, 59, 59, 0, time.UTC)
	}
	if req.GetSex() != nil {
		resp.Sex = req.GetSex().GetValue()
	}
	if req.GetHeightFrom() != nil {
		resp.HeightFrom = req.GetHeightFrom().GetValue()
	}
	if req.GetHeightTo() != nil {
		resp.HeightTo = req.GetHeightTo().GetValue()
	}
	if req.GetWeightTo() != nil {
		resp.WeightTo = req.GetWeightTo().GetValue()
	}
	if req.GetWeightFrom() != nil {
		resp.WeightFrom = req.GetWeightFrom().GetValue()
	}

	resp.Limit = req.GetLimit()
	resp.Page = req.GetPage()

	return resp
}

func catalogResponseModelToNeedResponse(tmpResp models.GetCatalogResponse) *profile1.GetCatalogResponse {
	pbItems := make([]*profile1.CatalogItem, len(tmpResp.Items))
	for i, item := range tmpResp.Items {
		pbItems[i] = &profile1.CatalogItem{
			UserId:    item.UserId,
			AvatarUrl: item.AvatarUrl,
			Name:      item.Name,
			LastName:  item.LastName,
		}
	}
	return &profile1.GetCatalogResponse{
		Items:      pbItems,
		TotalPages: tmpResp.TotalPages,
	}
}
