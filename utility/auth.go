package utility

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateJWT(signedby string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour)
	claims["authorized"] = true
	claims["signedby"] = signedby
	godotenv.Load()
	fmt.Println([]byte(os.Getenv("ENC_KEY")))
	sampleSecretKey := []byte(os.Getenv("ENC_KEY"))
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "Signing Error", err
	}
	return tokenString, err
}
func VerifyJWT(w http.ResponseWriter, r *http.Request, signedby string) bool {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tknStr := c.Value
	token, err := jwt.Parse(tknStr, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			Respond(http.StatusUnauthorized, "Error parsing the token", &w, false)
		}
		return "", nil

	})
	if err != nil {
		Respond(http.StatusUnauthorized, "Error parsing the token", &w, false)
		return false
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		if claims["signedby"] == "user" && signedby == "user" {
			return true
		}
		if claims["signedby"] == "admin" && signedby == "admin" {
			return true
		} else {
			Respond(http.StatusUnauthorized, "Invalid Access", &w, false)
			return false
		}
	} else {
		Respond(http.StatusUnauthorized, "Token not found in cookies", &w, false)
	}
	return false
}
