package auth

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository     *repositories.UserRepository
	jwtkey             []byte
}

func NewAutenticacaoService(userRepository *repositories.UserRepository, jwtKey []byte) *AuthService {
	return &AuthService{
		userRepository:     userRepository,
		jwtkey:             jwtKey,
	}
}

func (s *AuthService) GerarToken(user *models.User) (string, error) {
	var tokenAss string
	var err error
	var claim *Claim
	expirationTime := time.Now().Add(20 * time.Hour)

	claim = &Claim{
		ID:     user.ID,
		Nome:   user.Nome,
		CPF:    user.CPF,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "sbs-stock-api",
			Subject:   fmt.Sprintf("%d", user.ID),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	if tokenAss, err = token.SignedString(s.jwtkey); err != nil {
		return "", err
	}

	return tokenAss, nil
}

func (s *AuthService) AuthToken(tokenStr string) (*Claim, error) {
	claim := &Claim{}

	token, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return s.jwtkey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	return claim, nil
}

func (s *AuthService) UserAuth(c *gin.Context, credentials dto.UserCredentials) (string, error) {
	var err error
	var token string

	FCredentials := strings.ReplaceAll(credentials.CPF, ".", "")
	FCredentials = strings.ReplaceAll(FCredentials, "-", "")
	FCredentials = strings.ReplaceAll(FCredentials, " ", "")

	if !utils.IsValidCPF(FCredentials){
		return "", fmt.Errorf("CPF %s Inválido", FCredentials)
	}

	user, err := s.userRepository.FindUserByCPF(c, FCredentials)
	if err != nil {
		return "", fmt.Errorf("erro ao buscar pelo usuário %s", FCredentials)
	}

	var senha = user.SenhaHash

	if err = bcrypt.CompareHashAndPassword([]byte(senha), []byte(credentials.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", fmt.Errorf("credencial inválida: %v", err.Error())
		}
		err = fmt.Errorf("erro ao validar credenciais: %v", err.Error())
		return "", err
	}

	if token, err = s.GerarToken(user); err != nil {
		err = fmt.Errorf("erro ao gerar token: %v", err.Error())
		return "", err
	}

	return token, nil
}
