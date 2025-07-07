package domain

import (
	"context"
	"rest-api/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}