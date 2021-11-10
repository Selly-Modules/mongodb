package mongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// IsErrNotFound check err is not found document
func IsErrNotFound(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}
