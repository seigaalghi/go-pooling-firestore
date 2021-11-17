package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

func Connect(name string) (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cds-kompasid-dev")
	if err != nil {
		return nil, err
	}
	// defer fmt.Println(name, " is failed to connect to firestore")
	fmt.Println(name, " is connected to firestore")
	return client, nil
}
