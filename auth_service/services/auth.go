package services

import (
	"errors"
	"time"

	"auth_service/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	Register(user *models.User) error
	Login(credentials *models.AuthRequest) (string, error)
}

type AuthService struct {
	db        *gorm.DB
	jwtSecret string
}

var _ AuthServiceInterface = (*AuthService)(nil)

func NewAuthService(db *gorm.DB, jwtSecret string) AuthServiceInterface {
	return &AuthService{db: db, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.db.Create(user).Error
}

func (s *AuthService) Login(credentials *models.AuthRequest) (string, error) {
	var user models.User
	if err := s.db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.jwtSecret))
}
