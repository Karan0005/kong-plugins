package main

import (
	"io/ioutil"
	"strings"

	"github.com/Kong/go-pdk"
	"github.com/dgrijalva/jwt-go"
)

type Config struct{}

func New() interface{} {
	return &Config{}
}

// Custom Plugin Logic Space
func (conf Config) Access(kong *pdk.PDK) {
	var keyBase []byte
	var readErr error
	resHeader := make(map[string][]string)
	resHeader["Content-Type"] = append(resHeader["Content-Type"], "application/json")

	//If Authorization Token Not Provided
	clientToken, err := kong.Request.GetHeader("authorization")
	if clientToken == "" || err != nil {
		kong.Response.Exit(401, "Authorization Token is Required", resHeader)
		return
	}

	//If Incorrect Authorization Token Not Provided
	parsedToken := strings.Split(clientToken, "Bearer ")
	if len(parsedToken) != 2 {
		kong.Response.Exit(401, "Invalid Authorization Token Format", resHeader)
		return
	}

	//If Unable to Read Public Key File of Auth0
	keyBase, readErr = ioutil.ReadFile("/tmp/public.key")
	if readErr != nil {
		kong.Response.Exit(500, "Unable to Read Public Key File", resHeader)
		return
	}

	//If Unable to Parse Public Key File of Auth0
	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(keyBase)
	if err != nil {
		kong.Response.Exit(500, "Unable to Parse Public Key File", resHeader)
		return
	}

	//If Unable to Parse Provided Authorization Token with Public Key File of Auth0
	token, err := jwt.Parse(parsedToken[1], func(token *jwt.Token) (interface{}, error) { return parsedKey, nil })
	if !token.Valid && err != nil {
		ve := err.(*jwt.ValidationError)
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			kong.Response.Exit(401, "Token is Malformed", resHeader)
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			kong.Response.Exit(401, "Token is Expired", resHeader)
		} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			kong.Response.Exit(401, "Signature is Invalid", resHeader)
		} else {
			kong.Response.Exit(401, "Invalid Authorization Token", resHeader)
		}
		return
	}
}
