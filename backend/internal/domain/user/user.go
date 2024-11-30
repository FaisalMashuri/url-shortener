package user

import (
	"context"
	"database/sql"
	"backend/internal/domain/user/dto/requests"
	"backend/internal/domain/user/dto/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
	Email      string
	Password   string
	Phone      string       `gorm:"uniqueIndex"`
	VerifiedAt sql.NullTime `gorm:"default:null"`
}

func NewUser(request requests.RegisterRequest) User {
	return User{
		ID:       uuid.NewString(),
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
	}
}

type UserRepository interface {
	FindUserByEmail(db *gorm.DB, email string) (*User, error)
	CreateUser(db *gorm.DB, user *User) (*User, error)
	UpdateUser(db *gorm.DB, updatedUser *User) (*User, error)
	GetUserByID(db *gorm.DB, id string) (*User, error)
	GetAllUser(db *gorm.DB) ([]*User, error)
	DeleteUser(db *gorm.DB, id string) error
	VerifyUser(db *gorm.DB, id string) error
}

type UserService interface {
	Login(ctx context.Context, userRequest *requests.LoginRequest) (res *responses.LoginResponse, err error)
	RegisterUser(ctx context.Context, userRequest *requests.RegisterRequest) (user *User, err error)
	GetDetailUserById(ctx context.Context, id string) (user responses.UserDetail, err error)
	VerifyUser(ctx context.Context, verReq requests.VerifiedUserRequest) error
	ResendOTP(ctx context.Context, userId string) error
}

type UserController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	GetDetailUserJWT(ctx *fiber.Ctx) error
	VerifyUser(ctx *fiber.Ctx) error
	ResendOTP(ctx *fiber.Ctx) error
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
