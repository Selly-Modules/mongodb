package mongodb

import (
	"context"
	"fmt"

	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectTLSOpts ...
type ConnectTLSOpts struct {
	Host                       string
	DBName                     string
	ReplSet                    string
	CaFilePath                 string
	CertificateKeyFilePath     string
	CertificateKeyFilePassword string
}

var db *mongo.Database

// ConnectWithTLS ...
func ConnectWithTLS(opts ConnectTLSOpts) (*mongo.Database, error) {
	ctx := context.Background()
	uri := fmt.Sprintf("%s/?tls=true&tlsCAFile=%s&tlsCertificateKeyFile=%s&tlsCertificateKeyFilePassword=%s", opts.Host, opts.CaFilePath, opts.CertificateKeyFilePath, opts.CertificateKeyFilePassword)
	readPref := readpref.SecondaryPreferred()
	credential := options.Credential{
		AuthMechanism: "MONGODB-X509",
	}
	clientOpts := options.Client().SetAuth(credential).SetReadPreference(readPref).SetReplicaSet(opts.ReplSet).ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.SecondaryPreferred()); err != nil {
		return nil, err
	}
	db = client.Database(opts.DBName)
	return db, err
}

// Connect to mongo server
func Connect(host, user, password, dbName, mechanism, source string) (*mongo.Database, error) {
	connectOptions := options.ClientOptions{}
	// Set auth if existed
	if user != "" && password != "" {
		connectOptions.Auth = &options.Credential{
			AuthMechanism: mechanism,
			AuthSource:    source,
			Username:      user,
			Password:      password,
		}
	}

	// Connect
	client, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(host))
	if err != nil {
		fmt.Println("Error when connect to MongoDB database", host, err)
		return nil, err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO MONGODB: " + host + " --- DB: " + dbName))

	// Set data
	db = client.Database(dbName)
	return db, nil
}

// GetInstance ...
func GetInstance() *mongo.Database {
	return db
}
