package service

import "github.com/doug-martin/goqu/v9"

type UserRepository struct {
	db *goqu.Database
}