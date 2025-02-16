package utilities

import (
	"errors"
	"primeskills-test-api/internal/config"
	"primeskills-test-api/internal/entity"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *entity.User) (*string, error) {
	cfg := config.Cfg

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(cfg.Jwt.PrivateKey))
	if err != nil {
		return nil, errors.New("failed to parse private key: " + err.Error())
	}

	claims := entity.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			Issuer:    cfg.Jwt.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Jwt.ExpiresIn))),
		},
		Name:  user.Name,
		Email: user.Email,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func VerifyToken(tokenString string) (*entity.JwtTokenClaims, error) {
	cfg := config.Cfg

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cfg.Jwt.PublicKey))
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &entity.JwtTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entity.JwtTokenClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	if claims.Issuer != cfg.Jwt.Issuer {
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}

func ExtractClaims(ctx *gin.Context) *entity.JwtTokenClaims {
	claims, exist := ctx.Get("claims")
	if !exist {
		return nil
	}

	c, ok := claims.(*entity.JwtTokenClaims)
	if !ok {
		return nil
	}

	return c
}
