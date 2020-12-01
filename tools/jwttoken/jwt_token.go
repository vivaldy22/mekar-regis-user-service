package jwttoken

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	formatDateLayout = "2006-01-02 15:04:05"
)

func JwtEncoder(userName, customKey, hmacSampleSecret string) (string, error) {
	claims := jwt.MapClaims{
		"name":      userName,
		"customKey": customKey,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

/*
func JwtDecoder(tokenString, hmacSampleSecret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSampleSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["name"], claims["customKey"], claims["expiredAt"])
		expiredAt := claims["expiredAt"].(string)

		t, err := time.Parse(formatDateLayout, expiredAt)
		if err != nil {
			return nil, err
		}
		thisTimeString := time.Now().Format(formatDateLayout)
		thisTime, _ := time.Parse(formatDateLayout, thisTimeString)
		diffTime := t.Sub(thisTime).Seconds()

		if diffTime > 0 {
			return claims, nil
		} else {
			return nil, errors.New("expired token")
		}

	} else {
		fmt.Println(err)
		return nil, err
	}
}
*/
