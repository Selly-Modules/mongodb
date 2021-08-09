package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

// NewObjectID generates a new ObjectID based on the timestamp.
func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// NewStringID ...
func NewStringID() string {
	return NewObjectID().Hex()
}

// ConvertFromString ...
func ConvertFromString(s string) (value primitive.ObjectID, isValid bool) {
	id, err := primitive.ObjectIDFromHex(s)
	return id, err == nil
}

// IsValid ...
func IsValid(s string) bool {
	return primitive.IsValidObjectID(s)
}

