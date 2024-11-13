package repository

import (
	"context"
	"go-api-gin/config"
	"go-api-gin/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

// Constructor para inicializar el repositorio con la colección
func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: config.DB.Collection("users"), // Inicializamos la colección de usuarios
	}
}

// Método de repositorio para consultar un usuario por su ID
func (r *UserRepository) FindUserByID(id string) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user domain.User
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Método de repositorio para crear un nuevo usuario en la base de datos
func (r *UserRepository) CreateUser(user *domain.User) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), user)
}

// Método de repositorio para consultar un usuario por su email
func (r *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"email": email}

	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments // No se encontró el usuario
		}
		return nil, err // Otro tipo de error
	}

	return &user, nil
}
