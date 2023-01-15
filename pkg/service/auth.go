package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Akmyrza/todo-app"
	"github.com/Akmyrza/todo-app/pkg/repository"
)

const salt = "hashdgahsgd1h2gh13gh1j23"

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}