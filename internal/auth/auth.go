package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	microauth "github.com/micro/go-micro/v2/auth"
)

// ErrNotJWT is returned by AccountFromToken when the provided token is
// not a JWT. In such a case, the caller should use another method to
// inspect the token.
var ErrNotJWT = errors.New("auth: not a jwt token")

// authClaims to be encoded in the JWT
type authClaims struct {
	Type     string            `json:"type"`
	Scopes   []string          `json:"scopes"`
	Metadata map[string]string `json:"metadata"`

	jwt.StandardClaims
}

// AccountToToken converts account to JWT token
func AccountToToken(acc *microauth.Account) string {
	claims := authClaims{
		Type:     acc.Type,
		Scopes:   acc.Scopes,
		Metadata: acc.Metadata,
		StandardClaims: jwt.StandardClaims{
			Subject: acc.ID,
			Issuer:  acc.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte("noop"))
	return ss
}

// AccountFromToken restore account from the access JWT token
func AccountFromToken(token string) (*microauth.Account, error) {
	// check token format
	if len(strings.Split(token, ".")) != 3 {
		return nil, ErrNotJWT
	}

	// get the public key from env
	key := os.Getenv("MICRO_AUTH_PUBLIC_KEY")
	if key == "" {
		//logger.Info("env MICRO_AUTH_PUBLIC_KEY is not set")
		// try to decode JWT without verify signature
		res, _, err := new(jwt.Parser).ParseUnverified(token, &authClaims{})
		if err != nil {
			return nil, fmt.Errorf("can not parse jwt: %v", err)
		}
		claims, ok := res.Claims.(*authClaims)
		if !ok {
			return nil, fmt.Errorf("jwt claims type is incorrect")
		}
		return &microauth.Account{
			ID:       claims.Subject,
			Issuer:   claims.Issuer,
			Type:     claims.Type,
			Scopes:   claims.Scopes,
			Metadata: claims.Metadata,
		}, nil
	}

	// decode the public key
	pub, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("env MICRO_AUTH_PUBLIC_KEY is incorrect: %v", err)
	}

	// parse the public key
	res, err := jwt.ParseWithClaims(token, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(pub)
	})
	if err != nil {
		return nil, fmt.Errorf("parse jwt: %v", err)
	}

	// validate the token
	if !res.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := res.Claims.(*authClaims)
	if !ok {
		return nil, fmt.Errorf("can not type assert to authClaims")
	}

	// return the token
	return &microauth.Account{
		ID:       claims.Subject,
		Issuer:   claims.Issuer,
		Type:     claims.Type,
		Scopes:   claims.Scopes,
		Metadata: claims.Metadata,
	}, nil
}
