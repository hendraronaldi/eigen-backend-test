package databases

import (
	"app-library/internal/app_config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Instance struct {
	config        *app_config.AppConfig
	LibraryClient *mongo.Client
	LibraryDB     *mongo.Database
}

// New database instance constructor.
func New() *Instance {
	di := &Instance{}
	di.config = app_config.Get()
	return di
}

// ConnectLibraryDB: connect database instance of library mongodb database.
func (di *Instance) ConnectLibraryDB() {
	var err error
	if di.LibraryClient == nil {
		di.LibraryClient, err = mongo.Connect(context.TODO(), options.Client(), options.Client().ApplyURI(di.config.ENV.LIBRARY_SOURCE))
		if err != nil {
			log.Panicf(" (x) database error (connect): cannot connect to mongodb product\n")
		}
	}
	di.LibraryDB = di.LibraryClient.Database(di.config.ENV.LIBRARY_DATABASE)
}

// CloseLibraryDB: close database connection
func (di *Instance) CloseLibraryDB() {
	if di.LibraryClient == nil {
		return
	}

	err := di.LibraryClient.Disconnect(context.TODO())
	if err != nil {
		log.Println(" (x) error closing connection to mongodb library")
	}
}
