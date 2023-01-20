package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var sampleSecretKey = []byte("JWTSecret")


func GenerateJWT(id uint,email string)(string,error){
	
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = email
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(jwtToken string) ( jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(jwtToken,claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return sampleSecretKey, nil
	})

	if token == nil {
		return claims,fmt.Errorf("there was an error in parsing")
	}
	
	
	return claims,nil
}