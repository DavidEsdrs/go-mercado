package service

import (
	"fmt"
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

func (u *UserService) Signup(username, email, password string) (model.User, error) {
	var user model.User

	if user, err := u.userRepository.FindUserByEmail(email); user.ID != 0 || err != nil {
		return user, fmt.Errorf("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return user, fmt.Errorf("unable to hash password")
	}

	user.Username = username
	user.Email = email
	user.HashPassword = string(hash)

	if err := u.userRepository.Insert(&user); err != nil {
		return user, err
	}

	return user, nil
}
