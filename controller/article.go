package controller

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/go-pooling-firestore/database"
	"google.golang.org/api/iterator"
)

func GetArticle(ctx *gin.Context) {
	var output map[string]interface{}
	db, _ := ctx.MustGet("dbV1").(*firestore.Client)
	doc, err := db.Collection("publish_article/Kompas.id/Article").Doc("gelombang-keempat-mengintai-eropa-kembali-waspada").Get(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	doc.DataTo(&output)
	ctx.JSON(200, gin.H{
		"data": output,
	})
}

func GetArticleFirestore(ctx *gin.Context) {
	var output map[string]interface{}
	// db, _ := ctx.MustGet("dbV1").(*firestore.Client)
	doc := database.GetArticleFirestore(ctx)
	doc.DataTo(&output)
	ctx.JSON(200, gin.H{
		"data": output,
	})
}

func GetArticles(ctx *gin.Context) {
	var output []map[string]interface{}
	db, _ := ctx.MustGet("dbV2").(*firestore.Client)
	iter := db.Collection("publish_article/Kompas.id/Article").Limit(10).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		var result map[string]interface{}
		doc.DataTo(&result)

		output = append(output, result)
	}

	ctx.JSON(200, gin.H{
		"data": output,
	})
}
