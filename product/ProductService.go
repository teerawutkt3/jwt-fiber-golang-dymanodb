package auth

import (
	"fiber-poc-api/database/repository"
)

type ProductService struct {
	repo repository.UserRepository
}

func NewProductService(repo repository.UserRepository) ProductService {
	return ProductService{repo}
}

