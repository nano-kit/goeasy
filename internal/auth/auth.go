package auth

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	microauth "github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/logger"
)

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
func AccountFromToken(token string) (*microauth.Account, bool) {
	// check token format
	if len(strings.Split(token, ".")) != 3 {
		logger.Infof("not a jwt token: %v", token)
		return nil, false
	}

	// get the public key from env
	key := os.Getenv("MICRO_AUTH_PUBLIC_KEY")
	if key == "" {
		//logger.Info("env MICRO_AUTH_PUBLIC_KEY is not set")
		// try to decode JWT without verify signature
		res, _, err := new(jwt.Parser).ParseUnverified(token, &authClaims{})
		if err != nil {
			logger.Infof("can not parse jwt: %v", err)
			return nil, false
		}
		claims, ok := res.Claims.(*authClaims)
		if !ok {
			logger.Infof("jwt claims type is incorrect")
			return nil, false
		}
		return &microauth.Account{
			ID:       claims.Subject,
			Issuer:   claims.Issuer,
			Type:     claims.Type,
			Scopes:   claims.Scopes,
			Metadata: claims.Metadata,
		}, true
	}

	// decode the public key
	pub, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		logger.Infof("env MICRO_AUTH_PUBLIC_KEY is incorrect: %v", err)
		return nil, false
	}

	// parse the public key
	res, err := jwt.ParseWithClaims(token, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(pub)
	})
	if err != nil {
		logger.Infof("parse jwt: %v", err)
		return nil, false
	}

	// validate the token
	if !res.Valid {
		logger.Info("invalid token")
		return nil, false
	}
	claims, ok := res.Claims.(*authClaims)
	if !ok {
		logger.Info("can not type assert to authClaims")
		return nil, false
	}

	// return the token
	return &microauth.Account{
		ID:       claims.Subject,
		Issuer:   claims.Issuer,
		Type:     claims.Type,
		Scopes:   claims.Scopes,
		Metadata: claims.Metadata,
	}, true
}
