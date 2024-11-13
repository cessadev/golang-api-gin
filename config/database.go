package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() error {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	// Verificar la conexión
	if err := client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	// Selecciona la base de datos que usará la aplicación
	DB = client.Database("go_api_db")
	log.Println("Conectado a MongoDB")
	return err
}
