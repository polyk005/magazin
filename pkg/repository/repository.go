package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/polyk005/magazin/pkg/model"
	"github.com/polyk005/magazin/pkg/repository/postgres"
)	

type Authorization interface {
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}