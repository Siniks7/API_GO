package jwt_test

import (
	"api/configs"
	"api/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "siniks7@yandex.ru"
	conf := configs.LoadConfig()
	jwtService := jwt.NewJWT(conf.Auth.Secret)
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
