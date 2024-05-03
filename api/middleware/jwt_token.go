package middleware

import (
	"errors"
	"prototype/constant"
	"prototype/domain"
	jwt_con "prototype/domain"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func CreateTokenJWT(userId uuid.UUID, email string, role string) (string, error) {

	var userClaims = jwt_con.JwtCustomClaims{
		UserId: userId, UserEmail: email, Roles: role, RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	return token.SignedString([]byte(constant.PrivateKeyJWT()))
}

// Extrak jwt
func ExtractToken(token string) (uuid.UUID, string, string, error) {

	tokenParts := strings.Split(token, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "bearer" {
		return uuid.Nil, "", "", errors.New("Invalid token format")
	}

	parseToken, err := jwt.ParseWithClaims(tokenParts[1], &domain.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.PrivateKeyJWT()), nil
	})
	if err != nil {
		return uuid.Nil, "", "", err
	}

	if claims, ok := parseToken.Claims.(*domain.JwtCustomClaims); ok && parseToken.Valid {
		return claims.UserId, claims.UserEmail, claims.Roles, nil
	}

	return uuid.Nil, "", "", errors.New("Invalid token claims")
}
