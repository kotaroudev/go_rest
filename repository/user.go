package repository

import (
	"context"

	"github.com/kotaroudev/go_rest/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
}

// Repository
// Se basa en tener abstracciones(interfaces) de los metodos que se van a usar
// para el CRUD, y no en clases concretas. La idea es que puedas usar
// cualquier base de datos y setearla con SetRepository(inyeccion de dependencias);
// sin importar la db utilizada debe cumplir con la firma de los metodos
// del repository.
// La idea central es el "desacoplamiento" entre la capa de la logica del negocio
// y la capa de la de base de datos.
// Siempre se debe cumplir con el desacoplamiento y la inyeccion de dependencias.
// La inyeccion de dependencias es la que "inyecta" la implementacion concreta
// de la base de datos utilizada.

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserByID(ctx, id)
}
