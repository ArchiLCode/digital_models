package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/models"
	"gateway/utils"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"

	gc "gateway/grpc_client"
	_ "gateway/models"
	pf "github.com/kristrex/protos/gen/go/profile"
	pb "github.com/kristrex/protos/gen/go/sso"
)

type Handlers struct {
	gClient *gc.GrpcClient
}

func NewHandlers() *Handlers {
	return &Handlers{
		gClient: gc.NewGRPCClient("app:44044", "profile:44045"), // Адрес gRPC-сервиса
	}
}

// Register
// @Summary Register a new user
// @Description Creates a new user account by calling the gRPC service.
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body models.RegisterRequest true "Registration data"
// @Success 201 {object} models.RegisterResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/auth/register [post]
func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var req pb.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.gClient.AuthClient.Register(context.Background(), &req)
	if err != nil {
		log.Printf("Error calling Register: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	reqProfile := &pf.CreateProfileRequest{
		UserId:   resp.UserId,
		Name:     req.Name,
		LastName: req.LastName,
	}
	_, err = h.gClient.ProfileClient.CreateProfile(context.Background(), reqProfile)
	if err != nil {
		log.Printf("Error calling CreateProfile: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Login
// @Summary Login a user
// @Description Authenticates a user and returns an access token by calling the gRPC service.
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/auth/login [post]
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var req pb.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.gClient.AuthClient.Login(context.Background(), &req)
	if err != nil {
		log.Printf("Error calling Login: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// IsAdmin
func (h *Handlers) IsAdmin(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	req := &pb.IsAdminRequest{UserId: strconv.FormatInt(id, 10)}
	resp, err := h.gClient.AuthClient.IsAdmin(context.Background(), req)
	if err != nil {
		log.Printf("Error calling IsAdmin: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// RefreshToken
// @Summary Refresh user access token
// @Security ApiKeyAuth
// @Description Exchanges a refresh token for a new access token by calling the gRPC service.
// @Tags Auth
// @Accept json
// @Produce json
// @Param refresh_token body models.RefreshTokenRequest true "Refresh token stored in a cookie"
// @Success 200 {object} models.RefreshTokenResponse
// @Failure 401 {object} models.ErrorResponse "Invalid refresh token"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/auth/refresh-token [post]
func (h *Handlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req pb.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.gClient.AuthClient.RefreshToken(context.Background(), &req)
	if err != nil {
		strErr := fmt.Sprintf("Err: %v", err)
		log.Print("Error calling RefreshToken: ", strErr)
		http.Error(w, "Internal Server Error"+strErr, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Logout
// @Summary Logout user
// @Security ApiKeyAuth
// @Description Invalidates the user's refresh token by calling the gRPC service.
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token" default(Bearer <token>)
// @Success 204 "User logged out successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/auth/logout [get]
func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")
	token, err := utils.ExtractToken(authorizationHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req pb.LogoutRequest
	req.AccessToken = token

	if _, err := h.gClient.AuthClient.Logout(context.Background(), &req); err != nil {
		log.Printf("Error calling Logout: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Session
// @Summary Session info about use
// @Security ApiKeyAuth
// @Description Invalidates the user's refresh token by calling the gRPC service.
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token in the format 'Bearer <token>'"
// @Success 200 {object} models.SessionResponse
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/auth/session [get]
func (h *Handlers) Session(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")
	token, err := utils.ExtractToken(authorizationHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req pb.SessionRequest
	req.AccessToken = token

	resp, err := h.gClient.AuthClient.Session(context.Background(), &req)
	if err != nil {
		log.Printf("Error calling Session: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GetProfile godoc
// @Summary Get profile by ID
// @Security ApiKeyAuth
// @Description Returns the profile information for a user by their ID.
// @Tags profile
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} models.GetProfileResponse "Successful operation"
// @Failure 400 {string} string "Invalid user ID"
// @Failure 500 {string} string "Internal server error"
// @Router /api/profile [get]
func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Извлекаем user_id из query-параметров
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Создаем запрос с user_id
	req := &pf.GetProfileByIDRequest{
		UserId: userID,
	}

	// Вызываем метод gRPC клиента
	resp, err := h.gClient.ProfileClient.GetProfileByID(context.Background(), req)
	if err != nil {
		strErr := fmt.Sprintf("Err: %v", err)
		log.Print("Error calling GetProfile: ", strErr)
		http.Error(w, "Internal Server Error: "+strErr, http.StatusInternalServerError)
		return
	}

	// Возвращаем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// UpdateProfile обновляет профиль пользователя.
// @Summary Обновление профиля пользователя
// @Security ApiKeyAuth
// @Description Обновляет информацию о профиле пользователя по его ID.
// @Tags profile
// @Accept json
// @Produce json
// @Param user_id query string true "ID пользователя"
// @Param request body models.UpdateProfileRequest true "Данные для обновления профиля"
// @Success 200 "Успешное обновление профиля"
// @Failure 400 {string} string "Неверный запрос (некорректный user_id или тело запроса)"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/profile [put]
func (h *Handlers) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Извлекаем user_id из query-параметров
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var tmpReq models.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&tmpReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Print(tmpReq)
	req := &pf.UpdateProfileRequest{
		Name:        tmpReq.Name,
		LastName:    tmpReq.LastName,
		Sex:         tmpReq.Sex,
		DateOfBirth: timestamppb.New(tmpReq.DateOfBirth),
		City:        tmpReq.City,
		Height:      tmpReq.Height,
		ChestSize:   tmpReq.ChestSize,
		WaistSize:   tmpReq.WaistSize,
		HipSize:     tmpReq.HipSize,
		Weight:      tmpReq.Weight,
		Photos:      tmpReq.Photos,
		AvatarUrl:   tmpReq.AvatarUrl,
		UserId:      userID,
	}
	// Вызываем метод gRPC клиента
	resp, err := h.gClient.ProfileClient.UpdateProfile(context.Background(), req)
	if err != nil {
		strErr := fmt.Sprintf("Err: %v", err)
		log.Print("Error calling UpdateProfile: ", strErr)
		http.Error(w, "Internal Server Error: "+strErr, http.StatusInternalServerError)
		return
	}

	// Возвращаем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DeleteProfile удаляет профиль пользователя.
// @Summary Удаление профиля пользователя
// @Security ApiKeyAuth
// @Description Удаляет профиль пользователя по его ID.
// @Tags profile
// @Accept json
// @Produce json
// @Param user_id query string true "ID пользователя"
// @Success 200 "Успешное удаление профиля"
// @Failure 400 {string} string "Неверный запрос (некорректный user_id)"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/profile [delete]
func (h *Handlers) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	// Извлекаем user_id из query-параметров
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Создаем запрос с user_id
	req := &pf.DeleteProfileRequest{
		UserId: userID,
	}

	// Вызываем метод gRPC клиента
	resp, err := h.gClient.ProfileClient.DeleteProfile(context.Background(), req)
	if err != nil {
		strErr := fmt.Sprintf("Err: %v", err)
		log.Print("Error calling DeleteProfile: ", strErr)
		http.Error(w, "Internal Server Error: "+strErr, http.StatusInternalServerError)
		return
	}

	// Возвращаем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handlers) GetCatalog(w http.ResponseWriter, r *http.Request) {
	params, err := ParseCatalogQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	grpcReq := ToGRPCRequest(params)

	// Вызов gRPC сервиса
	response, err := h.gClient.GetCatalog(r.Context(), grpcReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа клиенту
	json.NewEncoder(w).Encode(response)
}

func ToGRPCRequest(p *models.CatalogQueryParams) *pf.GetCatalogRequest {
	req := &pf.GetCatalogRequest{
		Page:  p.Page,
		Limit: p.Limit,
	}

	if p.AgeFrom != nil {
		req.AgeFrom = wrapperspb.Int32(*p.AgeFrom)
	}
	if p.AgeTo != nil {
		req.AgeTo = wrapperspb.Int32(*p.AgeTo)
	}
	if p.Sex != nil {
		req.Sex = wrapperspb.String(*p.Sex)
	}
	if p.HeightFrom != nil {
		req.HeightFrom = wrapperspb.Int32(*p.HeightFrom)
	}
	if p.HeightTo != nil {
		req.HeightTo = wrapperspb.Int32(*p.HeightTo)
	}
	if p.WeightFrom != nil {
		req.WeightFrom = wrapperspb.Int32(*p.WeightFrom)
	}
	if p.WeightTo != nil {
		req.WeightTo = wrapperspb.Int32(*p.WeightTo)
	}

	return req
}

func ParseCatalogQueryParams(r *http.Request) (*models.CatalogQueryParams, error) {
	var params models.CatalogQueryParams
	queryDecoder := schema.NewDecoder()

	// Парсим query-параметры
	if err := queryDecoder.Decode(&params, r.URL.Query()); err != nil {
		return nil, fmt.Errorf("failed to decode query params: %w", err)
	}

	// Валидируем параметры
	if err := validator.New().Struct(params); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return &params, nil
}
