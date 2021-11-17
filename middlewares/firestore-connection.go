package middlewares

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func DbSetter(name string, db *firestore.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(name, db)
		c.Next()
	}
}
