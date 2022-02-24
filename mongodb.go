package mongodb

import (
	"context"
	"fmt"

	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Config ...
type Config struct {
	Host   string
	DBName string

	TLS        *ConnectTLSOpts
	Standalone *ConnectStandaloneOpts
}

// ConnectTLSOpts ...
type ConnectTLSOpts struct {
	ReplSet             string
	CaFile              string
	CertKeyFile         string
	CertKeyFilePassword string
}

// ConnectStandaloneOpts ...
type ConnectStandaloneOpts struct {
	AuthMechanism string
	AuthSource    string
	Username      string
	Password      string
}

var db *mongo.Database

// Connect to mongo server
func Connect(cfg Config) (*mongo.Database, error) {
	if cfg.TLS != nil && cfg.TLS.ReplSet != "" {
		return connectWithTLS(cfg)
	}
	connectOptions := options.ClientOptions{}
	opts := cfg.Standalone
	// Set auth if existed
	if opts.Username != "" && opts.Password != "" {
		connectOptions.Auth = &options.Credential{
			AuthMechanism: opts.AuthMechanism,
			AuthSource:    opts.AuthSource,
			Username:      opts.Username,
			Password:      opts.Password,
		}
	}

	// Connect
	client, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(cfg.Host))
	if err != nil {
		fmt.Println("Error when connect to MongoDB database", cfg.Host, err)
		return nil, err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO MONGODB: " + cfg.Host + " --- DB: " + cfg.DBName))

	// Set data
	db = client.Database(cfg.DBName)
	return db, nil
}

func connectWithTLS(cfg Config) (*mongo.Database, error) {
	ctx := context.Background()
	opts := cfg.TLS

	caFile, err := initFileFromBase64String("ca.pem", opts.CaFile)
	if err != nil {
		return nil, err
	}
	certFile, err := initFileFromBase64String("cert.pem", opts.CertKeyFile)
	if err != nil {
		return nil, err
	}
	pwd := base64DecodeToString(opts.CertKeyFilePassword)
	s := "%s/?tls=true&tlsCAFile=./%s&tlsCertificateKeyFile=./%s&tlsCertificateKeyFilePassword=%s&authMechanism=MONGODB-X509"
	uri := fmt.Sprintf(s, cfg.Host, caFile.Name(), certFile.Name(), pwd)
	readPref := readpref.SecondaryPreferred()
	clientOpts := options.Client().SetReadPreference(readPref).SetReplicaSet(opts.ReplSet).ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.SecondaryPreferred()); err != nil {
		return nil, err
	}
	db = client.Database(cfg.DBName)

	fmt.Println(aurora.Green("*** CONNECTED TO MONGODB: " + cfg.Host + " --- DB: " + cfg.DBName))
	return db, err
}

// GetInstance ...
func GetInstance() *mongo.Database {
	return db
}
