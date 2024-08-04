package service

import (
	"os"
	"time"

	"github.com/DavidEsdrs/go-mercado/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Repository[model.User]
	FindUserByEmail(email string) (model.User, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (u *UserService) Login(email, password string) (string, error) {
	user, err := u.userRepository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 2).Unix(), // 2 days
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
