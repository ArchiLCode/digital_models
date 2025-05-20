package tests

import (
	"auth/tests/suite"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v5"
	ssov1 "github.com/kristrex/protos/gen/go/sso"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const (
	emptyAppID = 0
	appID      = 1
	appSecret  = "test-secret"

	passDefaultLen = 10
)

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, passDefaultLen)
}

func TestAuthRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	// Генерация данных для регистрации
	name := gofakeit.Name()
	lastname := gofakeit.LastName()
	email := gofakeit.Email()
	goal := ""
	if gofakeit.Number(1, 10)%2 == 0 {
		goal = "model"
	} else {
		goal = "scout"
	}
	pass := randomFakePassword()

	// Регистрация пользователя
	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Name:     name,
		LastName: lastname,
		Email:    email,
		Goal:     goal,
		Password: pass,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	// Авторизация пользователя
	respLogin, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: pass,
		AppId:    appID,
	})
	require.NoError(t, err)

	// Проверка времени авторизации
	loginTime := time.Now()

	// Проверка access token
	accessToken := respLogin.GetAccessToken()
	require.NotEmpty(t, accessToken)

	accessParsed, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	require.NoError(t, err)

	claimsAccess, ok := accessParsed.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, respReg.GetUserId(), claimsAccess["uid"].(string))
	assert.Equal(t, email, claimsAccess["email"].(string))
	assert.Equal(t, appID, int(claimsAccess["app_id"].(float64)))
	const deltaSec = 1
	assert.InDelta(t, loginTime.Add(st.Cfg.TokenTTL).Unix(), claimsAccess["exp"].(float64), deltaSec)

	// Проверка refresh token
	refreshToken := respLogin.GetRefreshToken()
	require.NotEmpty(t, refreshToken)

	refreshParsed, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	require.NoError(t, err)

	claimsRefresh, ok := refreshParsed.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, respReg.GetUserId(), claimsRefresh["uid"].(string))
	assert.Equal(t, email, claimsRefresh["email"].(string))
	assert.Equal(t, appID, int(claimsRefresh["app_id"].(float64)))

	// Проверка времени истечения refresh token
	refreshTTL := st.Cfg.RefreshTTL // Предполагается, что RefreshTokenTTL определен в конфигурации
	assert.InDelta(t, loginTime.Add(refreshTTL).Unix(), claimsRefresh["exp"].(float64), deltaSec)
}

func TestRegisterLogin_DublicatedRegistration(t *testing.T) {
	ctx, st := suite.New(t)

	name := gofakeit.Name()
	lastname := gofakeit.LastName()
	email := gofakeit.Email()
	goal := ""
	if gofakeit.Number(1, 10)%2 == 0 {
		goal = "model"
	} else {
		goal = "scout"
	}
	pass := randomFakePassword()

	// Регистрация пользователя
	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Name:     name,
		LastName: lastname,
		Email:    email,
		Goal:     goal,
		Password: pass,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	respReg, err = st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Name:     name,
		LastName: lastname,
		Email:    email,
		Goal:     goal,
		Password: pass,
	})
	require.Error(t, err)
	assert.Empty(t, respReg.GetUserId())
	fmt.Println(err)
	assert.ErrorContains(t, err, "user already exists")
}
