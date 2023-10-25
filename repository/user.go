package repository

import (
	"context"

	"github.com/kotaroudev/go_rest/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	Close() error
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

// Arquitecturas limpias
// Se basan en que giran en torno al "dominio".
// El dominio no debe ser expuesto y debe estar totalmente desacoplado de las
// capas externas. El dominio debe funcionar sin importar todas las demas
// capas; sin importar si hay o no sistema.
// Una analogia podria ser que el dominio sea un programa de consola
// donde funciona toda la logica del negocio pero no tiene frameworks, ni base
// de datos, ni ninguna integracion; sino que todo funciona de manera basica
// mundana se podria decir, pero contiene todo el funcionamiento del core del negocio.
// Por ejemplo en una red social el core del negocio seria:
// - publicar posts
// - mandar mensajes
// - comentar posts
// - editar perfil
// - etc
// Estas funciones deberian poder ejecutarse con datos primitivos desacoplados
//  de las capas externas.
// En esto se basan todas las arquitecturas limpias.
// En nuestro ejemplo el "dominio o core" viene a ser:
// - intertar usuario
// - hacer login
// - obtener un usario
// - hacer un logout
// - etc

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

func Close() error {
	return implementation.Close()
}
