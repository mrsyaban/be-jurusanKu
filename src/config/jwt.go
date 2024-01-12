package config

import (
	"time"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Payload struct {
	ID 		 uuid.UUID `json:"id"`
	Email	 string    `json:"email"`
	Role     string    `json:"role"`
	IssuedAt time.Time `json:"issued_at"`
	Exp      time.Time `json:"expired_at"`
}

func CreateToken(email string, role string, duration time.Duration, key string) (string, *Payload, error) {
	payload := &Payload{
		Email: email,
		Role:     role,
		IssuedAt: time.Now(),
		Exp:      time.Now().Add(duration),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(key))
	return token, payload, err
}

func (p *Payload) Valid() error {
	if time.Now().After(p.Exp) {
		return errors.New("token has expired")
	}
	return nil
}

func VerifyToken(token string, key string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token is invalid")
		}
		return []byte(key), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New("token has expired")) {
			return nil, errors.New("token has expired")
		}
		return nil, errors.New("token is invalid")
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	return payload, nil
}

func GetUserEmailByToken(token string, key string) (string, error) {
	payload, err := VerifyToken(token, key)
	if err != nil {
		return "", err
	}
	return payload.Email, nil
}