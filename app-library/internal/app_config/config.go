package app_config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ENV ENV
}

type ENV struct {
	// database mongodb
	LIBRARY_SOURCE   string
	LIBRARY_DATABASE string

	// collection mongodb
	COLLECTION_BOOKS   string
	COLLECTION_MEMBERS string

	// app port
	PORT string
}

var (
	appConfig *AppConfig
)

// Get initialized config.
func Get() *AppConfig {
	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	if err := checkENV(); err != nil {
		err = godotenv.Load()
		if err != nil {
			log.Println(" (-) file .env not found, using global variable")
		}

		err = checkENV()
		if err != nil {
			log.Panic(err)
		}
	}

	return createAppConfig()
}

func checkENV() error {
	if _, ok := os.LookupEnv("LIBRARY_SOURCE"); !ok {
		return fmt.Errorf(" (x) config: LIBRARY_SOURCE has not been set")
	}
	if _, ok := os.LookupEnv("LIBRARY_DATABASE"); !ok {
		return fmt.Errorf(" (x) config: LIBRARY_DATABASE has not been set")
	}
	if _, ok := os.LookupEnv("COLLECTION_MEMBERS"); !ok {
		return fmt.Errorf(" (x) config: COLLECTION_MEMBERS has not been set")
	}
	if _, ok := os.LookupEnv("COLLECTION_BOOKS"); !ok {
		return fmt.Errorf(" (x) config: COLLECTION_BOOKS has not been set")
	}
	if _, ok := os.LookupEnv("PORT"); !ok {
		return fmt.Errorf(" (x) config: PORT has not been set")
	}
	return nil
}

func createAppConfig() *AppConfig {
	return &AppConfig{
		ENV: ENV{
			LIBRARY_SOURCE:     os.Getenv("LIBRARY_SOURCE"),
			LIBRARY_DATABASE:   os.Getenv("LIBRARY_DATABASE"),
			COLLECTION_MEMBERS: os.Getenv("COLLECTION_MEMBERS"),
			COLLECTION_BOOKS:   os.Getenv("COLLECTION_BOOKS"),
			PORT:               os.Getenv("PORT"),
		},
	}
}
