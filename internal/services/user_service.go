package services

import (
	"CRUD_Go_gin/internal/database"
	"CRUD_Go_gin/internal/domain"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

const salt = "skjfh_(_*gldfs354fsl(&j54347"

var Users *database.Users

func SingUp(u *domain.User) error {
	var err error
	u.PasswordEnc, err = Hash(u.PasswordEnc)
	if err != nil {
		return err
	}
	err = Users.CreateUser(u)
	if err != nil {
		return err
	}
	return nil
}

func SingIn(email string, password string) (string, error) {
	passwordEnc, err := Hash(password)
	if err != nil {
		return "", err
	}
	user, err := Users.GetUser(email, passwordEnc)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Id:        strconv.Itoa(int(user.Id)),
		IssuedAt:  time.Now().Unix(),
	})

	return token.SignedString([]byte("oprgueiohgnv8sluhg7ilhglsuy574sth87s57gh73gh8s"))
}

func Hash(password string) (string, error) {
	hash := sha1.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), nil
}

func ParseToken(token string) error {
	var secret = []byte("oprgueiohgnv8sluhg7ilhglsuy574sth87s57gh73gh8s")
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return err
	}

	if !t.Valid {
		return errors.New("invalid token")
	}

	_, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid claims")
	}

	return nil
}
