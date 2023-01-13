package usecases

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/Brigant/GoPetPorject/app/enteties"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	signingKey      = "sdFWlnxb13t&r"
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 24 * time.Hour
)

type Auth struct {
	Repo Repository
}

type tokenClaims struct {
	jwt.StandardClaims
	Id string
}

func NewAuth(repo Repository) Auth {
	return Auth{Repo: repo}
}

// The function CreateUser represents bissness logic layer
// and  execute some manipulation with data before
// saving  this data in the repository.
func (a Auth) CreateUser(user enteties.User) (string, error) {
	user.Password = a.sha256(user.Password)

	id, err := a.Repo.AddUser(user)
	if err != nil {
		return "", fmt.Errorf("error occures while AddUser: %w", err)
	}

	return id, nil
}

// The function sha256 returns the checksum of the string using SHA256 hash algorithms.
func (a Auth) sha256(someString string) string {
	h := sha256.Sum256([]byte(someString))

	return fmt.Sprintf("%x", h)
}

// The function GenerateToken represents bissness logic layer
// and  generate token.
func (a Auth) GenerateToken(phone int, password string) (accessToken, refreshToken string, err error) {
	userId, err := a.Repo.GetUser(phone, a.sha256(password))
	if err != nil {
		return "", "", fmt.Errorf("error occures while GetUser: %w", err)
	}

	accessToken, err = a.generateAccessToken(userId)
	if err != nil {
		return "", "", fmt.Errorf("cerror occures while generateAccessToken: %w", err)
	}

	refreshToken, err = a.generateRefreshToken(userId)
	if err != nil {
		return "", "", fmt.Errorf("error occures while generateRefreshToken: %w", err)
	}

	return accessToken, refreshToken, nil
}

// This function returns two tokens if refresh token is valid.
func (a Auth) RefreshTokens(token string) (accessToken, refreshToken string, err error) {
	if _, err := uuid.Parse(token); err != nil {
		return "", "", err
	}

	session, err := a.Repo.GetSession(token)
	if err != nil {
		return "", "", fmt.Errorf("cannot GetSesion with such refresh token: %w", err)
	}

	if session.Created.Add(session.ExpiresIn).Unix() < time.Now().Unix() {
		a.Repo.DeleteSession(session.ID)
		if err != nil {
			return "", "", fmt.Errorf("cannot DeleteSession: %w", err)
		}

		return "", "", enteties.ErrRefreshTokenExpired
	}

	accessToken, err = a.generateAccessToken(session.UserID)
	if err != nil {
		return "", "", fmt.Errorf("cannot generateAccessToken: %w", err)
	}

	updatedTTL := time.Duration(time.Now().Unix()-session.Created.Unix())*time.Second + refreshTokenTTL

	if err := a.Repo.UpdateSession(token, time.Duration(updatedTTL)); err != nil {
		return "", "", fmt.Errorf("an error occurred during UpdateSession: %w", err)
	}

	return accessToken, token, nil
}

// The function deletes the given token.
func (a Auth) DeleteToken(token string) error {
	if _, err := uuid.Parse(token); err != nil {
		return err
	}

	if err := a.Repo.DeleteSession(token); err != nil {
		return err
	}

	return nil
}

// The function returns user ID if accessToken is valid.
func (a Auth) ParseToken(accesToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing metod")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Id, nil
}

func (a Auth) generateAccessToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	accessToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("cannot get SignetString token: %w", err)
	}

	return accessToken, nil
}

func (a Auth) generateRefreshToken(userID string) (string, error) {
	token, err := a.Repo.AddSession(userID, refreshTokenTTL)
	if err != nil {
		return "", fmt.Errorf("cannot generate refresh tokes cause of: %w", err)
	}

	return token, nil
}
