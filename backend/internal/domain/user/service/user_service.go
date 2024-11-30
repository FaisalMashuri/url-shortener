package service

import (
	"context"

	redis_client "backend/infrastructure/redis"
	"backend/internal/domain/user"
	"backend/internal/domain/user/dto/requests"
	"backend/internal/domain/user/dto/responses"
	"gorm.io/gorm"
)

type userService struct {
	db        *gorm.DB
	repo      user.UserRepository
	redisRepo *redis_client.Repository
}

func NewService(db *gorm.DB, repo user.UserRepository, redis *redis_client.Repository) user.UserService {
	return &userService{
		db:        db,
		repo:      repo,
		redisRepo: redis,
	}
}

func (s *userService) Login(ctx context.Context, userRequest *requests.LoginRequest) (*responses.LoginResponse, error) {
	//TODO: Implement login
	return &responses.LoginResponse{}, nil
}

func (s *userService) RegisterUser(ctx context.Context, userRequest *requests.RegisterRequest) (userData *user.User, err error) {
	tx := s.db.WithContext(ctx).Begin()
	//TODO: Implement register
	userData, err = s.repo.CreateUser(tx, userData)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx = s.db.Commit()
	return userData, nil

}

func (s *userService) GetDetailUserById(ctx context.Context, id string) (res responses.UserDetail, err error) {
	//TODO: Implement getDetailUserbyId
	return res, nil
}

func (s *userService) VerifyUser(ctx context.Context, verReq requests.VerifiedUserRequest) error {
	// TODO: Implement verify user
	return nil
}

func (s *userService) ResendOTP(ctx context.Context, userId string) error {
	return nil
}
