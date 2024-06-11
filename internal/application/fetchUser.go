package application

import (
	"github.com/silvano-paulino/internal/domain"
	"github.com/silvano-paulino/internal/repository"
)

// Função que utiliza o Database para buscar um usuário por ID.
func FetchUser(db repository.Repository, id string) (*domain.User, error) {
	return db.GetUserById(id)
}