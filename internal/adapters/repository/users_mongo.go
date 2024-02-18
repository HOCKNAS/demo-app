package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/HOCKNAS/demo-app/pkg/database/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// toUserMongo define la estructura de un usuario para su almacenamiento en MongoDB.
// Contiene los campos necesarios para representar un usuario y utiliza las etiquetas `bson`
// para mapear cada campo con su correspondiente nombre en la base de datos de MongoDB.
type toUserMongo struct {
	ID           string `bson:"_id,omitempty"` // ID único del usuario, omitido si está vacío.
	Name         string `bson:"name"`          // Nombre del usuario.
	LastName     string `bson:"lastname"`      // Apellido del usuario.
	Username     string `bson:"username"`      // Nombre de usuario.
	Password     string `bson:"password"`      // Contraseña del usuario (debería estar hasheada).
	Email        string `bson:"email"`         // Correo electrónico del usuario.
	IsAdmin      bool   `bson:"isAdmin"`       // Indica si el usuario es administrador.
	IsActive     bool   `bson:"isActive"`      // Indica si la cuenta del usuario está activa.
	CreationDate string `bson:"creationdate"`  // Fecha de creación de la cuenta del usuario.
}

// toUserMongoList representa una lista de usuarios de MongoDB.
// Es un slice de toUserMongo que permite trabajar con múltiples usuarios.
type toUserMongoList []toUserMongo

// ToDomain convierte un toUserMongo a un objeto User del dominio.
// Este método se utiliza para transformar datos de la estructura de MongoDB
// a la estructura del dominio utilizada en la lógica de negocio.
func (toUserMongo *toUserMongo) ToDomain() *domain.User {
	return &domain.User{
		ID:           toUserMongo.ID,
		Name:         toUserMongo.Name,
		LastName:     toUserMongo.LastName,
		Username:     toUserMongo.Username,
		Password:     toUserMongo.Password,
		Email:        toUserMongo.Email,
		IsAdmin:      toUserMongo.IsAdmin,
		IsActive:     toUserMongo.IsActive,
		CreationDate: toUserMongo.CreationDate,
	}
}

// FromDomain recibe un objeto User del dominio y actualiza los campos de toUserMongo
// con los datos del usuario del dominio. Se usa para preparar los datos antes de
// ser guardados o manipulados en MongoDB.
func (userMongo *toUserMongo) FromDomain(user *domain.User) {
	if userMongo == nil {
		userMongo = &toUserMongo{}
	}
	userMongo.ID = user.ID
	userMongo.Name = user.Name
	userMongo.LastName = user.LastName
	userMongo.Username = user.Username
	userMongo.Password = user.Password
	userMongo.Email = user.Email
	userMongo.IsAdmin = user.IsAdmin
	userMongo.IsActive = user.IsActive
	userMongo.CreationDate = user.CreationDate
}

// ToDomain convierte un userMongoList (lista de usuarios en formato MongoDB)
// a una lista de usuarios en el formato del dominio de la aplicación.
//
// Recorre cada elemento de userMongoList, que son instancias de toUserMongo,
// y utiliza el método ToDomain de toUserMongo para convertir cada usuario
// de MongoDB a su equivalente en el dominio de la aplicación (domain.User).
//
// Retorna un slice de domain.User que contiene los usuarios convertidos.
func (userMongoList toUserMongoList) ToDomain() []domain.User {
	list := make([]domain.User, len(userMongoList))
	for index, toUserMongo := range userMongoList {
		User := toUserMongo.ToDomain()
		list[index] = *User
	}

	return list
}

// mongoDbRepository es una estructura que encapsula una conexión de cliente a MongoDB.
// Se utiliza como un repositorio para realizar operaciones de base de datos, tales como
// consultas, inserciones, actualizaciones y eliminaciones en MongoDB.
//
// El campo 'mongo' es un puntero a un cliente de MongoDB, que proporciona la funcionalidad
// necesaria para interactuar con la base de datos.
type usersRepository struct {
	db *mongo.Collection
}

// NewUsersRepository crea y retorna una nueva instancia de usersRepository.
// Este repositorio se utiliza para interactuar con la colección de usuarios en MongoDB.
//
// db: Es la base de datos de MongoDB en la que se encuentra la colección de usuarios.
//
// Retorna un nuevo repositorio que implementa la interfaz UsersRepositoy.
func NewUsersRepository(db *mongo.Database) ports.UsersRepository {
	return &usersRepository{
		db: db.Collection("Users"),
	}
}

// Create es un método de usersRepository que crea un nuevo usuario en la base de datos de MongoDB.
//
// ctx: Contexto para controlar el tiempo de vida de la solicitud de base de datos.
// user: Es el objeto de usuario del dominio que se va a almacenar en la base de datos.
//
// Realiza las siguientes operaciones:
// 1. Genera un nuevo UUID para el ID del usuario y establece la fecha de creación actual.
// 2. Convierte el objeto del dominio user a una estructura toUserMongo adecuada para MongoDB.
// 3. Inserta el usuario convertido en la base de datos de MongoDB.
// 4. Si el usuario ya existe (detectado por un error de duplicado), retorna un error específico.
//
// Retorna el usuario recién creado en el formato del dominio, o un error en caso de fallo.
func (r *usersRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	user.ID = uuid.New().String()
	user.CreationDate = time.Now().UTC().Format("2006-01-02T15:04:05.000000000")

	var toUserMongo toUserMongo = toUserMongo{}
	toUserMongo.FromDomain(user)

	_, err := r.db.InsertOne(ctx, toUserMongo)
	if mongodb.IsDuplicate(err) {
		return nil, domain.ErrUserAlreadyExistsDB
	}

	return toUserMongo.ToDomain(), err
}

func (r *usersRepository) Delete(ctx context.Context, id string) error {

	filter := bson.M{"_id": id}

	result, err := r.db.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return domain.ErrUserNotFoundForDeletionDB
	}

	return nil
}

func (r *usersRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var mongoUser toUserMongo
	filter := bson.M{"_id": id}
	err := r.db.FindOne(ctx, filter).Decode(&mongoUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w con ID: %s", domain.ErrUserNotFoundDB, id)
		}
		return nil, fmt.Errorf("%w al buscar el usuario con ID: %s, error: %v", domain.ErrDatabase, id, err)
	}

	return mongoUser.ToDomain(), err
}

func (r *usersRepository) Deactivate(ctx context.Context, id string) (*domain.User, error) {
	_, err := r.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFoundDB) {
			return nil, fmt.Errorf("%w con ID: %s", domain.ErrUserNotFoundDB, id)
		}
		return nil, fmt.Errorf("%w al verificar el usuario con ID: %s", domain.ErrDatabase, id)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isActive": false}}
	result, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("%w al desactivar el usuario con ID: %s", domain.ErrDatabase, id)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("%w con ID: %s", domain.ErrUserNotFoundDB, id)
	}

	updatedUser, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w al recargar el usuario desactivado con ID: %s", domain.ErrDatabase, id)
	}

	return updatedUser, err
}
