package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateIndex ...
func CreateIndex(colName string, ind mongo.IndexModel) {
	// Get collection
	col := db.Collection(colName)
	if col == nil {
		fmt.Printf("Collection %s not existed", colName)
		return
	}

	opts := options.CreateIndexes().SetMaxTime(time.Minute * 10)
	_, err := col.Indexes().CreateOne(context.Background(), ind, opts)
	if err != nil {
		fmt.Printf("Index collection %s err: %v", col.Name(), err)
	}
}

// NewIndexKey ...
func NewIndexKey(key ...string) mongo.IndexModel {
	var doc bson.D
	for _, s := range key {
		e := bson.E{
			Key:   s,
			Value: 1,
		}
		if strings.HasPrefix(s, "-") {
			e = bson.E{
				Key:   strings.Replace(s, "-", "", 1),
				Value: -1,
			}
		}
		doc = append(doc, e)
	}

	return mongo.IndexModel{Keys: doc}
}
