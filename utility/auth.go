package utility

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("my_secret_key")

type customClaims struct {
	Id   int
	Role string
	jwt.StandardClaims
}

func GenerateToken() (string, error) {
	claims := &customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
func AuthToken(w http.ResponseWriter, r *http.Request) (status bool) {
	c, err := r.Cookie("token")
	if err != nil {
		Respond(http.StatusBadRequest, "Not authorized to perform this request", &w, false)
		return false
	}
	token, err := jwt.ParseWithClaims(c.Value, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return token.Valid
}
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
