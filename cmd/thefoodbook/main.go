package main

import (
	"context"
	"github.com/mariamihai/thefoodbook/configs"
	"log"
	"time"
)

func main() {
	c, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Error encountered when loading the configuration: %v", err)
	}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	driver, err := configs.NewDriver(ctxWithTimeout, c.DBUri, c.DBUser, c.DBPass)

	if err != nil {
		log.Fatalf("Couldn't connect to the db: %v", err)
	}

	err = configs.CloseDriver(ctxWithTimeout, driver)
	if err != nil {
		log.Fatalf("Couldn't close the db: %v", err)
	}
}
