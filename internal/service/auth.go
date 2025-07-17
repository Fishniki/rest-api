package service

import (
	"context"
	"errors"
	"rest-api/domain"
	"rest-api/dto"
	"rest-api/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}


func NewAuth(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return AuthService{
		conf:           cnf,
		userRepository: userRepository,
	}
}


	// Login implements domain.AuthService.
func (a AuthService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {

	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return  dto.AuthResponse{}, err
	}

	if user.Id == "" {
		return  dto.AuthResponse{}, errors.New("autentikasi gagal1")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		 return dto.AuthResponse{}, errors.New("autentikasi gagal2")
	}

	claim := jwt.MapClaims{
		"id": user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return  dto.AuthResponse{}, errors.New("autentikasi gagal3")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil

}
