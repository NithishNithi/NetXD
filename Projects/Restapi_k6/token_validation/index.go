package tokenvalidation

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)
const secretKey = "your-secret-key"

var jwtKey = []byte("your-secret-key")

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("1")
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println("2")
		return nil, err
	}
	if !token.Valid {
		fmt.Println("3")
		return nil, fmt.Errorf("unvalid token")
	}
	fmt.Println("11")
	return token, nil
}

func GenerateToken()string {

	// Create a token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the "cusid" claim
	claims := token.Claims.(jwt.MapClaims)
	claims["cusid"] = 123                            // Replace with your desired cusid
	claims["exp"] = time.Now().Add(time.Hour).Unix() // Token expiration time (e.g., 1 hour)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error creating token:", err)
		return ""
	}
	return tokenString
}

// func main(){
// 	tokenStr:=GenerateToken();
// 	fmt.Println(tokenStr)
// }
