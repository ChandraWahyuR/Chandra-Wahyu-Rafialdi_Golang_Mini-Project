package middleware

import (
	"prototype/constant"
	jwt_con "prototype/domain"
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
