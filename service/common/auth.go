package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

// private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// read the key file
func initKeys() {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}
}

// GenerateJWT token
func GenerateJWT(name, role string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set claims for JWT token
	claims := t.Claims.(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// set the expire time for the JWT token
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateToken(token *jwt.Token) (interface{}, error) {
	// Verify the token with the public key
	return verifyKey, nil
}

// Authorize middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// validate the token
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, validateToken)

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return
			default:
				DisplayAppError(w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}

		default:
			DisplayAppError(w,
				err,
				"Error while parsing Access Token!",
				500)
			return
		}
	}
	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}
