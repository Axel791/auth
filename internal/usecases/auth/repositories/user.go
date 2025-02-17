package repositories

import "github.com/jmoiron/sqlx"

// SqlUserRepository - структура репозитория
type SqlUserRepository struct {
	db *sqlx.DB
}

// NewUserRepository - конструктор репозитория пользователя
func NewUserRepository(db *sqlx.DB) *SqlUserRepository {
	return &SqlUserRepository{db: db}
}

// CreateUser - Создание пользователя
func (r *SqlUserRepository) CreateUser() {

}

// GetUserById - получение пользователя по ID
func (r *SqlUserRepository) GetUserById() {

}
