package auth

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Auth struct {
	expirationDuration time.Duration
	secretKey          []byte
}
type AuthConf struct {
	ExpirationDuration time.Duration
	SecretKey          []byte
}

func NewAuthService(conf *AuthConf) *Auth {
	return &Auth{
		expirationDuration: conf.ExpirationDuration,
		secretKey:          conf.SecretKey,
	}
}

type JWTClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id,omitempty"`
}

func (a *Auth) CreateToken(userId string) (string, error) {
	//get token expiry
	expiry := time.Now().Add(a.expirationDuration)
	claims := &JWTClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "tanvirs",
			ExpiresAt: expiry.Unix(),
		},
		userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secretKey)
}

func (a *Auth) ValidateToken(tokenStr string) (*JWTClaims, error) {
	// claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})
	if err != nil {
		log.Println("error while token authorization", err)
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		log.Println("error while token authorization, invalid token", err)
		return nil, err
	}
	return claims, nil
}
