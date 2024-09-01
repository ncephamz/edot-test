package jwt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	modelUser "user-service/models/user"
	cErr "user-service/pkg/err"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	Secret string
}

func (j Jwt) CreateToken(claims modelUser.Claims) (modelUser.Token, error) {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims.Id,
		"exp":     claims.ExpiredIn,
	})

	result := modelUser.Token{}

	t, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return result, err
	}

	result.AccessToken = t

	t, err = j.createRefreshToken(result.AccessToken)
	if err != nil {
		return result, err
	}

	result.RefreshToken = t

	return result, nil
}

func (j Jwt) ValidateToken(accessToken string) (modelUser.Claims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.Secret), nil
	})

	claims := modelUser.Claims{}
	if err != nil {
		return claims, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		claims.Id = payload["user_id"].(string)
		claims.ExpiredIn = payload["exp"].(int64)

		return claims, nil
	}

	return claims, cErr.InvalidToken
}

func (j Jwt) createRefreshToken(accesToken string) (string, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, j.Secret)

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		fmt.Println(err.Error())

		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	refreshToken := base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(accesToken), nil))

	return refreshToken, nil
}

func (j Jwt) ValidateRefreshToken(model modelUser.Token) (modelUser.Claims, error) {
	var stringToken = strings.Split(model.AccessToken, " ")
	model.AccessToken = stringToken[1]

	sha1 := sha1.New()
	io.WriteString(sha1, j.Secret)

	user := modelUser.Claims{}
	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return user, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return user, err
	}

	data, err := base64.URLEncoding.DecodeString(model.RefreshToken)
	if err != nil {
		return user, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return user, err
	}

	if string(plain) != model.AccessToken {
		return user, cErr.InvalidToken
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(model.AccessToken, claims)

	if err != nil {
		return user, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return user, cErr.InvalidToken
	}

	user.Id = payload["user_id"].(string)

	return user, nil
}
