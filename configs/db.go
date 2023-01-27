package configs

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func NewDriver(ctx context.Context, uri, username, password string) (neo4j.DriverWithContext, error) {
	// Create Driver
	driverWithContext, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))

	// Handle any driver creation errors
	if err != nil {
		return nil, err
	}

	// Verify Connectivity
	err = driverWithContext.VerifyConnectivity(ctx)

	// If connectivity fails, handle the error
	if err != nil {
		return nil, err
	}

	log.Println("Obtained new driver with context")
	return driverWithContext, nil
}

// CloseDriver call on application exit
func CloseDriver(ctx context.Context, driver neo4j.DriverWithContext) error {
	log.Println("Closing the driver")
	return driver.Close(ctx)
}
