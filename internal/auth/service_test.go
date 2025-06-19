package auth_test

import (
	"api/internal/auth"
	"api/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "siniks7@yandex.ru",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "siniks7@yandex.ru"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "siniks7")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("Email %s do not math %s", email, initialEmail)
	}
}
