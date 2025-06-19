package auth_test

import (
	"api/configs"
	"api/internal/auth"
	"api/internal/user"
	"api/pkg/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLoginSuccess(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Failed init mock db")
		return
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		t.Fatal("Failed init gorm")
		return
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
}
