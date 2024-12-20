package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("supersecret") // needs to be converted to byte for the key of es256 https://golang-jwt.github.io/jwt/usage/signing_methods/#signing-methods-and-key-types
func GenerateToken(email string, userId int64) (string, error) {
	// first argument =>
	// identifier of the signin approach that should be used. it will be checked if the token was generated by this server etc

	// HS256
	// one shared secret key is used for both signing and verifying. It is simpler but might be less secure if that one key is exposed
	//
	// ES256
	// it is a signing method for JWTs that uses two keys: one to sign (create) the token and another to verify it.
	// Why two keys? => it is called asymmetric encryption. it is more secure of certain application
	// because the signing key (private key) can be kept secret while verifying key (public key) can be shared openly

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (int64, error) {
	// when token is created in GenerateToken, the token's signature is created using `secretKey`.
	// To verify if a token is authentic, jwt.Parse needs the same key to recreate the signature and compare it with the one on the token.

	// token signature:
	// cryptographic hash that ensures token's authenticity and integrity.
	// It is like a digital seal that allows the receiver to verify that the token hasn't been tampered with and it was indeed created by the expected source

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// verifying signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // it is type checking syntax if the value of token.Method field has a type of *jwt.SigningMethodHMAC

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) // checking if it is a same method we used to create a Map (in line 23)

	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// // we can extract second parameter (ok bool) to verify but we already check if the token is valid above so we are omitting it
	// email := claims["email"].(string)

	// userId is stored as int64 in GenerateToken,
	// but when the userId is encoded as part of the JWT claims in jwt.MapClaims, it's serialized as JSON.
	// In JSON, all numbers are represented as float64 by default, so we have to convert to int64 manually
	userId := int64(claims["userId"].(float64))

	return userId, nil

}
