package mongodb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// CreateIndex ...
func CreateIndex(colName string, ind mongo.IndexModel) {
	// Get collection
	col := db.Collection(colName)
	if col == nil {
		fmt.Printf("Collection %s not existed", colName)
		return
	}

	fmt.Println("ind", ind)

	opts := options.CreateIndexes().SetMaxTime(time.Minute * 10)
	_, err := col.Indexes().CreateOne(context.Background(), ind, opts)
	if err != nil {
		fmt.Printf("Index collection %s err: %v", col.Name(), err)
	}
}

// NewIndexKey ...
func NewIndexKey(key ...string) mongo.IndexModel {
	var doc bsonx.Doc
	for _, s := range key {
		e := bsonx.Elem{
			Key:   s,
			Value: bsonx.Int32(1),
		}
		if strings.HasPrefix(s, "-") {
			e = bsonx.Elem{
				Key:   strings.Replace(s, "-", "", 1),
				Value: bsonx.Int32(-1),
			}
		}
		doc = append(doc, e)
	}

	return mongo.IndexModel{Keys: doc}
}