package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

var (
	// Pool - redis pool
	fdb *firestore.Client
)

func init() {
	firestore, err := Connect("connection1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("sekali init aja di awal")

	fdb = firestore
}

func GetArticleFirestore(ctx *gin.Context) *firestore.DocumentSnapshot {
	doc, err := fdb.Collection("publish_article/Kompas.id/Article").Doc("gelombang-keempat-mengintai-eropa-kembali-waspada").Get(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	return doc
}

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
