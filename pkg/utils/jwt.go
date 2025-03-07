package utils

import (
	"errors"
	"primeskills-test-api/internal/domain/entity"
	"primeskills-test-api/pkg/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a JWT token for a given user.
// It takes a pointer to an entity.User and returns a pointer to the generated token string and an error if any.
// The token contains the user's ID, name, email, and other registered claims.
func GenerateToken(user *entity.User) (*string, error) {
	cfg := config.Cfg

	secretKey := []byte(cfg.Jwt.SecretKey)

	claims := entity.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			Issuer:    cfg.Jwt.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.Jwt.ExpiresIn)),
		},
		Name:  user.Name,
		Email: user.Email,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// VerifyToken verifies a JWT token string and returns the claims if the token is valid.
// It takes a token string as input and returns a pointer to the JwtTokenClaims and an error if any.
// The function checks the token's validity, parses the claims, and verifies the issuer.
func VerifyToken(tokenString string) (*entity.JwtTokenClaims, error) {
	cfg := config.Cfg

	secretKey := []byte(cfg.Jwt.SecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &entity.JwtTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
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

// ExtractClaims extracts the JWT claims from the Gin context.
// It takes a Gin context as input and returns a pointer to the JwtTokenClaims if they exist in the context.
// If the claims do not exist or are of the wrong type, it returns nil.
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
